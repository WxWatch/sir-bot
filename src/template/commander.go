package template

import (
	"github.com/bwmarrin/discordgo"
	"wxwatch.dev/bot/src/discord"
)

var applicationCommands = []*discordgo.ApplicationCommand{
	{
		Name: "Template",
		Type: discordgo.MessageApplicationCommand,
	},
}

type TemplateCommander struct{}

func NewTemplateCommanderr() *TemplateCommander {
	return &TemplateCommander{}
}

func (c *TemplateCommander) Setup(bot *discord.Bot) {
	applicationHandler := func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		})
	}

	bot.SetupApplicationCommands(applicationCommands, applicationHandler)
}
