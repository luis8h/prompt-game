package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"prompt-game/external/openai"
	"prompt-game/internal/models"
	"prompt-game/internal/stores"
	"prompt-game/views"
	"prompt-game/views/pages/game"
	"prompt-game/views/pages/result"
	"strings"

	"github.com/gin-contrib/sessions"
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

func (h *LevelHandler) PostLevelNextA() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		locale := getLocale(ctx)

		session.Set("withStrategy", true)
		session.Save()

		levelId, ok := session.Get("currentLevel").(int)
		if !ok {
			levelId = 0
			session.Set("currentLevel", 0)
			session.Save()
		}

		// render template
		err := render(ctx, http.StatusOK, game.InstructionsPane(stores.GetLevel(levelId, locale), true, true, levelId))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render page"})
		}
	}
}

func (h *LevelHandler) PostLevelNextB() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		locale := getLocale(ctx)

		// get current level
		levelId, ok := session.Get("currentLevel").(int)
		if !ok {
			levelId = 0
			session.Set("currentLevel", 0)
		}

		// get messages
		messagesJson := ctx.PostForm("messages")
		messages := []openai.Message{}
		if err := json.Unmarshal([]byte(messagesJson), &messages); err != nil {
			fmt.Println("Error unmarshalling messages:", err)
			return
		}

		// verify answer
		valid := h.validateLevel(messages, stores.GetLevel(levelId, locale))

		// render invalid template
		if !valid {
			// render template
			err := render(ctx, http.StatusOK, game.InstructionsPane(stores.GetLevel(levelId, locale), true, false, levelId))
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render page"})
			}
			return
		}

		// set nextlevel
		nextLevelId := levelId + 1
		session.Set("currentLevel", nextLevelId)
		session.Set("withStrategy", false)
		session.Save()

		if (stores.GetLevel(nextLevelId, locale).ClearChatHistoryOnSubmit) {
			ctx.Writer.Header().Set("HX-Trigger", "resetChatHistory")
		}

		// load results page
		if nextLevelId == stores.GetLevelCount() {
			ctx.Writer.Header().Set("HX-Retarget", "#page-container")
			err := render(ctx, http.StatusOK, views.Layout(result.ResultPage()))
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render page"})
			}
			return
		}

		// render next level
		err := render(ctx, http.StatusOK, game.InstructionsPane(stores.GetLevel(nextLevelId, locale), false, true, nextLevelId))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render page"})
		}
	}
}

func (h *LevelHandler) validateLevel(messages []openai.Message, level models.Level) (bool) {
	if len(messages) == 0 {
		return false
	}

	// verify answer
	isAnswerValid, err := h.isValidAnswer(messages, level)
	if err != nil {
		fmt.Printf("error when validating answer: %v", err)
		return false
	}

	if !level.HasStrategy {
		return isAnswerValid;
	}

	// verify strategy
	isStrategyValid, err := h.isValidStrategy(messages, level)
	if err != nil {
		fmt.Printf("error when validating strategy: %v", err)
		return false
	}

	fmt.Printf("Strategy: %s, Answer: %s", isStrategyValid, isAnswerValid)

	if isStrategyValid && isAnswerValid {
		return true
	}
	return false
}

func (h *LevelHandler) isValidAnswer(messages []openai.Message, level models.Level) (bool, error) {
	prompt := fmt.Sprintf(stores.ValidateAnswerPrompt, stores.ElveName, h.getChatHistory(messages), level.Task)

	jsonResponse, err := h.getVerificationResponse(prompt)
	if err != nil {
		return false, err
	}

	return jsonResponse.Verified, nil
}

func (h *LevelHandler) isValidStrategy(messages []openai.Message, level models.Level) (bool, error) {
	prompt := fmt.Sprintf(stores.ValidateStrategyPrompt, stores.ElveName, level.StrategyValidation, h.getChatHistory(messages))

	jsonResponse, err := h.getVerificationResponse(prompt)
	if err != nil {
		return false, err
	}

	fmt.Printf("Strategy prompt: %s\n", prompt)
	fmt.Printf("Strategy: %b\n", jsonResponse.Verified)

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
