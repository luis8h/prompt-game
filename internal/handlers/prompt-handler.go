package handlers

import (
	"fmt"
	"net/http"
	"prompt-game/external/openai"
	"prompt-game/views/components"

	"github.com/gin-gonic/gin"
)

type PromptHandler struct {
    api *openai.Api
    messages []openai.Message
}

func NewPromptHandler(apiKey string) *PromptHandler {
	return &PromptHandler{
        api: openai.NewApi(apiKey),
        messages: []openai.Message{},
    }
}

// TODO: save message history per session -> chatgpt chat
func (h *PromptHandler) PostPromptUser() gin.HandlerFunc {
    return func (ctx *gin.Context)  {
        message := ctx.PostForm("prompt_input")

        if message == "" {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": "empty request body"})
            return
        }

        viewMessage := components.Message{Role: "user", Content: message}
        err := render(ctx, http.StatusOK, components.ChatMessage(viewMessage))

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render page"})
		}
    }
}

func (h *PromptHandler) PostPromptAssistant() gin.HandlerFunc {
	return func(ctx *gin.Context) {
        message := ctx.PostForm("prompt_input")
        fmt.Println(ctx)

        if message == "" {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": "empty request body"})
            return
        }

        newMessage := openai.Message{Role: "user", Content: message}
		h.messages = append(h.messages, newMessage)

        resp, err := h.api.Get(message, h.messages)

        if (err != nil) {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error in openai api request"})
            return
        }

		assistantMessage := openai.Message{Role: "assistant", Content: resp}
		h.messages = append(h.messages, assistantMessage)

        viewMessage := components.Message{Role: assistantMessage.Role, Content: assistantMessage.Content}
		err = render(ctx, http.StatusOK, components.ChatMessage(viewMessage))

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render page"})
		}
	}
}
