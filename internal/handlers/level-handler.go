package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"prompt-game/external/openai"
	"prompt-game/internal/models"
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
        valid, err := h.isValidStrategy(messageSlice)

        err = render(ctx, http.StatusOK, components.LevelFeedbackHtml(valid))
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render page"})
        }
    }
}

func (h *PromptHandler) isValidStrategy(messages []openai.Message) (bool, error) {
    // TODO: use the right prompt
    prompt := "hallo"
    strResponse, err := h.api.GetAnswer(prompt, []openai.Message{})
    if err != nil {
        return false, err
    }

    var jsonResponse models.VerificationResponse
    if err := json.Unmarshal(([]byte(strResponse)), &jsonResponse); err != nil {
        fmt.Println("Error unmarshalling api response:", err)
        return false, err
    }

    return jsonResponse.Verified, nil
}

