package stores

import (
	"prompt-game/internal/models"
)

// var levels models.TranslatedLevels = models.TranslatedLevels{
//     {
//         "en": {Title: "Start", Description: "Start the game", Strategy: "Basic Strategy"},
//         "de": {Title: "Start", Description: "Starte das Spiel", Strategy: "Grundstrategie"},
//     },
//     {
//         "en": {Title: "Level 2", Description: "Explore the cave", Strategy: "Use a torch"},
//         "de": {Title: "Level 2", Description: "Erkunde die Höhle", Strategy: "Benutze eine Fackel"},
//     },
// }

func GetLevelCount() int {
	return len(levels)
}

func GetLevel(levelID int, langCode string) models.Level {
	if level, exists := levels[levelID][langCode]; exists {
		return level
	}
	// Fallback to English
	return levels[levelID]["en"]
}

var levels models.TranslatedLevels = models.TranslatedLevels{
	{
		"en": {
			Title: "Introduce yourself",
			Description: "**The king wants to have** an introduction to a person before he meets somebody. This should be a text which should be written in formal medieval language.\n\n" +
				"He only trusts persons who can show that they are aware of good language.\n\n" +
				"Give the oracle some details about yourself and let it write this introduction letter. (you do not have to use your real identity, just make up something)\n\n",
			Strategy: "RTF framework (role - task - format)",
		},
	},
	// the king could accept the invitation, but to proof, that the oracle really could help the king, he asks to solve the riddle of the caesar cipher ->
	{
		"en": {
			Title: "Caesar Cipher",
			Description: `
			Secret Book with nonsense. Librarian thinks it might be encoded with caesar cipher, but he is not able to get the right shift count.
			He gives you the following sentence and asks if you can find out which shift was used using the oracle.
			"Dtz hwfhpji ymj jshwduynts. Mfaj kzs bnym ymnx afqzfgqj pstbqjilj."
		`,
			Strategy: "Generated Knowledge. First let the llm generate knowledge about a topic and then tell it the task.",
		},
	},

	{
		"en": {
			Title: "Caesar Cipher 2",
			Description: `
			Take shift from last Level and use it to decode a long text.
			Try decoding the following text by just asking the oracle to decode it. But as the text is pretty long, the result might be a real mess.
			But there is a trick for such long, but repetitive tasks. Just ask the oracle to give you the program code to solve this problem in javascript.
			You can try out the result right here in the browser, by clicking f12 and entering the console. If you paste the generted code here, you can see the result.
		`,
			Strategy: "Asking the llm to write code to solve the problem instead of just genrating an answer.",
		},
	},
	{
		"en": {
			Title: "The mystirious Recipe",
			Description: `
			Since a few days, the king doesn't feel very well. The royal doctor wants to cook a special potion which will help the king to heal.
			But the recipe is a bit messed up and its hard to tell which amount of each ingredient is required. A small mistake in the ingredient composition could create an oposite effect.
			Help him by using the llm to tell how much milliliters are needed of each ingredient.

			The recipe:
			Behold the mystical concoction that awaits your alchemy skills—an enchanting potion infused with the rarest of ingredients. Begin with Dragon's Breath, a potent essence that is twice the amount of Phoenix Feather. This fiery breath will ignite the very heart of your potion. The second key ingredient, Phoenix Feather, is half the amount of the Dragon's Breath, yet crucial for bringing balance to the brew. Then, introduce the Unicorn Tears, a magical and ethereal addition to the mix. Their amount is the sum of Dragon's Breath and Phoenix Feather combined—an elixir of pure purity. Weave the delicate magic of the Elven realm into your potion with Elven Essence, which amounts to half of the Unicorn Tears, creating a soft but powerful undercurrent in the brew. Now, blend the power of the Dragon's Breath and Elven Essence—when their combined essence is divided by two, you’ll discover the perfect amount of Moonstone Dust needed to complete your creation. After carefully adding each ingredient, your potion will reach a total volume of 44.4 ml, a precise and harmonious blend of magic and mystery. With each step, you draw closer to unlocking the ancient magic of this unique elixir.
		`,
			Strategy: "Zero shot chain of thought strategy. Tell the llm to think/go step by step at the end of your prompt.",
		},
	},
}
