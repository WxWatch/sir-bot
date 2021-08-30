package commands

import "github.com/bwmarrin/discordgo"

var Pong = &Command{
	Name:    "pong",
	Execute: pong,
}

func pong(s *discordgo.Session, m *discordgo.MessageCreate, r *CommandRouter) error {

	s.ChannelMessageSend(m.ChannelID, "Ping!")

	return nil
}
