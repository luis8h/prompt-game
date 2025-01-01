package stores

import (
	"prompt-game/internal/models"
)

var Levels = []models.Level{
	{
		Title: "The Grand Healing Potion",
		Description: `
			The king has been struck down by a mysterious illness, and only the Grand Healing Potion can save him. The royal doctor has found an ancient recipe, but it's riddled with cryptic instructions.
			The potion requires a precise balance of ingredients, and the total volume must be **48ml**.

			The following clues describe the amounts needed for each ingredient:

			1. The amount of **Dragon's Breath** is **twice the amount of Phoenix Feather**.
			2. The amount of **Unicorn Tears** is **1.5 times the amount of Dragon's Breath**.
			3. The amount of **Fairy Dust** is **3ml less than the amount of Unicorn Tears**.
			4. The amount of **Moonlight Dew** is **half the amount of Fairy Dust**.
			5. The amount of **Dragon's Breath** is **4ml more than the amount of Moonlight Dew**.
			6. The total amount of **Dragon's Breath**, **Phoenix Feather**, **Unicorn Tears**, **Fairy Dust**, and **Moonlight Dew** must add up to **48ml**.

			Using these clues, help the royal doctor figure out the exact amounts of each ingredient, ensuring the total potion volume is exactly **48ml**.
		`,
		Strategy: "Use the zero-shot chain of thought strategy to break down the problem step-by-step and solve the system of equations.",
	},

	{
		Title: "The Healing Brew",
		Description: `
			The king is gravely ill, and the royal doctor has discovered an ancient recipe for a potion that might save him.
			The recipe is cryptic, and only the correct balance of ingredients will produce a healing brew. The total volume of the potion must be exactly **36ml**, and the doctor has made the following notes:

			1. The amount of **Dragon's Breath** is **twice the amount of Phoenix Feather**.
			2. The amount of **Unicorn Tears** is **1.5 times the amount of Dragon's Breath**.
			3. The amount of **Dragon's Breath** is **3ml more than the amount of Fairy Dust**.
			4. The total amount of **Dragon's Breath**, **Phoenix Feather**, **Unicorn Tears**, and **Fairy Dust** must add up to **36ml**.
			5. The amount of **Phoenix Feather** is **equal to the amount of Fairy Dust**.

			Help the doctor determine the correct amounts of each ingredient, ensuring that the total volume is exactly **36ml**.
		`,
		Strategy: "Use the zero-shot chain of thought strategy to reason through the problem and solve the system of equations.",
	},
	{
		Title: "The Enigmatic Illness",
		Description: `
			The king has fallen ill, and the royal doctor is frantic. His cure rests on a legendary potion whose recipe has been passed down through cryptic notes.
			The potion must total exactly **36ml**, but a single misstep in the composition could have dire consequences.

			The notes include the following instructions about the potion's ingredients:
			1. The amount of **Dragon's Breath** must be exactly **twice the amount of Phoenix Feather**.
			2. The amount of **Unicorn Tears** must equal the **sum of Dragon's Breath and Phoenix Feather**.
			3. The ratio of Phoenix Feather to the total potion volume must be exactly **1:9**.

			Decipher the correct measurements for each ingredient using the guidance of the LLM, ensuring the potion's total volume is 36ml. Only a precise composition can save the king.
		`,
		Strategy: "Use the zero-shot chain of thought strategy to reason step by step and calculate the quantities based on the constraints.",
	},
	{
		Title: "The illness",
		Description: `
			Since a few days, the king doesn't feel very well. The royal doctor wants to cook a special potion which will help the king to heal.
			But the recipe is a bit messed up and its hard to tell how much he needs of each ingredient is required. A small mistake in the ingredient composition could create an oposite effect.
			Help him by using the llm to tell how much milliliters are needed of each ingredient if the total volume should be 36ml in the end.

			The recipe:
				- The amount of Dragons's Breath is twice the amount of Phoenix Feather.
				- The amount of Unicorn Tears is the sum of Dragon's Breath and Phoenix Feather.
		`,
		Strategy: "Use the zero shot chain of thought strategy",
	},
	{
		Title: "The Healing Brew",
		Description: `
			The king has fallen ill, and the royal doctor needs to prepare a special potion. The recipe is ancient, and although the ingredients are known, the exact measurements are not.
			However, there are clues hidden within the instructions that will help you decipher the correct amounts.

			The total volume of the potion must be **36ml**, and the recipe contains the following conditions:

			1. The amount of **Dragon's Breath** is **three times the amount of Unicorn Tears**.
			2. The amount of **Phoenix Feather** is **half the amount of Dragon's Breath**.
			3. The total amount of **Dragon's Breath, Unicorn Tears, and Phoenix Feather** combined must be **36ml**.
			4. The amount of **Dragon's Breath** must also be **4ml more than Phoenix Feather**.

			Using these clues, help the doctor figure out the exact amounts of each ingredient. Make sure the potion totals exactly **36ml**.
		`,
		Strategy: "Use the zero-shot chain of thought strategy to break down the problem step-by-step and solve the system of equations.",
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
