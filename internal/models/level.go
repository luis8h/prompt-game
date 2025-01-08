package models

type TranslatedLevels []map[string]Level

type Level struct {
	Title       string
	Description string
    StrategyExplanation string
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
