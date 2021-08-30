package commands

import (
	"math/rand"

	"github.com/bwmarrin/discordgo"
)

var CruiseFact = &Command{
	Name:    "cruisefact",
	Execute: cruiseFact,
}

var CruiseFactSlash = &discordgo.ApplicationCommand{
	Name:        "cruisefact",
	Description: "Displays a random fact about Cruise and/or onions.",
}

var CruiseFactHandler = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: getCruiseFact(),
		},
	})
}

func cruiseFact(s *discordgo.Session, m *discordgo.MessageCreate, r *CommandRouter) error {
	s.ChannelMessageSend(m.ChannelID, getCruiseFact())

	return nil
}

func getCruiseFact() string {
	facts := []string{
		"Cruise loves onions.",
		"According to the National Onion Association, U.S. onion consumption has increased 50% in the last 20 years.",
		"Wild onions grow on nearly every continent.",
		"Onions have been around since the Bronze Age! The oldest know onion harvest dates back to around 5,000 BC, over 7,000 years ago!",
		"The sulfuric compounds in onions cause is to cry when we chop them. To cut down on the crying, chill the onion and cut into the root end of the onion last.",
		"There are less than 1,000 onion farmers in the United States. That’s a pretty low number. Altogether, about 125,000 acres of onions are planted in the US each year.",
		"The onion was worshiped by ancient Egyptians. They believed that its spherical shape and concentric rings symbolized eternity. They used to cover the tombs of their rulers with onion pictures and onions played a vital role in burial rituals. They believed that onions would help the dead succeed in the afterlife.",
		"In an old English Rhyme, the thickness of an onion skin was thought to help predict the severity of the upcoming winter. Thin skins mean a mild winter is coming while thick skins indicate a rough winter ahead.",
		"Before it was known as the Big Apple, New York City was called the Big Onion because it was a place where you could peel off layer after layer without ever reaching the core.",
		"The Guinness Book of World Records lists the largest onion ever grown as a whopping 10 pound, 14 ounce onion!",
		"During the Middle Ages, onions were actually used as gifts and even currency. People used to pay for services and goods and even paid rent using onions.",
		"General Ulysses S. Grant, during American Civil War, sent a telegram. Addressed to the War Department, it read – “I will not move my army without onions”. The result was that three train cars loaded with onions were immediately shipped.",
		"Onions are the world’s 6th most popular vegetable crop (based on the volume produced per year).",
	}

	idx := rand.Intn(len(facts))

	return facts[idx]
}
