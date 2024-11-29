package models

type Level struct {
	Title       string
	Description string
	Strategy    string
}

type LevelValidation struct {
	Strategy bool
	Answer   bool
	Ignore   bool
}

type VerificationResponse struct {
	Verified bool `json:"verified"`
}
