package commander

import (
	"context"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Command struct {
	Use           string
	Usage         string
	Description   string
	ArgsValidator ArgValidator
	Run           runFunc

	args     []string
	commands []*Command
	parent   *Command
}

func NewCommand(use string, opts ...Option) *Command {
	options := options{
		usage:         "",
		description:   "",
		argsValidator: ArbitraryArgs,
		run:           nil,
		ctx:           context.Background(),
	}

	for _, o := range opts {
		o.apply(&options)
	}

	return &Command{
		Use:           use,
		Usage:         options.usage,
		Description:   options.description,
		ArgsValidator: options.argsValidator,
		Run:           options.run,

		args:     make([]string, 0),
		commands: make([]*Command, 0),
		parent:   nil,
	}
}

func (c *Command) DefaultHelp() *discordgo.MessageEmbed {
	embed := &discordgo.MessageEmbed{
		Title:       c.Use,
		Description: c.Description,
		Fields:      []*discordgo.MessageEmbedField{},
	}

	for _, child := range c.commands {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   child.Usage,
			Value:  child.Description,
			Inline: false,
		})
	}

	return embed
}

// Execute
func (c *Command) Execute(s *discordgo.Session, e *discordgo.MessageCreate) {
	if e.Author.Bot {
		return
	}

	// Parse Args
	err := c.ParseArgs(e.Content)
	if err != nil {
		return
	}

	// Call a child command instead if there is one
	if len(c.args) > 0 {
		possibleCommand := c.args[0]
		for _, command := range c.commands {
			if command.Use == possibleCommand {
				command.Execute(s, e)
				return
			}
		}
	}

	// Validate args
	err = c.ValidateArgs()
	if err != nil {
		return
	}

	// Do it
	go c.Run(c, c.args, s, e)
}

func (c *Command) ValidateArgs() error {
	if c.ArgsValidator == nil {
		return nil
	}
	return c.ArgsValidator(c, c.args)
}

func (c *Command) ParseArgs(raw string) error {
	splitArgs := strings.Split(raw, " ")

	// Get args after the instance of c.Use
	for i, arg := range splitArgs {
		if (arg == c.Use || arg == fmt.Sprintf("!%v", c.Use)) && len(splitArgs) > i {
			c.args = splitArgs[i+1:]
			return nil
		}
	}

	return nil
}

// AddCommand
func (c *Command) AddCommand(command *Command) error {
	if c == command {
		return fmt.Errorf("cannot add command as a subcommand of itself")
	}

	command.SetParent(c)
	c.commands = append(c.commands, command)

	return nil
}

func (c *Command) SetParent(command *Command) error {
	if c == command {
		return fmt.Errorf("command cannot be a parent of itself")
	}
	c.parent = command
	return nil
}
