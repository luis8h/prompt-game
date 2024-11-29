package handlers

import (
	"net/http"
	"prompt-game/internal/stores"
	"prompt-game/views"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type IndexHandler struct {
}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

func (h *IndexHandler) IndexPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)

        // get current level
        currentLevel, ok := session.Get("currentLevel").(int)
        if !ok {
            currentLevel = 0
            session.Set("currentLevel", 0)
            session.Save()
        }

		// render page
		err := render(ctx, http.StatusOK, views.Index(stores.Levels[currentLevel]))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render page"})
		}
	}
}
