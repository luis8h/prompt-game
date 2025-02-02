package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"prompt-game/external/openai"
	"prompt-game/internal/stores"
	"prompt-game/internal/utils"
	"prompt-game/views/pages/game"

	"github.com/gin-gonic/gin"
)

type PromptHandler struct {
	api *openai.Api
}

func NewPromptHandler(apiKey string) *PromptHandler {
	return &PromptHandler{
		api: openai.NewApi(apiKey),
	}
}

func (h *PromptHandler) PostHistory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		messagesJson := ctx.PostForm("messages")

		messages := []game.Message{}
		if err := json.Unmarshal([]byte(messagesJson), &messages); err != nil {
			fmt.Println("Error unmarshalling messages:", err)
			return
		}

		render(ctx, http.StatusOK, game.ChatHistory(messages))
	}
}

func (h *PromptHandler) PostMessageUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		message := ctx.PostForm("message")

		if message == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "empty request body"})
			return
		}

		utils.GameLogger.PrintS(ctx, "postet message: " + message)

		viewMessage := game.Message{Role: "user", Content: message}

		render(ctx, http.StatusOK, game.ChatMessage(viewMessage))
	}
}

func (h *PromptHandler) PostMessageAssistant() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		message := ctx.PostForm("message")

		if message == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "empty request body"})
			return
		}

		utils.GameLogger.PrintS(ctx, "assistant replied: "+message)

		viewMessage := game.Message{Role: "assistant", Content: message}

		render(ctx, http.StatusOK, game.ChatMessage(viewMessage))
	}
}

func (h *PromptHandler) PostPrompt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var payload struct {
			Messages []openai.Message `json:"messages"`
		}

		locale := getLocale(ctx)
		withStrategy := GetWithStrategy(ctx)
		levelId := GetCurrentLevel(ctx)
		level := stores.GetLevel(levelId, locale)

		// set system messages
		systemMessages := []openai.Message{
			{
				Role:    "system",
				Content: stores.ElfSysPrompt,
			},
		}

		if !withStrategy && level.HasStrategy {
			systemMessages = append(systemMessages, openai.Message{
				Role:    "system",
				Content: stores.BadSysPrompt,
			})
			systemMessages = append(systemMessages, openai.Message{
				Role:    "system",
				Content: level.BadPrompt,
			})
		}

		// Convert systemMessages to a pretty-printed JSON string
		systemMessagesJSON, err := json.MarshalIndent(systemMessages, "", "  ")
		if err != nil {
			utils.GameLogger.PrintS(ctx, "Error marshalling systemMessages: "+err.Error())
		} else {
			utils.GameLogger.PrintS(ctx, "Prompting OpenAI with systemMessages:\n"+string(systemMessagesJSON))
		}

		// Parse the JSON payload from the request
		if err := ctx.ShouldBindJSON(&payload); err != nil {
			fmt.Println("Error binding JSON:", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON payload"})
			return
		}

		// Validate that messages are present
		if len(payload.Messages) == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "messages field is empty"})
			return
		}

		// openai api call
		resp, err := h.api.RequestApiSystem(payload.Messages, systemMessages)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error in openai api request"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"answer": resp})
	}
}
