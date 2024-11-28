package main

import (
	"fmt"
	"os"
	// "prompt-game/external/openai"
	"prompt-game/internal"

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

	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	app := internal.Config{Router: router, ApiKey: os.Getenv("OPENAI_API_KEY")}

	router.Static("/static", "./static")

	app.Routes()

	router.Run(":8080")
}
