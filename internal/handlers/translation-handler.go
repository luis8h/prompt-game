package handlers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type TranslationHandler struct {
}

func NewTranslationHandler() *TranslationHandler {
    return &TranslationHandler{}
}

func (h *TranslationHandler) PostLanguage() gin.HandlerFunc {
    return func(ctx *gin.Context) {
		lang := ctx.PostForm("lang")
		ctx.SetCookie("lang", lang, 3600*24*30, "/", "", false, true)
		ctx.JSON(http.StatusOK, gin.H{"message": "Language updated to " + lang})
    }
}
