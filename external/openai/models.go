package openai

// OpenAIRequest represents the request body to OpenAI API
type OpenAIRequest struct {
	Model    string   `json:"model"`
	Messages []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
}

// Message represents a single message in the conversation
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// OpenAIResponse represents the API response from OpenAI
type OpenAIResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Choices []struct {
		Index   int    `json:"index"`
		Message Message `json:"message"`
	} `json:"choices"`
}
