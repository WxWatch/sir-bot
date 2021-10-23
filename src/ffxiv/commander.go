package ffxiv

import (
	"github.com/bwmarrin/discordgo"
	"wxwatch.dev/bot/pkg/xivapi"
	"wxwatch.dev/bot/src/discord"
)

const (
	characterNameOption = "character-name"
	serverOption        = "server"
	showClassesOption   = "show-classes"
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
						Name:        characterNameOption,
						Description: "Name of the character to lookup",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        serverOption,
						Description: "Optional server of the character",
					},
					{
						Type:        discordgo.ApplicationCommandOptionBoolean,
						Name:        showClassesOption,
						Description: "If true, show all classes",
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

		if i.ApplicationCommandData().Name != "ffxiv" {
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
			c.characterHandler(s, i)
		}
	}

	bot.SetupApplicationCommands(applicationCommands, applicationHandler)
}
