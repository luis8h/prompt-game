package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"prompt-game/external/openai"
	"prompt-game/internal/models"
	"prompt-game/views/components"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LevelHandler struct {
	api *openai.Api
}

func NewLevelHandler(apiKey string) *PromptHandler {
	return &PromptHandler{
		api: openai.NewApi(apiKey),
	}
}

func (h *PromptHandler) GetLevelSubmit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)

		// get messages from session
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

		// create sample level
		level := models.Level{
			Title:       "Sample level",
			Description: "Use the llm to write a poem in the style of the 19th century. To get better results you should give the llm a role first.",
			Strategy:    "The user should use the Strategy called Role-Prompting. That means he should give the ai assistant a role which suits his needs.",
		}

		// verify using openai api
		valid, err := h.isValidStrategy(messageSlice, level)
        if err != nil {
            fmt.Println("error when validating strategy: %v", err)
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

		err = render(ctx, http.StatusOK, components.LevelFeedbackHtml(valid))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render page"})
		}
	}
}

func (h *PromptHandler) trimResponse(response string) (string, error) {
    jsonStart := strings.Index(response, "{")
    jsonEnd := strings.Index(response, "}")

    if jsonStart == -1 || jsonEnd == -1 || jsonStart >= jsonEnd {
        return "", errors.New("invalid json object: missing or misplaced '{' or '}'")
    }

    return response[jsonStart : jsonEnd+1], nil
}

func (h *PromptHandler) isValidStrategy(messages []openai.Message, level models.Level) (bool, error) {
	// configure prompt
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
    `, chatHistory, level.Description, level.Strategy)

	// get ai response
	strResponse, err := h.api.GetAnswer(prompt, []openai.Message{})
	if err != nil {
		return false, fmt.Errorf("error in api request %v", err)
	}

    // trim response (only take content between { })
    trimmed, err := h.trimResponse(strResponse)
    if err != nil {
        return false, fmt.Errorf("error when trimming responsed: %v", err)
    }

	// convert json to object
	var jsonResponse models.VerificationResponse
	if err := json.Unmarshal(([]byte(trimmed)), &jsonResponse); err != nil {
		return false, fmt.Errorf("failed to parse ai response '%s' to object: %v", trimmed, err)
	}

	return jsonResponse.Verified, nil
}
