package handlers

import (
	"fmt"
	"net/http"

	"prompt-game/internal/utils"
	"prompt-game/views"
	"prompt-game/views/pages/result"

	"github.com/gin-gonic/gin"
)

type ResultHandler struct {
}

func NewResultHandler() *PromptHandler {
	return &PromptHandler{}
}

func (h *PromptHandler) GetResultRestart() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		SetCurrentLevel(ctx, 0)
		SetStoryId(ctx, 0)
		SetShowTask(ctx, false)
		SetWithStrategy(ctx, false)

		utils.GameLogger.PrintS(ctx, fmt.Sprintf("restarting game"))

		// redirect
		ctx.Writer.Header().Set("HX-Redirect", "/")
		ctx.Writer.Header().Set("HX-Trigger", "resetChatHistory")
		ctx.Status(http.StatusOK)
	}
}

func (h *PromptHandler) GetResult() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		render(ctx, http.StatusOK, views.Layout(result.ResultPage(), GetSessionId(ctx)))
	}
}
