package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SessionIdMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		if (session.Get("sessionId") == nil) {
			newId := uuid.New().String()
			session.Set("sessionId", newId)
			session.Save()
		}
		ctx.Next()
	}
}

