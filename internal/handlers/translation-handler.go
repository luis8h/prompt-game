package handlers

import (
	"net/http"
	"prompt-game/internal/utils"

	"github.com/gin-gonic/gin"
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
		utils.GameLogger.PrintS(ctx, "change language to " + lang)
		ctx.JSON(http.StatusOK, gin.H{"message": "Language updated to " + lang})
    }
}
