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

    // translation
    translationHandler := handlers.NewTranslationHandler()
    app.Router.POST("/language", translationHandler.PostLanguage())

    // game
    gameHandler := handlers.NewGameHandler()
    app.Router.GET("/game", gameHandler.GetGamePage())

    // prompt
    promptHandler := handlers.NewPromptHandler(app.ApiKey)
    app.Router.POST("/message/assistant", promptHandler.PostMessageAssistant())
    app.Router.POST("/message/user", promptHandler.PostMessageUser())
    app.Router.POST("/prompt", promptHandler.PostPrompt())
    app.Router.POST("/history", promptHandler.PostHistory())

    // level
    levelHandler := handlers.NewLevelHandler(app.ApiKey)
    app.Router.POST("/level/nexta", levelHandler.PostLevelNextA())
    app.Router.POST("/level/nextb", levelHandler.PostLevelNextB())
    app.Router.GET("/level/story/next", levelHandler.GetLevelStoryNext())
    app.Router.GET("/level/story/prev", levelHandler.GetLevelStoryPrev())
	app.Router.GET("/level/set/:id", levelHandler.SetLevel())
    app.Router.GET("/level", levelHandler.GetLevel())

    // result
    resultHandler := handlers.NewResultHandler()
    app.Router.GET("/result", resultHandler.GetResult())
    app.Router.GET("/result/restart", resultHandler.GetResultRestart())
}
