package models

// TODO: add verification of task itself (maybe in a different request to seperate the concerns and get better answers)
// -> could maybe also include the result of the users prompt

const ScenarioPromptP1 = `
Background: I created a game, where the users can learn prompt engineering by solving different tasks using various prompting techniques.
Your task is to decide wheter in the following prompt, a specific prompt engineering technique was used.

prompt of the user: "
`

const ScenarioPromptP2 = `"

prompt technique description: "
`

const ScenarioPromptP3 = `"

task description: "
`

const ScenarioPromptP4 = `"

Reply in a json string and **nothing else** which has an attribute called verification. This attribute should contain true (if the technique was used **correctly**) and false (if not).
`

type Level struct {
	Title       string
	Description string
	Strategy    string
}

type VerificationResponse struct {
	Verified bool `json:"verified"`
}
