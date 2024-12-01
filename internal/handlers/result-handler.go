package handlers

import (
	"net/http"

	"prompt-game/views"
	"prompt-game/views/pages/result"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type ResultHandler struct {
}

func NewResultHandler() *PromptHandler {
	return &PromptHandler{}
}

func (h *PromptHandler) GetResultRestart() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// reset levelid
		session := sessions.Default(ctx)
		session.Set("currentLevel", 0)
		session.Save()

		// redirect
		ctx.Writer.Header().Set("HX-Redirect", "/")
		ctx.Writer.Header().Set("HX-Trigger", "resetChatHistory")
		ctx.Status(http.StatusOK)
	}
}

func (h *PromptHandler) GetResult() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := render(ctx, http.StatusOK, views.Layout(result.ResultPage()))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render page"})
		}
	}
}
