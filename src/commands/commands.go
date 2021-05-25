package commands

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"wxwatch.dev/bot/src/cache"
)

const prefix = "!"

type Command struct {
	Name    string
	Execute func(s *discordgo.Session, e *discordgo.MessageCreate, router *CommandRouter) error
}

type CommandRouter struct {
	commands []*Command
	cache    *cache.Cache
}

type Options struct {
	Cache *cache.Cache
}

func NewCommandRouter(options *Options) *CommandRouter {
	return &CommandRouter{
		commands: []*Command{
			Help,
			Ping,
			Pong,
			CruiseFact,
		},
		cache: options.Cache,
	}
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

}
