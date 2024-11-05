package main

import (
	// "fmt"
	"log"
	// "os"
	// "prompt-game/external/openai"
    "prompt-game/internal"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Initialize function to load environment variables
func initEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on system environment variables.")
	}
}

func main() {
    initEnv()

    // apiKey := os.Getenv("OPENAI_API_KEY")
    // testprompt := "hallo wie geht es dir?"
    // api := openai.NewApi(apiKey)
    // resp, err := api.Get(testprompt)
    //
    // if (err != nil) {
    //     fmt.Errorf("request failed: %v", err)
    // }
    // fmt.Println(resp)


    router := gin.Default()
    app := internal.Config{Router: router}

    router.Static("/static", "./static")

    app.Routes()

    router.Run(":8080")
}
