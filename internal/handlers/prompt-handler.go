package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"prompt-game/external/openai"
	"prompt-game/views/pages/index"

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

        messages := []index.Message{}
		if err := json.Unmarshal([]byte(messagesJson), &messages); err != nil {
			fmt.Println("Error unmarshalling messages:", err)
			return
		}

        err := render(ctx, http.StatusOK, index.ChatHistory(messages))
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render page"})
        }
    }
}

func (h *PromptHandler) PostMessageUser() gin.HandlerFunc {
    return func(ctx *gin.Context) {
		message := ctx.PostForm("message")

		if message == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "empty request body"})
			return
		}

		viewMessage := index.Message{Role: "user", Content: message}

		err := render(ctx, http.StatusOK, index.ChatMessage(viewMessage))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render page"})
		}
    }
}

func (h *PromptHandler) PostMessageAssistant() gin.HandlerFunc {
    return func(ctx *gin.Context) {
		message := ctx.PostForm("message")

		if message == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "empty request body"})
			return
		}

		viewMessage := index.Message{Role: "assistant", Content: message}

		err := render(ctx, http.StatusOK, index.ChatMessage(viewMessage))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render page"})
		}
    }
}

func (h *PromptHandler) PostPrompt() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        var payload struct {
            Messages []openai.Message `json:"messages"`
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
		resp, err := h.api.RequestApi(payload.Messages)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error in openai api request"})
			return
		}

        ctx.JSON(http.StatusOK, gin.H{"answer": resp})
    }
}
