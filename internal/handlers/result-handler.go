package handlers

import (
	"net/http"

	"prompt-game/views"
	"prompt-game/views/pages/result"

	"github.com/gin-gonic/gin"
)

type ResultHandler struct {
}

func NewResultHandler() *PromptHandler {
	return &PromptHandler{}
}

func (h *PromptHandler) GetResult() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := render(ctx, http.StatusOK, views.Layout(result.ResultPage()))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render page"})
		}
	}
}
