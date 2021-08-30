package commands

import (
	"math/rand"

	"github.com/bwmarrin/discordgo"
)

var Help = &Command{
	Name:    "help",
	Execute: help,
}

var HelpSlash = &discordgo.ApplicationCommand{
	Name:        "help",
	Description: "Displays another unhelpful message",
}

var HelpHandler = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: getHelp(),
		},
	})
}

func help(s *discordgo.Session, m *discordgo.MessageCreate, r *CommandRouter) error {
	s.ChannelMessageSend(m.ChannelID, getHelp())

	return nil
}

func getHelp() string {
	copypastas := []string{
		"This does not change the fact that in Antarctica there are 21 million penguins and in Malta there are 502,653 inhabitants. So if the penguins decide to invade Malta, each Maltese will have to fight 42 penguins.",
		"If you ask Rick Astley for a DVD of the movie Up, he won’t give it to you because he’s never gonna give you Up. However, by not giving you Up like you asked for it, he’s letting you down. This is known as the Astley paradox.",
		"I hope Zoe wins xD. I’m a Zoe main and she’s just so fun!! People get so trolled by the bubble, and her voice lines are so cute like when she sings about chocolate cake LOL! She’s super random but also smarter than she looks, just like me xD",
		"now ᴘʟᴀʏɪɴɢ: Who asked (Feat: Nobody) ───────────⚪────── ◄◄⠀▐▐⠀►► 𝟸:𝟷𝟾 / 𝟹:𝟻𝟼⠀───○ 🔊",
		"this streamer is look very nerd :/",
		"( ͡° ͜ʖ├┬┴┬",
		"So you're going by 'Octavian' now plebian? Haha what's up spurcifer, it's Tannerius from Rome. Remember me? Me and the other legionaries used to give a hard time. Sorry you were just an easy target. I can see not much has changed. Remember Seira, the girl you had a crush on? Yeah, she's my concubine now. I make over 200 sesterces a year and drive a quadriga chariot. I guess some things never change huh? Nice catching up. Patheticus.",
		"Guys can you please not spam the chat. My mom bought me this new laptop and it gets really hot when the chat is being spammed. Now my leg is starting to hurt because it is getting so hot. Please, if you don't want me to get burned, then dont spam the chat",
	}

	idx := rand.Intn(len(copypastas))

	return copypastas[idx]
}
