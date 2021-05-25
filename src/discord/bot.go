package discord

import (
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

func (b *Bot) Setup(handlers []interface{}) {
	for _, handler := range handlers {
		b.session.AddHandler(handler)
	}

	b.session.Identify.Intents = discordgo.IntentsGuildMessages
}

func (b *Bot) Connect() error {
	err := b.session.Open()

	b.session.UpdateListeningStatus("!help and maybe other stuff :shrug:")

	return err
}

func (b *Bot) Disconnect() {
	b.session.Close()
}
