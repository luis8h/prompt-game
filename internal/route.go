package internal

import (
	"prompt-game/internal/handlers"
	"prompt-game/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type Config struct {
    Router *gin.Engine
    ApiKey string
}

// func (app *Config) Routes() {
//     // index
//     indexHandler := handlers.NewIndexHandler()
//     app.Router.GET("/", indexHandler.IndexPage())
//
//     // game
//     gameHandler := handlers.NewGameHandler()
//     app.Router.GET("/game", gameHandler.GetGamePage())
//
//     // prompt
//     promptHandler := handlers.NewPromptHandler(app.ApiKey)
//     app.Router.POST("/message/assistant", promptHandler.PostMessageAssistant())
//     app.Router.POST("/message/user", promptHandler.PostMessageUser())
//     app.Router.POST("/prompt", promptHandler.PostPrompt())
//     app.Router.POST("/history", promptHandler.PostHistory())
//
//     // level
//     levelHandler := handlers.NewLevelHandler(app.ApiKey)
//     app.Router.POST("/level/submit", levelHandler.PostLevelSubmit())
//
//     // result
//     resultHandler := handlers.NewResultHandler()
//     app.Router.GET("/result", resultHandler.GetResult())
//     app.Router.GET("/result/restart", resultHandler.GetResultRestart())
// }

func (app *Config) Routes() {
    // Create a language-based group
    langGroup := app.Router.Group("/:lang") // Define a dynamic language prefix

    // Add the language middleware to the group
    langGroup.Use(middlewares.NewLanguageMiddleware())

    // Define all routes under this group
    {
        // index
        indexHandler := handlers.NewIndexHandler()
        langGroup.GET("/", indexHandler.IndexPage())

        // game
        gameHandler := handlers.NewGameHandler()
        langGroup.GET("/game", gameHandler.GetGamePage())

        // prompt
        promptHandler := handlers.NewPromptHandler(app.ApiKey)
        langGroup.POST("/message/assistant", promptHandler.PostMessageAssistant())
        langGroup.POST("/message/user", promptHandler.PostMessageUser())
        langGroup.POST("/prompt", promptHandler.PostPrompt())
        langGroup.POST("/history", promptHandler.PostHistory())

        // level
        levelHandler := handlers.NewLevelHandler(app.ApiKey)
        langGroup.POST("/level/submit", levelHandler.PostLevelSubmit())

        // result
        resultHandler := handlers.NewResultHandler()
        langGroup.GET("/result", resultHandler.GetResult())
        langGroup.GET("/result/restart", resultHandler.GetResultRestart())
    }
}

