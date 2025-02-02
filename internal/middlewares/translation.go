package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/invopop/ctxi18n"
)

func I18nMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Default language fallback
		lang := "de"

		// Check for language stored in the cookie
		if cookie, err := c.Cookie("lang"); err == nil {
			lang = cookie
		}

		// Create a context with the desired language
        ctx, err := ctxi18n.WithLocale(c.Request.Context(), lang)
        if err != nil {
			log.Printf("error setting locale: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "error setting locale"})
			c.Abort()
			return
		}

		// Pass the modified context to the Gin context
		c.Request = c.Request.WithContext(ctx)

		// Proceed with the request
		c.Next()
	}
}

