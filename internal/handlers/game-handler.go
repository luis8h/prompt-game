package handlers

import (
	"net/http"
	"prompt-game/internal/stores"
	"prompt-game/internal/utils"
	"prompt-game/views"
	"prompt-game/views/pages/game"
	"prompt-game/views/pages/result"

	"github.com/gin-gonic/gin"
)

type GameHandler struct {
}

func NewGameHandler() *GameHandler {
	return &GameHandler{}
}

func (h *GameHandler) GetGamePage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		locale := getLocale(ctx)
		currentLevel := GetCurrentLevel(ctx)
		storyId := GetStoryId(ctx)
		showTask := GetShowTask(ctx)
		withStrategy := GetWithStrategy(ctx)

		utils.GameLogger.PrintS(ctx, "entered game")

		// check for results
		if currentLevel >= stores.GetLevelCount() {
			render(ctx, http.StatusOK, views.Layout(result.ResultPage(), GetSessionId(ctx)))
			return
		}

		if (len(stores.GetLevel(currentLevel, locale).Story) <= 1) {
			showTask = SetShowTask(ctx, true)
		}

		// render page
		render(ctx, http.StatusOK, views.Layout(game.GamePage(stores.GetLevel(currentLevel, locale), withStrategy, currentLevel, storyId, showTask), GetSessionId(ctx)))
	}
}
