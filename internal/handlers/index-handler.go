package handlers

import (
	"net/http"
	"prompt-game/views"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

type IndexHandler struct {
}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{
	}
}

// Helper function to render templ.Component with status
func render(ctx *gin.Context, status int, template templ.Component) error {
	ctx.Status(status)
	return template.Render(ctx.Request.Context(), ctx.Writer)
}

// index page
func (h *IndexHandler) IndexPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := render(ctx, http.StatusOK, views.Index())

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render page"})
		}
	}
}

