package middlewares

import (
	"strings"
	"net/http"
	"log"
	"github.com/invopop/ctxi18n"
	"github.com/gin-gonic/gin"
)

func NewLanguageMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := "en" // Default language
		pathSegments := strings.Split(c.Request.URL.Path, "/")
		if len(pathSegments) > 1 {
			lang = pathSegments[1]
		}

		ctx, err := ctxi18n.WithLocale(c.Request.Context(), lang)
		if err != nil {
			log.Printf("error setting locale: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "error setting locale"})
			c.Abort()
			return
		}

		// Pass the updated context to Gin
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
