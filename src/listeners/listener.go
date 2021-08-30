package listeners

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

type MessageCreateListener struct {
}

func NewMessageCreateListener() *MessageCreateListener {
	return &MessageCreateListener{}
}

func (l *MessageCreateListener) Listen(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	if strings.Contains(m.Content, "onion") {
		s.ChannelMessageSendReply(m.ChannelID, "<@199749017150816256>", m.Reference())
	}
}
