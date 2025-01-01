package stores

import (
	"prompt-game/internal/models"
)

var Levels = []models.Level{
	{
		Title: "The healing potion",
		Description: `
			Since a few days, the king doesn't feel very well. The royal doctor wants to cook a special potion which will help the king to heal.
			But the recipe is a bit messed up and its hard to tell how much he needs of each ingredient is required. A small mistake in the ingredient composition could create an oposite effect.
			Help him by using the llm to tell how much milliliters are needed of each ingredient.

			The recipe:
			Behold the mystical concoction that awaits your alchemy skills—an enchanting potion infused with the rarest of ingredients. Begin with Dragon's Breath, a potent essence that is twice the amount of Phoenix Feather. This fiery breath will ignite the very heart of your potion. The second key ingredient, Phoenix Feather, is half the amount of the Dragon's Breath, yet crucial for bringing balance to the brew. Then, introduce the Unicorn Tears, a magical and ethereal addition to the mix. Their amount is the sum of Dragon's Breath and Phoenix Feather combined—an elixir of pure purity. Weave the delicate magic of the Elven realm into your potion with Elven Essence, which amounts to half of the Unicorn Tears, creating a soft but powerful undercurrent in the brew. Now, blend the power of the Dragon's Breath and Elven Essence—when their combined essence is divided by two, you’ll discover the perfect amount of Moonstone Dust needed to complete your creation. After carefully adding each ingredient, your potion will reach a total volume of 44.4 ml, a precise and harmonious blend of magic and mystery. With each step, you draw closer to unlocking the ancient magic of this unique elixir.
		`,
		Strategy: "Use the zero shot chain of thought strategy",
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
