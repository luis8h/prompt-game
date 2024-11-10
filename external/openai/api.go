package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Api struct {
	apiKey string
}

func NewApi(apiKey string) *Api {
	a := &Api{
		apiKey: apiKey,
	}
	return a
}

func (a *Api) validateStrategy(prompt string) (bool, error) {
    a.Get(prompt, []Message{})

	return false, nil
}

func (a *Api) Get(content string, messages []Message) (string, error) {
	// add new message to history
	newMessage := Message{Role: "user", Content: content}
	messages = append(messages, newMessage)

	requestBody := OpenAIRequest{
		Model:    "gpt-4o-mini",
		Messages: messages,
	}
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %v", err)
	}

	// create request
	apiURL := "https://api.openai.com/v1/chat/completions"
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+a.apiKey)

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	// Read and handle the response
	body, err := io.ReadAll(resp.Body) // Use io.ReadAll instead of ioutil.ReadAll
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, body)
	}

	// Parse JSON response
	var apiResponse OpenAIResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal response body: %v", err)
	}

	// Check if there are choices returned
	if len(apiResponse.Choices) == 0 {
		return "", fmt.Errorf("no response choices found")
	}

	// Return the content of the first choice
	return apiResponse.Choices[0].Message.Content, nil
}
