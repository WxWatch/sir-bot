package ffxiv

import (
	"github.com/bwmarrin/discordgo"
	"wxwatch.dev/bot/pkg/xivapi"
	"wxwatch.dev/bot/src/discord"
)

var applicationCommands = []*discordgo.ApplicationCommand{
	{
		Name:        "ffxiv",
		Description: "Final Fantasy XIV",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "about",
				Description: "About Stuff",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
			},
			{
				Name:        "character",
				Description: "Search for a character",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "character-name",
						Description: "Name of the character to lookup",
						Required:    true,
					},
				},
			},
		},
	},
}

type FFXIVCommander struct {
	client *xivapi.Client
}

func NewFFXIVCommander() *FFXIVCommander {
	return &FFXIVCommander{
		client: xivapi.NewClient(),
	}
}

func (c *FFXIVCommander) Setup(bot *discord.Bot) {
	applicationHandler := func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if len(i.ApplicationCommandData().Options) == 0 {
			return
		}

		logger.Infof("Slash Command Executed: %v %v", i.ApplicationCommandData().Name, i.ApplicationCommandData().Options[0].Name)
		switch i.ApplicationCommandData().Options[0].Name {
		case "about":
			content :=
				"There's a leaderboard command, at least"
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: content,
				},
			})
		case "character":
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
			})

			characterName := i.ApplicationCommandData().Options[0].Options[0].StringValue()
			character, err := c.GetCharacter(characterName)
			if err != nil {
				logger.Error(err)
			}
			embed := c.BasicCharacterEmbed(character)

			s.InteractionResponseEdit(s.State.User.ID, i.Interaction, &discordgo.WebhookEdit{
				Embeds: []*discordgo.MessageEmbed{embed},
			})

			character, _ = c.GetCharacterDetails(character.ID)
			embed, _ = c.CharacterDetailEmbed(character)

			s.InteractionResponseEdit(s.State.User.ID, i.Interaction, &discordgo.WebhookEdit{
				Embeds: []*discordgo.MessageEmbed{embed},
			})
		}
	}

	bot.SetupApplicationCommands(applicationCommands, applicationHandler)
}
