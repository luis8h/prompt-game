package main

import (
	"fmt"
	"os"

	// "prompt-game/external/openai"
	"log"
	"prompt-game/internal"

	"prompt-game/locales"
	"prompt-game/internal/middlewares"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/invopop/ctxi18n"
	"github.com/joho/godotenv"
)

// func newLanguageMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		lang := "en" // Default language
// 		pathSegments := strings.Split(c.Request.URL.Path, "/")
// 		if len(pathSegments) > 1 {
// 			lang = pathSegments[1]
// 		}
//
// 		ctx, err := ctxi18n.WithLocale(c.Request.Context(), lang)
// 		if err != nil {
// 			log.Printf("error setting locale: %v", err)
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "error setting locale"})
// 			c.Abort()
// 			return
// 		}
//
// 		// Pass the updated context to Gin
// 		c.Request = c.Request.WithContext(ctx)
// 		c.Next()
// 	}
// }

// Initialize function to load environment variables
func initEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found, relying on system environment variables.")
	}
}

func main() {
	initEnv()

	if err := ctxi18n.Load(locales.Content); err != nil {
		log.Fatalf("error loading locales: %v", err)
	}

	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.Use(middlewares.NewLanguageMiddleware())


	// Group routes under a language prefix
	langGroup := router.Group("/:lang")
	{
		// Add language middleware to the group
		langGroup.Use(middlewares.NewLanguageMiddleware())

		// Serve static files under the language prefix
		langGroup.Static("/static", "./static")

		// Define application routes in the language group
		app := internal.Config{Router: router, ApiKey: os.Getenv("OPENAI_API_KEY")}
		app.Routes()
	}

	// Redirect root `/` to `/en`
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/en")
	})

	router.Run(":8080")
}
