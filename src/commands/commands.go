package commands

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"wxwatch.dev/bot/pkg/commander"
	"wxwatch.dev/bot/src/cache"
	"wxwatch.dev/bot/src/discord"
	"wxwatch.dev/bot/src/leveling"
	"wxwatch.dev/bot/src/storage"
)

const prefix = "!"

type Command struct {
	Name    string
	Execute func(s *discordgo.Session, e *discordgo.MessageCreate, router *CommandRouter) error
}

type CommandRouter struct {
	commands    []*Command
	newCommands []*commander.Command
	cache       *cache.Cache
}

type Options struct {
	Cache   *cache.Cache
	Storage storage.Storage
}

func NewCommandRouter(options *Options) *CommandRouter {
	leveling := leveling.NewLeveling(
		leveling.WithStorage(options.Storage),
	)
	levelingCommand := leveling.RegisterCommands()
	return &CommandRouter{
		commands: []*Command{
			Help,
			Ping,
			Pong,
			CruiseFact,
		},
		newCommands: []*commander.Command{
			levelingCommand,
		},
		cache: options.Cache,
	}
}

func (r *CommandRouter) SetupSlashCommands(bot *discord.Bot) {
	bot.SetupApplicationCommands(applicationCommands, func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := handlers[i.ApplicationCommandData().Name]; ok {
			logger.Infof("Slash Command Executed: %v", i.ApplicationCommandData().Name)
			h(s, i)
		}
	})

}

func (r *CommandRouter) Route(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	for _, command := range r.commands {
		if strings.HasPrefix(m.Content, fmt.Sprintf("%s%s", prefix, command.Name)) {
			logger.Infof("Command Executed: %v", command.Name)
			command.Execute(s, m, r)
		}
	}

	for _, command := range r.newCommands {
		if strings.HasPrefix(m.Content, fmt.Sprintf("%s%s", prefix, command.Use)) {
			logger.Infof("Command Executed: %v", command.Use)
			command.Execute(s, m)
		}
	}

}
