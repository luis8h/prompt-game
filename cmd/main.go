package main

import (
    "prompt-game/internal"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    app := internal.Config{Router: router}

    app.Routes()

    router.Run(":8080")
}
