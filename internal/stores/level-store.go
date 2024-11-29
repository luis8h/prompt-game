package stores

import (
	"prompt-game/internal/models"
)

var Levels = []models.Level{
	{
		Title:       "Sample level",
		Description: "Use the llm to write a poem in the style of the 19th century. To get better results you should give the llm a role first.",
		Strategy:    "The user should use the Strategy called Role-Prompting. That means he should give the ai assistant a role which suits his needs.",
	},
	{
		Title:       "Level 2",
		Description: "Use the llm to write a poem in the style of the 19th century. To get better results you should give the llm a role first.",
		Strategy:    "The user should use the Strategy called Role-Prompting. That means he should give the ai assistant a role which suits his needs.",
	},
	{
		Title:       "Level 3",
		Description: "Use the llm to write a poem in the style of the 19th century. To get better results you should give the llm a role first.",
		Strategy:    "The user should use the Strategy called Role-Prompting. That means he should give the ai assistant a role which suits his needs.",
	},
}
