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

type PromptHandler struct {
	api      *openai.Api
}

func NewPromptHandler(apiKey string) *PromptHandler {
	return &PromptHandler{
		api:      openai.NewApi(apiKey),
	}
}

func (h *PromptHandler) DeletePromptReset() gin.HandlerFunc {
    return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
        emptyMessages := []components.Message{}
        messageBytes, _ := json.Marshal(emptyMessages)
        session.Set("messages", messageBytes)
        session.Save()

        err := render(ctx, http.StatusOK, components.ChatHistory(emptyMessages))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render page"})
		}
    }
}

func (h *PromptHandler) PostPromptUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		message := ctx.PostForm("prompt_input")

		if message == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "empty request body"})
			return
		}

		viewMessage := components.Message{Role: "user", Content: message}

		err := render(ctx, http.StatusOK, components.ChatMessage(viewMessage))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render page"})
		}
	}
}

func (h *PromptHandler) PostPromptAssistant() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		message := ctx.PostForm("prompt_input")
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

        // check wether the given message is valid
		if message == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "empty request body"})
			return
		}

        // create new user message
		newMessage := openai.Message{Role: "user", Content: message}

        // openai api call
		resp, err := h.api.GetAnswer(message, messageSlice)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error in openai api request"})
			return
		}

        // create new assistant message
		assistantMessage := openai.Message{Role: "assistant", Content: resp}

        // save messages to session
        messageSlice = append(messageSlice, newMessage, assistantMessage)
		messageBytes, _ := json.Marshal(messageSlice)
		session.Set("messages", messageBytes)
		session.Save()

        // render content
		viewMessage := components.Message{Role: assistantMessage.Role, Content: assistantMessage.Content}
		err = render(ctx, http.StatusOK, components.ChatMessage(viewMessage))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render page"})
		}
	}
}
