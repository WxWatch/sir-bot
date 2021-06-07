package leveling

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bwmarrin/discordgo"
	"wxwatch.dev/bot/pkg/commander"
	"wxwatch.dev/bot/src/storage"
)

type Leveling struct {
	storage storage.Storage
}

func NewLeveling(opts ...Option) *Leveling {
	options := options{
		storage: storage.NewInMemoryStorage(),
	}

	for _, o := range opts {
		o.apply(&options)
	}

	return &Leveling{
		storage: options.storage,
	}
}

func (l *Leveling) RegisterCommands() *commander.Command {
	rootCommand := commander.NewCommand("rpg",
		commander.WithUsage("!rpg"),
		commander.WithDescription("Misc. stuff related to RPG-like features (xp, levels, inventory, etc :eyes:)"),
		commander.WithRunFunc(l.rootRun),
	)

	leaderboardCommand := commander.NewCommand("leaderboard",
		commander.WithUsage("leaderboard"),
		commander.WithDescription("Show Leaderboard"),
		commander.WithRunFunc(l.leaderboard),
	)

	rootCommand.AddCommand(leaderboardCommand)

	return rootCommand
}

func (l *Leveling) rootRun(cmd *commander.Command, args []string, s *discordgo.Session, e *discordgo.MessageCreate) {
	s.ChannelMessageSendEmbed(e.ChannelID, cmd.DefaultHelp())
}

func (l *Leveling) leaderboard(cmd *commander.Command, args []string, s *discordgo.Session, e *discordgo.MessageCreate) {
	users, _ := l.storage.GetUsers(e.GuildID)
	sort.Sort(storage.ByLevel(users))

	var names strings.Builder
	var levels strings.Builder
	for _, user := range users {
		discordUser, _ := s.User(user.ID)
		names.WriteString(fmt.Sprintln(discordUser.Username))
		levels.WriteString(fmt.Sprintln(user.Level))
	}

	embed := &discordgo.MessageEmbed{
		Title: "Leaderboard",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "User",
				Value:  names.String(),
				Inline: true,
			},
			{
				Name:   "Level",
				Value:  levels.String(),
				Inline: true,
			},
		},
	}

	s.ChannelMessageSendEmbed(e.ChannelID, embed)
}
