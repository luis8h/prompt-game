package internal

import (
	"github.com/gin-gonic/gin"
    "prompt-game/internal/handlers"
)

type Config struct {
    Router *gin.Engine
    ApiKey string
}

func (app *Config) Routes() {
    // index
    indexHandler := handlers.NewIndexHandler()
    app.Router.GET("/", indexHandler.IndexPage())

    // prompt
    promptHandler := handlers.NewPromptHandler(app.ApiKey)
    app.Router.POST("/prompt", promptHandler.PostPrompt())
}
