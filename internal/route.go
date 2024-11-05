package internal

import (
	"github.com/gin-gonic/gin"
    "prompt-game/internal/handlers"
)

type Config struct {
    Router *gin.Engine
}

func (app *Config) Routes() {
    // index
    indexHandler := handlers.NewIndexHandler()
    app.Router.GET("/", indexHandler.IndexPage())

    // prompt
    promptHandler := handlers.NewPromptHandler()
    app.Router.POST("/prompt", promptHandler.PostPrompt())
}
