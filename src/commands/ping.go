package commands

import "github.com/bwmarrin/discordgo"

var Ping = &Command{
	Name:    "ping",
	Execute: ping,
}

func ping(s *discordgo.Session, m *discordgo.MessageCreate, r *CommandRouter) error {

	s.ChannelMessageSend(m.ChannelID, "Pong!")

	return nil
}
