package handlers

import (
    "prompt-game/external/openai"
    "github.com/gin-gonic/gin"
    "net/http"
)

type PromptHandler struct {
    api openai.Api
}

func NewPromptHandler() *PromptHandler {
	return &PromptHandler{}
}

func (h *PromptHandler) PostPrompt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "hallo zurück")
	}
}
