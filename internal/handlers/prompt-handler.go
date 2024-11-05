package handlers

import (
	"net/http"
	"prompt-game/external/openai"

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

func (h *PromptHandler) PostPrompt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
        message := ctx.PostForm("prompt-input")

        if message == "" {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": "empty request body"})
            return
        }

        resp, err := h.api.Get(message)

        if (err != nil) {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error in openai api request"})
            return
        }

		ctx.JSON(http.StatusOK, resp)
	}
}
