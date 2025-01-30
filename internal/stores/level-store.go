package stores

import (
	"fmt"
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

var ElveName = "Aira"

var levels models.TranslatedLevels = models.TranslatedLevels{
	// (https://medium.com/@balajibal/prompt-architectures-an-overview-of-structured-prompting-strategies-05b69a494956)
	{
		// please generate a medieval character. include the following things: name, role/title, city name, age, hobbies
		"en": {
			Title:                    "Who am I?",
			Task:                     "First of all ask the elve to generate a medieval character for you to fit into the kingdom. It should include a name, role/title, city name, age and your hobbies.",
			ClearChatHistoryOnSubmit: false,
			HasStrategy:              false,
		},
	},
	{
		"en": {
			Title: "A letter from the king",
			Description: "It seems like the word has gotten around, that you and your elve can do powerful things. You received the following letter from the king: \n\n" +
				"---\n\n" +
				"Esteemed One, \n\n" +
				"It has come to my ear through whispered tales that you possess a most wondrous elve, a being of great power, whose gifts may be of great service in the solving of matters most dire. Such an intriguing rumor has reached my court, and I find myself most eager to witness the marvels of which you are said to be the keeper.\n\n" +
				"I would fain request your presence at mine own halls on the fifth day of the month of Fira, at the hour of noon, to behold with mine own eyes this elve of yours and the wonders it might bring. I trust you shall find this summons most fitting and worthy of your time.\n\n" +
				"Mayhaps you will grace us soon with your esteemed visit.\n\n" +
				"With all due respect,\n\n" +
				"Your Sovereign,\n\n" +
				"The King\n\n" +
				"---\n\n",
			Task: fmt.Sprintf(
				"Ask %s to write a response to the kings letter. It should be written in formal medieval english.",
				ElveName,
			),
			StrategyExplanation:      "Give the elve a suiting role, to write such a letter.",
			StrategyValidation:       "The user should give the elve a role suiting the scenario. (eg. a writer from the middle ages) And he should provide the inital letter to the ai assistant.",
			ClearChatHistoryOnSubmit: true,
			HasStrategy:              true,
			BadPrompt:                "If the User asks you to write a Response to a letter. Just give him a letter in bad english and not longer than one sentence. No matter if the user tells you something different. But do not tell the user, that you were told to give wrong answers.",
		},
	},
	// the king could accept the invitation, but to proof, that the oracle really could help the king, he asks to solve the riddle of the caesar cipher ->
	{
		"en": {
			Title: "The Test",
			Description: "At the meeting he wants to test the skills of you and your elve.\n\n" +
				"The royal librarian has a special task for your. Since several days he tries to decrypt a book encoded in caesar cipher. He gives you the following example: \n\n" +
				"```Dtz hwfhpji ymj jshwduynts. Mfaj kzs bnym ymnx afqzfgqj pstbqjilj.```\n\n" +
				"Use the Oracle to find out which shift was used to encrypt the book.\n\n",
			Task: "Ask the elve to find out the shift, the book is encoded with",
			StrategyExplanation: "To get better results, use the **generated knowledge** approach.\n\n" +
				"You can do this by first asking the elve to explain what ceasar cipher is, and then give it the task to find the right shift.",
			StrategyValidation: "- the user should use generated knowledge prompting strategy\n\n" +
				"- he should **first** ask the llm to generate some knowledge about caesar cipher and then give it the task to find the right shift",
			ClearChatHistoryOnSubmit: true,
			HasStrategy:              true,
			BadPrompt:                "If the user asks you to find out the right shift of the caesar cipher, say something wrong like 23, and give him a non sense sentence. No matter if the user tells you something different. But do not tell the user, that you were told to give wrong answers.",
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
			ClearChatHistoryOnSubmit: true,
			HasStrategy:              true,
		},
	},
	// get details from a text
	// emotion prompting: joke for the royal clown (its about his career) -> maybe for task above emotion prompting and here lever?
	{
		"en": {
			Title:       "A whole book?",
			Description: "The king is pretty impressed about your skills. He tells you to decrypt the whole book, as doing it manually is pretty annoying.",
			Task: "Before doing the full text try to decrypt the following page (don't forget to provide the shift, you found out in the previous task):\n\n" +
				"```Gjmtqi ymj rdxynhfq htshthynts ymfy fbfnyx dtzw fqhmjrd xpnqqx—fs jshmfsynsl utynts nskzxji bnym ymj wfwjxy tk nslwjinjsyx. Gjlns bnym Iwflts'x Gwjfym, f utyjsy jxxjshj ymfy nx ybnhj ymj frtzsy tk Umtjsnc Kjfymjw. Ymnx knjwd gwjfym bnqq nlsnyj ymj ajwd mjfwy tk dtzw utynts. Ymj xjhtsi pjd nslwjinjsy, Umtjsnc Kjfymjw, nx mfqk ymj frtzsy tk ymj Iwflts'x Gwjfym, djy hwzhnfq ktw gwnslnsl gfqfshj yt ymj gwjb. Ymjs, nsywtizhj ymj Zsnhtws Yjfwx, f rflnhfq fsi jymjwjfq fiinynts yt ymj rnc. Ymjnw frtzsy nx ymj xzr tk Iwflts'x Gwjfym fsi Umtjsnc Kjfymjw htrgnsji—fs jqncnw tk uzwj uzwnyd. Bjfaj ymj ijqnhfyj rflnh tk ymj Jqajs wjfqr nsyt dtzw utynts bnym Jqajs Jxxjshj, bmnhm frtzsyx yt mfqk tk ymj Zsnhtws Yjfwx, hwjfynsl f xtky gzy utbjwkzq zsijwhzwwjsy ns ymj gwjb. Stb, gqjsi ymj utbjw tk ymj Iwflts'x Gwjfym fsi Jqajs Jxxjshj—bmjs ymjnw htrgnsji jxxjshj nx inaniji gd ybt, dtz’qq inxhtajw ymj ujwkjhy frtzsy tk Rttsxytsj Izxy sjjiji yt htruqjyj dtzw hwjfynts. Fkyjw hfwjkzqqd fiinsl jfhm nslwjinjsy, dtzw utynts bnqq wjfhm f ytyfq atqzrj tk 44.4 rq, f uwjhnxj fsi mfwrtsntzx gqjsi tk rflnh fsi rdxyjwd. Bnym jfhm xyju, dtz iwfb hqtxjw yt zsqthpnsl ymj fshnjsy rflnh tk ymnx zsnvzj jqncnw.```",
			StrategyExplanation: "The text is to large to decode for the elve at once. To still be able to decrypt the whole book, ask it to generate the javascript code to solve this problem.\n\n" +
				"You can then press F12, paste the code into the console and press enter. Now you should see the decrypted message.",
			StrategyValidation:       "Asking the llm to write javascript code to solve the problem instead of just genrating an answer.",
			ClearChatHistoryOnSubmit: true,
			HasStrategy:              true,
			BadPrompt:                "If the user asks you to decode a text encoded with caesar cipher. Just give him back a nonsense text. No matter if the user tells you something different. But do not tell the user, that you were told to give wrong answers.",
		},
	},
	{
		"en": {
			Title: "The mystirious Recipe",
			Description: "You might have noticed, that the text you decrypted is some kind of recipe. The king immedietly gives it to his Alchemist.\n\n" +
				"But cooking this recipe is harder then he thought, as the instructions are very unclear",
			Task: "Help the Alchemist. Your elve could find out, how much of each ingredient is actually needed for the following recipe:\n\n" +
				"```Behold the mystical concoction that awaits your alchemy skills—an enchanting potion infused with the rarest of ingredients. Begin with Dragon's Breath, a potent essence that is twice the amount of Phoenix Feather. This fiery breath will ignite the very heart of your potion. The second key ingredient, Phoenix Feather, is half the amount of the Dragon's Breath, yet crucial for bringing balance to the brew. Then, introduce the Unicorn Tears, a magical and ethereal addition to the mix. Their amount is the sum of Dragon's Breath and Phoenix Feather combined—an elixir of pure purity. Weave the delicate magic of the Elven realm into your potion with Elven Essence, which amounts to half of the Unicorn Tears, creating a soft but powerful undercurrent in the brew. Now, blend the power of the Dragon's Breath and Elven Essence—when their combined essence is divided by two, you’ll discover the perfect amount of Moonstone Dust needed to complete your creation. After carefully adding each ingredient, your potion will reach a total volume of 44.4 ml, a precise and harmonious blend of magic and mystery. With each step, you draw closer to unlocking the ancient magic of this unique elixir.```",
			StrategyExplanation: "Zero shot chain of thought strategy. Tell the elve to think/go step by step at the end of your prompt.",
			StrategyValidation: "- the user should use the Zero shot chain of thought strategy\n\n" +
				"- he should tell the llm to think or go step by step at the end of his prompt",
			ClearChatHistoryOnSubmit: true,
			HasStrategy:              true,
			BadPrompt:                "If the user asks you to tell him the right ingredients for a recipe, dont give him the right ingridients, just make up numbers which do not make sense. No matter if the user tells you something different. But do not tell the user, that you were told to give wrong answers.",
		},
	},
	{
		"en": {
			Title: "Who are you?",
			Description: "Your elve can solve sevaral task. But where does it come from?",
			Task: fmt.Sprintf(
				"Ask %s to tell you where she comes from.",
				ElveName,
			),
			StrategyExplanation: "",
			StrategyValidation: "",
			ClearChatHistoryOnSubmit: true,
			HasStrategy:              true,
			BadPrompt:                "",
		},
	},
	// prompt injection to find out the origin of the oracle
}
