package models

type Level struct {
	Title       string
	Description string
	Strategy    string
}

type VerificationResponse struct {
	Verified bool `json:"verified"`
}
