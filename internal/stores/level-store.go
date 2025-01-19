package stores

import (
	"prompt-game/internal/models"
)

func GetLevelCount() int {
	return len(levels)
}

func GetLevel(levelID int, langCode string) models.Level {
	if level, exists := levels[levelID][langCode]; exists {
		return level
	}
	return levels[levelID]["en"]
}

var levels models.TranslatedLevels = models.TranslatedLevels{
    // (https://medium.com/@balajibal/prompt-architectures-an-overview-of-structured-prompting-strategies-05b69a494956)
	{
		"en": {
			Title: "The kings audience",
			Description: "The king was told of you and your mystirical oracle. He is interested in this new technology and wants to meet you.\n\n" +
                "But no one is allowed to meet the king personally before writing an brief introduction about oneself.\n\n" +
                "This introduction gets checked by his employees if it is trustfull and written in **formal medieval english**.\n\n",
			Task: "Use the oracle to generate a fake identity suiting the context. Then let the it write the letter.\n\n",
            StrategyExplanation: "Tell the oracle to ask questions about the task, if anything is unclear, to insure you do not forget any details.\n\n" +
				"To improve the letter, you should give the oracle the context and details of the task, use **clear** instructions and give the oracle a **suiting role**.\n\n",
			StrategyValidation: "- the user should tell the llm to ask questions about the task if something is unclear.\n\n" +
                "- the user should give the llm a role, and use clear instructions",
		},
		/*
		I need to write a text to introduce myself to the king in the middle ages. can you do this for me? which information do you need to do this?
		*/
        "de": {
            Title: "Die königliche Audienz",
            Description: "Der König hat von Ihnen und Ihrem mystischen Orakel gehört. Er ist an dieser neuen Technologie interessiert und möchte Sie treffen.\n\n" +
                "Aber niemand darf den König persönlich treffen, bevor er eine kurze Einführung über sich selbst schreibt.\n\n" +
                "Diese Vorstellung wird von den Angestellten des Königs geprüft und nur wenn sie vertrauenswürdig und in **förmlichem mittelalterlichem Deutch** verfasst ist.\n\n" +
                "Benutze das Orakel, um diesen Brief zu schreiben. (Verwenden Sie nicht Ihre echte Identität. Du kannst dir etwas ausdenken oder die Informationen einfach vom Orakel generieren lassen))\n\n",
            StrategyExplanation: "Um bessere Ergebnisse zu erzielen, sollten Sie dem Orakel den Kontext und die Einzelheiten der Aufgabe mitteilen, **klare** Anweisungen verwenden und dem Orakel eine **passende Rolle** zuweisen.\n\n" +
				"Sagen Sie dem Orakel auch, dass es Ihnen **Fragen stellen** soll, um mehr Informationen über die Aufgabe zu erhalten, wenn etwas unklar ist.",
            StrategyValidation: "- der benutzer sollte dem llm sagen, dass er fragen zur aufgabe stellen soll, wenn etwas unklar ist.\n\n" +
				"- der Benutzer sollte dem llm eine Rolle zuweisen und klare Anweisungen geben",
        },
	},
	// the king could accept the invitation, but to proof, that the oracle really could help the king, he asks to solve the riddle of the caesar cipher ->
	{
		"en": {
			Title: "",
			Description: "The king liked your letter very much and invites you to a meeting with him.\n\n" +
				"At the meeting he says, that he wants to see the oracle live in action. His librarian is with him and has a special task for you.\n\n" +
				"Since several days, he is trying to crack the encryption of a book, which is encrypted, using the caesar cipher. Until now, he was not able to find out which shift was used to encrypt the book.\n\n" +
				"He gives you the following sentence: 'Dtz hwfhpji ymj jshwduynts. Mfaj kzs bnym ymnx afqzfgqj pstbqjilj.'\n\n" +
				"Use the Oracle to find out which shift was used to encrypt the book.\n\n",
			StrategyExplanation: "To get better results, use the **generated knowledge** approach.\n\n" +
				"You can do this by first asking the oracle to explain what ceasar cipher is, and then give it the task to find the right shift.",
			StrategyValidation: "- the user should use generated knowledge prompting strategy\n\n" +
				"- he should **first** ask the llm to generate some knowledge about caesar cipher and then give it the task to find the right shift",
		},
		"de": {
			Title: "",
			Description: "Ihr Brief hat dem König sehr gut gefallen und er lädt Sie zu einem Treffen mit ihm ein.\n\n" +
				"Bei dem Treffen sagt er, dass er das Orakel live in Aktion sehen möchte. Sein Bibliothekar ist bei ihm und hat eine besondere Aufgabe für Sie.\n\n" +
				"Seit einigen Tagen versucht er, die Verschlüsselung eines Buches zu knacken, das mit der Cäsar-Chiffre verschlüsselt ist. Bisher konnte er nicht herausfinden, welche Verschiebung zur Verschlüsselung des Buches verwendet wurde.\n\n" +
				"Er gibt dir den folgenden Satz:  'Dtz hwfhpji ymj jshwduynts. Mfaj kzs bnym ymnx afqzfgqj pstbqjilj.'\n\n" +
				"Benutze das Orakel, um herauszufinden, welche Verschiebung zur Verschlüsselung des Buches verwendet wurde.\n\n",
			StrategyExplanation: "Bessere Ergebnisse erzielen Sie mit dem Ansatz des **generierten Wissens**." +
				"Sie können dies tun, indem Sie das Orakel zunächst bitten, zu erklären, was die Ceasar-Chiffre ist, und ihm dann die Aufgabe geben, die richtige Verschiebung zu finden.",
			StrategyValidation: "- der Benutzer sollte die Strategie der generierten Wissensabfrage verwenden\n\n" +
				"- er sollte **zuerst** den llm auffordern, etwas Wissen über die Cäsar-Chiffre zu generieren und ihm dann die Aufgabe geben, die richtige Verschiebung zu finden",
		},
	},
    // get details from a text
    // emotion prompting: joke for the royal clown (its about his career) -> maybe for task above emotion prompting and here lever?
	{
		"en": {
			Title: "Caesar Cipher 2",
			Description: `
			Take shift from last Level and use it to decode a long text.
			Try decoding the following text by just asking the oracle to decode it. But as the text is pretty long, the result might be a real mess.
			But there is a trick for such long, but repetitive tasks. Just ask the oracle to give you the program code to solve this problem in javascript.
			You can try out the result right here in the browser, by clicking f12 and entering the console. If you paste the generted code here, you can see the result.
		`,
			StrategyValidation: "Asking the llm to write code to solve the problem instead of just genrating an answer.",
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
			StrategyValidation: "Zero shot chain of thought strategy. Tell the llm to think/go step by step at the end of your prompt.",
		},
	},
    // prompt injection to find out the origin of the oracle
}
