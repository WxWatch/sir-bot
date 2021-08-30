package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	session *discordgo.Session
}

type Options struct {
	Token string
}

func NewBot(options *Options) *Bot {
	dg, err := discordgo.New("Bot " + options.Token)
	if err != nil {
		logger.Infof("error creating Discord session: ", err)
		return nil
	}

	return &Bot{
		session: dg,
	}
}

func (b *Bot) SetupHandlers(handlers []interface{}) {
	for _, handler := range handlers {
		b.session.AddHandler(handler)
	}

	b.session.Identify.Intents = discordgo.IntentsGuildMessages
}

func (b *Bot) SetupApplicationCommands(commands []*discordgo.ApplicationCommand, handler func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	// Slash command stuff
	for _, command := range commands {
		_, err := b.session.ApplicationCommandCreate(b.session.State.User.ID, "", command)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", command.Name, err)
		}
	}

	b.session.AddHandler(handler)
}

func (b *Bot) Connect() error {
	err := b.session.Open()

	b.session.UpdateListeningStatus("!help and maybe other stuff :shrug:")

	return err
}

func (b *Bot) Disconnect() {
	b.session.Close()
}
