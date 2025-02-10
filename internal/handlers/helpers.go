package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/invopop/ctxi18n"
)

func render(ctx *gin.Context, status int, template templ.Component) {
	ctx.Status(status)
	err := template.Render(ctx.Request.Context(), ctx.Writer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render template"})
	}
}

func getLocale(ctx *gin.Context) string {
	locale := "en"
	if cookie, err := ctx.Cookie("lang"); err == nil {
		locale = cookie
	}
	return locale
}

func GetTranslationContext(ctx *gin.Context) context.Context {
    lang := getLocale(ctx)
    transCtx, err := ctxi18n.WithLocale(ctx.Request.Context(), lang)
	if err != nil {
		fmt.Println("Error unmarshalling messages:", err)
	}
	return transCtx
}

func GetSessionId(ctx *gin.Context) string {
	session := sessions.Default(ctx)
	sessionId, ok := session.Get("sessionId").(string)
	if !ok {
		sessionId = "undefined session"
	}
	return sessionId
}

func SetStoryId(ctx *gin.Context, storyId int) int {
	session := sessions.Default(ctx)
	session.Set("storyId", storyId)
	session.Save()
	return storyId
}

func GetStoryId(ctx *gin.Context) int {
	session := sessions.Default(ctx)
	storyId, ok := session.Get("storyId").(int)
	if !ok {
		storyId = SetStoryId(ctx, 0)
	}
	return storyId
}

func SetCurrentLevel(ctx *gin.Context, levelId int) int {
	session := sessions.Default(ctx)
	session.Set("currentLevel", levelId)
	session.Save()
	return levelId
}

func GetCurrentLevel(ctx *gin.Context) int {
	session := sessions.Default(ctx)
	levelId, ok := session.Get("currentLevel").(int)
	if !ok {
		levelId = SetCurrentLevel(ctx, 0)
	}
	return levelId
}

func SetWithStrategy(ctx *gin.Context, withStrategy bool) bool {
	session := sessions.Default(ctx)
	session.Set("withStrategy", withStrategy)
	session.Save()
	return withStrategy
}

func GetWithStrategy(ctx *gin.Context) bool {
	session := sessions.Default(ctx)
	withStrategy, ok := session.Get("withStrategy").(bool)
	if !ok {
		withStrategy = SetWithStrategy(ctx, false)
	}
	return withStrategy
}

func SetShowTask(ctx *gin.Context, showTask bool) bool {
	session := sessions.Default(ctx)
	session.Set("showTask", showTask)
	session.Save()
	return showTask
}

func GetShowTask(ctx *gin.Context) bool {
	session := sessions.Default(ctx)
	showTask, ok := session.Get("showTask").(bool)
	if !ok {
		showTask = SetWithStrategy(ctx, false)
	}
	return showTask
}
