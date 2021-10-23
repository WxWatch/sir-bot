package gacha

import (
	"github.com/bwmarrin/discordgo"
	"wxwatch.dev/bot/src/discord"
)

const numPullsOption = "amount"

var applicationCommands = []*discordgo.ApplicationCommand{
	{
		Name:        "gamers",
		Description: "Any gamers in chat?",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "pull",
				Description: "Displays the user level leaderboard",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionInteger,
						Name:        numPullsOption,
						Description: "Number of characters to pull (max 10)",
						Required:    false,
					},
				},
			},
		},
	},
}

type GachaCommander struct{}

func NewGachaCommander() *GachaCommander {
	return &GachaCommander{}
}

func (c *GachaCommander) Setup(bot *discord.Bot) {
	applicationHandler := func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if len(i.ApplicationCommandData().Options) == 0 {
			return
		}

		if i.ApplicationCommandData().Name != "gamers" {
			return
		}

		logger.Infof("Slash Command Executed: %v %v", i.ApplicationCommandData().Name, i.ApplicationCommandData().Options[0].Name)
		switch i.ApplicationCommandData().Options[0].Name {
		case "pull":
			c.pullHandler(s, i)
		}
	}

	bot.SetupApplicationCommands(applicationCommands, applicationHandler)
}
