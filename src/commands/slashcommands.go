package commands

import "github.com/bwmarrin/discordgo"

var applicationCommands = []*discordgo.ApplicationCommand{
	CruiseFactSlash,
	HelpSlash,
}

var handlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	CruiseFactSlash.Name: CruiseFactHandler,
	HelpSlash.Name:       HelpHandler,
}
