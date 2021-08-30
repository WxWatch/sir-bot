package leveling

import "github.com/bwmarrin/discordgo"

var applicationCommands = []*discordgo.ApplicationCommand{
	{
		Name:        "rpg",
		Description: "Top-Level RPG Command",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "help",
				Description: "Displays an unhelpful help message",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
			},
			{
				Name:        "leaderboard",
				Description: "Displays the user level leaderboard",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
			},
		},
	},
}
