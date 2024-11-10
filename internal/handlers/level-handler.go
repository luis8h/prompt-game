package handlers

import (
	"net/http"
	"prompt-game/external/openai"
	"prompt-game/views/components"

	"github.com/gin-gonic/gin"
)

type LevelHandler struct {
	api      *openai.Api
}

func NewLevelHandler(apiKey string) *PromptHandler {
	return &PromptHandler{
		api:      openai.NewApi(apiKey),
	}
}

func (h *PromptHandler) PostLevelSubmit() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        verified := false

        err := render(ctx, http.StatusOK, components.LevelFeedbackHtml(verified))
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render page"})
        }
    }
}

