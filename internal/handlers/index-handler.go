package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"prompt-game/views"
	"prompt-game/views/components"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type IndexHandler struct {
}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

func (h *IndexHandler) IndexPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)

		// get messages from session
		// TODO: put this into an extra function in the model file to use in different functions (maybe a generic function)
		messageData, ok := session.Get("messages").([]byte)
		if !ok {
			messageSlice := []components.Message{}
			emptyMessageData, err := json.Marshal(messageSlice)
			if err != nil {
				fmt.Println("Error initializing messages:", err)
				return
			}
			session.Set("messages", emptyMessageData)
            session.Save()
			messageData = emptyMessageData
		}
		var messageSlice []components.Message
		if err := json.Unmarshal(messageData, &messageSlice); err != nil {
			fmt.Println("Error unmarshalling messages:", err)
			return
		}

        // get current level
        currentLevel, ok := session.Get("currentLevel").(int)
        if !ok {
            currentLevel = 0
            session.Set("currentLevel", 0)
        }

		// render page
		err := render(ctx, http.StatusOK, views.Index(messageSlice, currentLevel))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render page"})
		}
	}
}
