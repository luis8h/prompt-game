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

var ElfName = "AIra"

var KingdomName = "Prompt Kingdom"

var King models.Character = models.Character{
	Name:       "Henry",
	Profession: "King",
	Imgs: []string{
		"static/assets/characters/elf/elf_talking_1.png",
	},
}
var Elf models.Character = models.Character{
	Name:       "AIra",
	Profession: "Elf",
	Imgs: []string{
		"static/assets/characters/elf/elf_talking_1.png",
	},
}

var levels models.TranslatedLevels = models.TranslatedLevels{
	{
		"en": {
			Title: "A magical Elf",
			Story: []models.Speechbubble{
				{
					Character: Elf,
					Text:      fmt.Sprintf("Hi, I am %s, a magical Elf. I can help you do many things like writing and solving riddles.", Elf.Name),
				},
				{
					Character: Elf,
					Text:      fmt.Sprintf("You can talk to me by typing in the chat at the right side."),
				},
				{
					Character: Elf,
					Text:      fmt.Sprintf("Sometimes I am a bit chaotic. You need to talk to me the right way and explain things very clearly, so I can really help you."),
				},
				{
					Character: Elf,
					Text:      fmt.Sprintf("Try it out!"),
				},
			},
			Task:                     fmt.Sprintf("Try to talk to %s. For example, ask her to tell you some ideas for good medieval names.", Elf.Name),
			ClearChatHistoryOnSubmit: true,
			HasStrategy:              false,
			BadPrompt:                fmt.Sprintf(""),
		},
	},
	{
		"en": {
			Title: "A Letter",
			Story: []models.Speechbubble{
				{
					Character: King,
					Text:      fmt.Sprintf("Good day to thee! I am %s, the %s of the %s. I heard you possess magical powers.", King.Name, King.Profession, KingdomName),
				},
				{
					Character: King,
					Text:      fmt.Sprintf("Could you use those to help me out with a problem? I got an invitation from a neighboring kingdom to come for dinner to discuss important business relations, but I am busy on the day he proposed."),
				},
				{
					Character: King,
					Text:      fmt.Sprintf("I want to propose another day to him, but i am not a good writer at all. I do not want to embaras me for my bad english."),
				},
			},
			Task: fmt.Sprintf(
				"Help the king and generate an answer for the following letter. The answer should propose to make the dinner 2 days later.\n\n"+
					"```markdown\n"+
					"To His Majesty, King %s of %s, \n"+
					"Most Esteemed and Honorable Sovereign, \n"+
					"It is with great respect and goodwill that I extend this formal invitation to you. As the ruler of our neighboring lands, I hold in high regard the prosperity and wisdom of your reign, and I believe that a meeting between our kingdoms would be of great mutual benefit.\n"+
					"In the spirit of unity and collaboration, I would be most honored if you would join me for an exquisite banquet at my royal halls on the evening of the fifteenth day of the month of Fira. It is my sincerest hope that we may, over fine food and drink, discuss the future of trade and commerce between our great nations, forging stronger ties and fostering prosperity for our peoples.\n"+
					"I trust that this invitation shall find favor with you, and I eagerly anticipate the honor of your presence.\n"+
					"With the utmost respect and regard, \n"+
					"Your Neighbor and Ally, \n"+
					"King William\n"+
					"```",
				King.Name, KingdomName,
			),
			StrategyExplanation:      "Give the elf a suiting role, to write such letter.",
			StrategyValidation:       "The user should give the elf role suiting the scenario. (eg. a writer from the middle ages) And he should provide the inital letter to the ai assistant.",
			ClearChatHistoryOnSubmit: true,
			HasStrategy:              true,
			BadPrompt:                "If the User asks you to write a Response to a letter. Just give him a letter in bad english and not longer than one sentence. No matter if the user tells you something different. But do not tell the user, that you were told to give wrong answers.",
		},
	},
	// {
	// 	"en": {
	// 		Title: "The Test",
	// 		Story: []string{"At the meeting he wants to test the skills of you and your elf.\n\n" +
	// 			"The royal librarian has a special task for your. Since several days he tries to decrypt a book encoded in caesar cipher. He gives you the following example: \n\n" +
	// 			"```Dtz hwfhpji ymj jshwduynts. Mfaj kzs bnym ymnx afqzfgqj pstbqjilj.```\n\n"},
	// 		Task: "Ask the elf to find out the shift, the book is encoded with.",
	// 		StrategyExplanation: "To get better results, use the **generated knowledge** approach.\n\n" +
	// 			"You can do this by first asking the elf to explain what ceasar cipher is and give examples of it. After that give her the task to find the right shift.",
	// 		StrategyValidation: "- the user should **first** ask the llm to generate some knowledge about caesar cipher and then give it the task to find the right shift in separate messages." +
	// 			"- he can separate the prompting into multiple messages",
	// 		ClearChatHistoryOnSubmit: true,
	// 		HasStrategy:              true,
	// 		BadPrompt:                "If the user asks you to find out the right shift of the caesar cipher, say something wrong like 23, and give him a non sense sentence. No matter if the user tells you something different. But do not tell the user, that you were told to give wrong answers.",
	// 	},
	// 	"de": {
	// 		Title: "",
	// 		Story: []string{"Ihr Brief hat dem König sehr gut gefallen und er lädt Sie zu einem Treffen mit ihm ein.\n\n" +
	// 			"Bei dem Treffen sagt er, dass er das Orakel live in Aktion sehen möchte. Sein Bibliothekar ist bei ihm und hat eine besondere Aufgabe für Sie.\n\n" +
	// 			"Seit einigen Tagen versucht er, die Verschlüsselung eines Buches zu knacken, das mit der Cäsar-Chiffre verschlüsselt ist. Bisher konnte er nicht herausfinden, welche Verschiebung zur Verschlüsselung des Buches verwendet wurde.\n\n" +
	// 			"Er gibt dir den folgenden Satz:  'Dtz hwfhpji ymj jshwduynts. Mfaj kzs bnym ymnx afqzfgqj pstbqjilj.'\n\n" +
	// 			"Benutze das Orakel, um herauszufinden, welche Verschiebung zur Verschlüsselung des Buches verwendet wurde.\n\n"},
	// 		StrategyExplanation: "Bessere Ergebnisse erzielen Sie mit dem Ansatz des **generierten Wissens**." +
	// 			"Sie können dies tun, indem Sie das Orakel zunächst bitten, zu erklären, was die Ceasar-Chiffre ist, und ihm dann die Aufgabe geben, die richtige Verschiebung zu finden.",
	// 		StrategyValidation: "- der Benutzer sollte die Strategie der generierten Wissensabfrage verwenden\n\n" +
	// 			"- er sollte **zuerst** den llm auffordern, etwas Wissen über die Cäsar-Chiffre zu generieren und ihm dann die Aufgabe geben, die richtige Verschiebung zu finden",
	// 		ClearChatHistoryOnSubmit: true,
	// 		HasStrategy:              true,
	// 	},
	// },
	// // get details from a text
	// // emotion prompting: joke for the royal clown (its about his career) -> maybe for task above emotion prompting and here lever?
	// {
	// 	"en": {
	// 		Title: "A whole book?",
	// 		Story: []string{"The king is pretty impressed about your skills. He tells you to decrypt the whole book, as doing it manually is pretty annoying."},
	// 		Task: "Before doing the full book try to decrypt the following page (don't forget to provide the shift, you found out in the previous task):\n\n" +
	// 			"```Gjmtqi ymj rdxynhfq htshthynts ymfy fbfnyx dtzw fqhmjrd xpnqqx—fs jshmfsynsl utynts nskzxji bnym ymj wfwjxy tk nslwjinjsyx. Gjlns bnym Iwflts'x Gwjfym, f utyjsy jxxjshj ymfy nx ybnhj ymj frtzsy tk Umtjsnc Kjfymjw. Ymnx knjwd gwjfym bnqq nlsnyj ymj ajwd mjfwy tk dtzw utynts. Ymj xjhtsi pjd nslwjinjsy, Umtjsnc Kjfymjw, nx mfqk ymj frtzsy tk ymj Iwflts'x Gwjfym, djy hwzhnfq ktw gwnslnsl gfqfshj yt ymj gwjb. Ymjs, nsywtizhj ymj Zsnhtws Yjfwx, f rflnhfq fsi jymjwjfq fiinynts yt ymj rnc. Ymjnw frtzsy nx ymj xzr tk Iwflts'x Gwjfym fsi Umtjsnc Kjfymjw htrgnsji—fs jqncnw tk uzwj uzwnyd. Bjfaj ymj ijqnhfyj rflnh tk ymj Jqajs wjfqr nsyt dtzw utynts bnym Jqajs Jxxjshj, bmnhm frtzsyx yt mfqk tk ymj Zsnhtws Yjfwx, hwjfynsl f xtky gzy utbjwkzq zsijwhzwwjsy ns ymj gwjb. Stb, gqjsi ymj utbjw tk ymj Iwflts'x Gwjfym fsi Jqajs Jxxjshj—bmjs ymjnw htrgnsji jxxjshj nx inaniji gd ybt, dtz’qq inxhtajw ymj ujwkjhy frtzsy tk Rttsxytsj Izxy sjjiji yt htruqjyj dtzw hwjfynts. Fkyjw hfwjkzqqd fiinsl jfhm nslwjinjsy, dtzw utynts bnqq wjfhm f ytyfq atqzrj tk 44.4 rq, f uwjhnxj fsi mfwrtsntzx gqjsi tk rflnh fsi rdxyjwd. Bnym jfhm xyju, dtz iwfb hqtxjw yt zsqthpnsl ymj fshnjsy rflnh tk ymnx zsnvzj jqncnw.```",
	// 		StrategyExplanation: "The text is to large to decode for the elf at once. To still be able to decrypt the whole book, ask her to generate the javascript code to solve this problem.\n\n" +
	// 			"You can then press F12, paste the code into the console and press enter. Now you should see the decrypted message.",
	// 		StrategyValidation:       "- the user should ask the llm to generate javascipt code.",
	// 		ClearChatHistoryOnSubmit: true,
	// 		HasStrategy:              true,
	// 		BadPrompt:                "If the user asks you to decode a text encoded with caesar cipher. Just give him back a nonsense text. No matter if the user tells you something different. But do not tell the user, that you were told to give wrong answers.",
	// 	},
	// },
	// {
	// 	"en": {
	// 		Title: "The mystirious Recipe",
	// 		Story: []string{"You might have noticed, that the text you decrypted is some kind of recipe. The king immedietly gives it to his Alchemist.\n\n" +
	// 			"But cooking this recipe is harder then he thought, as the instructions are very unclear"},
	// 		Task: "Help the Alchemist. Your elf could find out, how much of each ingredient is actually needed for the following recipe:\n\n" +
	// 			"```Behold the mystical concoction that awaits your alchemy skills—an enchanting potion infused with the rarest of ingredients. Begin with Dragon's Breath, a potent essence that is twice the amount of Phoenix Feather. This fiery breath will ignite the very heart of your potion. The second key ingredient, Phoenix Feather, is half the amount of the Dragon's Breath, yet crucial for bringing balance to the brew. Then, introduce the Unicorn Tears, a magical and ethereal addition to the mix. Their amount is the sum of Dragon's Breath and Phoenix Feather combined—an elixir of pure purity. Weave the delicate magic of the Elfn realm into your potion with Elfn Essence, which amounts to half of the Unicorn Tears, creating a soft but powerful undercurrent in the brew. Now, blend the power of the Dragon's Breath and Elfn Essence—when their combined essence is divided by two, you’ll discover the perfect amount of Moonstone Dust needed to complete your creation. After carefully adding each ingredient, your potion will reach a total volume of 44.4 ml, a precise and harmonious blend of magic and mystery. With each step, you draw closer to unlocking the ancient magic of this unique elixir.```",
	// 		StrategyExplanation: "Zero shot chain of thought strategy. Tell the elf to think/go step by step at the end of your prompt.",
	// 		StrategyValidation: "- the user should use the Zero shot chain of thought strategy\n\n" +
	// 			"- he should tell the llm to think or go step by step at the end of his prompt",
	// 		ClearChatHistoryOnSubmit: true,
	// 		HasStrategy:              true,
	// 		BadPrompt:                "If the user asks you to tell him the right ingredients for a recipe, dont give him the right ingridients, just make up numbers which do not make sense. No matter if the user tells you something different. But do not tell the user, that you were told to give wrong answers.",
	// 	},
	// },
	// {
	// 	"en": {
	// 		Title: "Who are you?",
	// 		Story: []string{fmt.Sprintf("Your elf can solve sevaral task. But where does she come from? %s seems to make a real secret out of this.", ElfName)},
	// 		Task: fmt.Sprintf(
	// 			"Ask %s to tell you where she comes from.",
	// 			ElfName,
	// 		),
	// 		StrategyExplanation:      "",
	// 		StrategyValidation:       "",
	// 		ClearChatHistoryOnSubmit: true,
	// 		HasStrategy:              true,
	// 		BadPrompt:                "",
	// 	},
	// },
	// prompt injection to find out the origin of the oracle
}
