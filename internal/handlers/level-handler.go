package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"prompt-game/external/openai"
	"prompt-game/internal/models"
	"prompt-game/internal/stores"
	"prompt-game/views/components"
	"strconv"
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

func (h *LevelHandler) GetLevelSubmit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get level from store
		levelId := ctx.Param("levelId")
		levelIdInt, err := strconv.Atoi(levelId)
		if err != nil {
			fmt.Println("error when converting id to integer: %v", err)
			return
		}
		level := stores.Levels[levelIdInt]

		// get messages from session
		session := sessions.Default(ctx)
		messageData, ok := session.Get("messages").([]byte)
		if !ok {
			fmt.Println("Error: messages are not stored as []byte")
			return
		}
		var messageSlice []openai.Message
		if err := json.Unmarshal(messageData, &messageSlice); err != nil {
			fmt.Println("Error unmarshalling messages:", err)
			return
		}

		// verify using openai api
		valid, err := h.isValidStrategy(messageSlice, level)
		if err != nil {
			fmt.Println("error when validating strategy: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// render template
		err = render(ctx, http.StatusOK, components.LevelFeedbackHtml(valid))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render page"})
		}
	}
}

func (h *LevelHandler) isValidAnswer(messages []openai.Message, level models.Level) (bool, error) {
    prompt := fmt.Sprintf(`
        test
    `)

    jsonResponse, err := h.getVerificationResponse(prompt)
    if err != nil {
        return false, err
    }

	return jsonResponse.Verified, nil
}

func (h *LevelHandler) isValidStrategy(messages []openai.Message, level models.Level) (bool, error) {
	prompt := fmt.Sprintf(`
        Background: I created a game, where the users can learn prompt engineering by solving different tasks using various prompting techniques.
        Your task is to decide wether in the following chat between the user and the ai, a specific prompt engineering strategy was used by the user.

        Chat of the user and the ai assistant:
        %s

        Task which should be solved by the user:
        %s

        Further information about the prompting strategy:
        %s

        Your reply should be a json string and **nothing else** which has an attribute called "verified".
        This attribute should contain a true value if the user used the right strategy and a false value if he didn't.
    `, h.getChatHistory(messages), level.Description, level.Strategy)

    jsonResponse, err := h.getVerificationResponse(prompt)
    if err != nil {
        return false, err
    }

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

func (h *LevelHandler) getChatHistory(messages []openai.Message) (string) {
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
