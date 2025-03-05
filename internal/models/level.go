package models

type TranslatedLevels []map[string]Level

type Level struct {
	Title                    string
	Story                    []Speechbubble
	Task                     string
	TaskValidation           string
	StrategyExplanation      string
	StrategyValidation       string
	ClearChatHistoryOnSubmit bool
	HasStrategy              bool
	BadPrompt                string
	GoodPrompt               string
	InfoText                 string
}

type Speechbubble struct {
	Character Character
	Text      string
}

type Character struct {
	Name       string
	Profession string
	Imgs       []string
}

type LevelValidation struct {
	Strategy bool
	Answer   bool
	Ignore   bool
}

type VerificationResponse struct {
	Verified bool `json:"verified"`
}
