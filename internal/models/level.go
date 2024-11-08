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
    Description string
    Technique string
}

var sampleLevel = Level{
    Description: "Use the large language model to write a poem about pigs. To get better results, give the large language model a role.",
    Technique: "The User should give the model he is talking to a role. In this context, the role should be something like a writer.",
}
