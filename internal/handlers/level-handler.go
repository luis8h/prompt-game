package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"prompt-game/external/openai"
	"prompt-game/internal/models"
	"prompt-game/internal/stores"
	"prompt-game/internal/utils"
	"prompt-game/views"
	"prompt-game/views/pages/game"
	"prompt-game/views/pages/result"
	"strings"

	"github.com/gin-gonic/gin"
)

type LevelHandler struct {
	api *openai.Api
}

func NewLevelHandler(apiKey string) *LevelHandler {
	return &LevelHandler{
		api: openai.NewApi(apiKey),
	}
}

func (h *LevelHandler) GetLevelStoryPrev() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		locale := getLocale(ctx)

		levelId := GetCurrentLevel(ctx)
		level := stores.GetLevel(levelId, locale)

		storyId := GetStoryId(ctx)

		newStoryId := storyId - 1
		if (len(level.Story) <= newStoryId || newStoryId < 0) {
			newStoryId = 0
		}

		SetStoryId(ctx, newStoryId)

		story := level.Story[newStoryId]

		utils.GameLogger.PrintS(ctx, fmt.Sprintf("switched to new story id with back %d", newStoryId))

		render(ctx, http.StatusOK, game.StoryHtml(story, newStoryId, len(level.Story)))
	}
}

func (h *LevelHandler) GetLevelStoryNext() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		locale := getLocale(ctx)

		levelId := GetCurrentLevel(ctx)
		level := stores.GetLevel(levelId, locale)

		storyId := GetStoryId(ctx)

		newStoryId := storyId + 1

		if (len(level.Story)-1 == newStoryId) {
			SetShowTask(ctx, true)
			ctx.Writer.Header().Set("HX-Trigger", "refreshLevel")
		}

		if (len(level.Story) <= newStoryId || newStoryId < 0) {
			newStoryId = 0
		}

		SetStoryId(ctx, newStoryId)
		story := level.Story[newStoryId]

		utils.GameLogger.PrintS(ctx, fmt.Sprintf("switched to new story id with next %d", newStoryId))

		render(ctx, http.StatusOK, game.StoryHtml(story, newStoryId, len(level.Story)))
	}
}

func (h *LevelHandler) GetLevel() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		locale := getLocale(ctx)
		levelId := GetCurrentLevel(ctx)
		storyId := GetStoryId(ctx)
		withStrategy := GetWithStrategy(ctx)
		showTask := GetShowTask(ctx)

		level := stores.GetLevel(levelId, locale)

		render(ctx, http.StatusOK, game.LevelHtml(level, withStrategy, false, levelId, storyId, showTask))
	}
}

func (h *LevelHandler) PostLevelNextA() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		locale := getLocale(ctx)

		withStrategy := SetWithStrategy(ctx, true)
		showTask := SetShowTask(ctx, true)

		levelId := GetCurrentLevel(ctx)
		storyId := GetStoryId(ctx)

		utils.GameLogger.PrintS(ctx, fmt.Sprintf("revealed strategy"))

		render(ctx, http.StatusOK, game.InstructionsPane(stores.GetLevel(levelId, locale), true, withStrategy, levelId, storyId, showTask))
	}
}

func (h *LevelHandler) PostLevelNextB() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		locale := getLocale(ctx)
		levelId := GetCurrentLevel(ctx)
		storyId := GetStoryId(ctx)
		showTask := GetShowTask(ctx)

		// get messages
		messagesJson := ctx.PostForm("messages")
		messages := []openai.Message{}
		if err := json.Unmarshal([]byte(messagesJson), &messages); err != nil {
			fmt.Println("Error unmarshalling messages:", err)
			return
		}

		utils.GameLogger.PrintS(ctx, fmt.Sprintf("submitted"))

		// verify answer
		valid := h.validateLevel(ctx, messages, stores.GetLevel(levelId, locale))

		// render invalid template
		if !valid {
			utils.GameLogger.PrintS(ctx, fmt.Sprintf("solution NOT valid"))
			// render template
			ctx.Writer.Header().Set("HX-Trigger", "invalidAnswer")
			render(ctx, http.StatusOK, game.InstructionsPane(stores.GetLevel(levelId, locale), true, false, levelId, storyId, showTask))
			return
		}

		utils.GameLogger.PrintS(ctx, fmt.Sprintf("solution valid"))

		// set nextlevel
		nextLevelId := levelId + 1
		SetCurrentLevel(ctx, nextLevelId)
		withStrategy := SetWithStrategy(ctx, false)
		newStoryId := SetStoryId(ctx, 0)
		showTask = SetShowTask(ctx, false)

		if (stores.GetLevel(levelId, locale).ClearChatHistoryOnSubmit) {
			ctx.Writer.Header().Set("HX-Trigger", "resetChatHistory")
		}

		// load results page
		if nextLevelId == stores.GetLevelCount() {
			utils.GameLogger.PrintS(ctx, fmt.Sprintf("finished - rendering result page"))
			ctx.Writer.Header().Set("HX-Retarget", "#page-container")
			render(ctx, http.StatusOK, views.Layout(result.ResultPage(), GetSessionId(ctx)))
			return
		}

		// render next level
		render(ctx, http.StatusOK, game.InstructionsPane(stores.GetLevel(nextLevelId, locale), withStrategy, true, nextLevelId, newStoryId, showTask))
	}
}

func (h *LevelHandler) validateLevel(ctx *gin.Context, messages []openai.Message, level models.Level) (bool) {
	if len(messages) == 0 {
		return false
	}

	// verify answer
	isAnswerValid, err := h.isValidAnswer(ctx, messages, level)
	if err != nil {
		fmt.Printf("error when validating answer: %v", err)
		return false
	}

	if !level.HasStrategy {
		return isAnswerValid;
	}

	// verify strategy
	isStrategyValid, err := h.isValidStrategy(ctx, messages, level)
	if err != nil {
		fmt.Printf("error when validating strategy: %v", err)
		return false
	}

	if isStrategyValid && isAnswerValid {
		return true
	}
	return false
}

func (h *LevelHandler) isValidAnswer(ctx *gin.Context, messages []openai.Message, level models.Level) (bool, error) {
	prompt := fmt.Sprintf(stores.ValidateAnswerPrompt, stores.ElfName, h.getChatHistory(messages), level.Task)

	jsonResponse, err := h.getVerificationResponse(prompt)
	if err != nil {
		return false, err
	}

	utils.GameLogger.PrintS(ctx, fmt.Sprintf("validating answer was %t with following prompt:\n", jsonResponse.Verified, prompt))

	return jsonResponse.Verified, nil
}

func (h *LevelHandler) isValidStrategy(ctx *gin.Context, messages []openai.Message, level models.Level) (bool, error) {
	prompt := fmt.Sprintf(stores.ValidateStrategyPrompt, stores.ElfName, level.StrategyValidation, h.getChatHistory(messages))

	jsonResponse, err := h.getVerificationResponse(prompt)
	if err != nil {
		return false, err
	}

	utils.GameLogger.PrintS(ctx, fmt.Sprintf("validating strategy was %t with following prompt:\n", jsonResponse.Verified, prompt))

	return jsonResponse.Verified, nil
}

func (h *LevelHandler) getVerificationResponse(prompt string) (*models.VerificationResponse, error) {
	// get ai response
	strResponse, err := h.api.GetAnswer(prompt, []openai.Message{})
	if err != nil {
		return nil, fmt.Errorf("error in api request %v", err)
	}

	// trim response (only take content between { })
	trimmed, err := h.trimResponse(strResponse)
	if err != nil {
		return nil, fmt.Errorf("error when trimming responsed: %v", err)
	}

	// convert json to object
	var jsonResponse models.VerificationResponse
	if err := json.Unmarshal(([]byte(trimmed)), &jsonResponse); err != nil {
		return nil, fmt.Errorf("failed to parse ai response '%s' to object: %v", trimmed, err)
	}

	return &jsonResponse, nil
}

func (h *LevelHandler) getChatHistory(messages []openai.Message) string {
	var chatHistory string

	for _, message := range messages {
		if message.Role == "user" {
			chatHistory = chatHistory + "\n\nuser:\n"
		} else if message.Role == "assistant" {
			chatHistory = chatHistory + "\n\nassistant:\n"
		} else {
			fmt.Println("role appart from 'assistant' and 'user' was found in message")
			continue
		}

		chatHistory = chatHistory + fmt.Sprintf("'%s'", message.Content)
	}

	return chatHistory
}

func (h *LevelHandler) trimResponse(response string) (string, error) {
	jsonStart := strings.Index(response, "{")
	jsonEnd := strings.Index(response, "}")

	if jsonStart == -1 || jsonEnd == -1 || jsonStart >= jsonEnd {
		return "", errors.New("invalid json object: missing or misplaced '{' or '}'")
	}

	return response[jsonStart : jsonEnd+1], nil
}
