package handlers

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func render(ctx *gin.Context, status int, template templ.Component) error {
	ctx.Status(status)
	return template.Render(ctx.Request.Context(), ctx.Writer)
}

func getLocale(ctx *gin.Context) string {
    locale := "en"
    if cookie, err := ctx.Cookie("lang"); err == nil {
        locale = cookie
    }
    return locale
}
