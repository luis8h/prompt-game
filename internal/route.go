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
    app.Router.POST("/prompt/user", promptHandler.PostPromptUser())
    app.Router.POST("/prompt/assistant", promptHandler.PostPromptAssistant())
    app.Router.POST("/message/assistant", promptHandler.PostMessageAssistant())
    app.Router.POST("/message/user", promptHandler.PostMessageUser())
    app.Router.POST("/prompt", promptHandler.PostPrompt())
    app.Router.DELETE("/prompt/reset", promptHandler.DeletePromptReset())

    // level
    levelHandler := handlers.NewLevelHandler(app.ApiKey)
    app.Router.GET("/level/submit/:levelId", levelHandler.GetLevelSubmit())
}
