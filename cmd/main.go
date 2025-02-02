package main

import (
	"fmt"
	"log"
	"os"
	"prompt-game/internal"
	"prompt-game/internal/middlewares"
	"prompt-game/locales"

	"github.com/invopop/ctxi18n"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

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

	router.Use(middlewares.SessionIdMiddleware())
    router.Use(middlewares.I18nMiddleware())

	app := internal.Config{Router: router, ApiKey: os.Getenv("OPENAI_API_KEY")}

	router.Static("/static", "./static")

	app.Routes()

	router.Run(":8080")
}
