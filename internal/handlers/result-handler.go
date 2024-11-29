package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
    "prompt-game/views/pages"
)

type ResultHandler struct {
}

func NewResultHandler() *PromptHandler {
	return &PromptHandler{}
}

func (h *PromptHandler) GetResult() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := render(ctx, http.StatusOK, pages.ResultPage())
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render page"})
		}
	}
}
