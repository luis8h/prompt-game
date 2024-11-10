package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"prompt-game/external/openai"
	"prompt-game/views/components"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LevelHandler struct {
	api      *openai.Api
}

func NewLevelHandler(apiKey string) *PromptHandler {
	return &PromptHandler{
		api:      openai.NewApi(apiKey),
	}
}

func (h *PromptHandler) GetLevelSubmit() gin.HandlerFunc {
    return func(ctx *gin.Context) {
		session := sessions.Default(ctx)

        // get messages from session
		messageData, ok := session.Get("messages").([]byte)
		if !ok {
			fmt.Println("Error: messages are not stored as []byte")
			return
		}
		var messageSlice []openai.Message
		if err := json.Unmarshal(messageData, &messageSlice); err != nil {
			fmt.Println("Error unmarshalling messages:", err)
			return
		}

        // verify using openai api

        verified := false

        err := render(ctx, http.StatusOK, components.LevelFeedbackHtml(verified))
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render page"})
        }
    }
}

