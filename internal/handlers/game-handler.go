package handlers

import (
	"net/http"
	"prompt-game/internal/stores"
	"prompt-game/views"
	"prompt-game/views/pages/game"
	"prompt-game/views/pages/result"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type GameHandler struct {
}

func NewGameHandler() *GameHandler {
	return &GameHandler{}
}

func (h *GameHandler) GetGamePage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
        locale := getLocale(ctx)

		// get current level
		currentLevel, ok := session.Get("currentLevel").(int)
		if !ok {
			currentLevel = 0
			session.Set("currentLevel", 0)
			session.Save()
		}

        // check for results
        if currentLevel >= stores.GetLevelCount() {
            err := render(ctx, http.StatusOK, views.Layout(result.ResultPage()))
            if err != nil {
			    ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render page"})
            }
            return
        }

		// render page
		err := render(ctx, http.StatusOK, views.Layout(game.GamePage(stores.GetLevel(currentLevel, locale))))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render page"})
		}
	}
}
