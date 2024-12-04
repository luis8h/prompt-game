package handlers

import (
	"net/http"
	"prompt-game/views"
	"prompt-game/views/pages/index"

	"github.com/gin-gonic/gin"
)

type IndexHandler struct {
}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

func (h *IndexHandler) IndexPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := render(ctx, http.StatusOK, views.Layout(index.IndexPage()))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render page"})
		}
	}
}
