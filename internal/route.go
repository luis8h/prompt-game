package internal

import (
	"github.com/gin-gonic/gin"
    "prompt-game/internal/handlers"
)

type Config struct {
    Router *gin.Engine
}

func (app *Config) Routes() {
    indexHandler := handlers.NewIndexHandler()

    //views
    app.Router.GET("/", indexHandler.IndexPage())
}
