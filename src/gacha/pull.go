package gacha

import (
	"fmt"
	"math"
	"math/rand"
	"sort"

	"github.com/bwmarrin/discordgo"
)

func (c *GachaCommander) pullHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})

	numPulls := 1
	for _, option := range i.ApplicationCommandData().Options[0].Options {
		switch option.Name {
		case numPullsOption:
			numPulls = int(option.IntValue())
		}
	}

	numPulls = int(math.Min(float64(numPulls), float64(10)))

	membersInGuild, err := s.GuildMembers(i.GuildID, "", 1000)
	if err != nil {
		logger.Error(err)
	}

	picks := make([]Pick, 0)
	pickString := ""
	for i := 0; i < numPulls; i++ {
		// Pick member
		randomIndex := rand.Intn(len(membersInGuild))
		pick := membersInGuild[randomIndex]

		// Calculate rarity
		rarity := 0
		if pick.User.Bot {
			// Bots can only be 3* :(
			rarity = 3
		} else {
			switch rndm := rand.Intn(99); {
			case rndm == 0:
				rarity = 5
			case rndm <= 10:
				rarity = 4
			default:
				rarity = 3
			}
		}

		picks = append(picks, Pick{
			name:   pick.User.Username,
			rarity: rarity,
		})

	}

	sort.Slice(picks, func(i, j int) bool {
		return picks[i].rarity > picks[j].rarity
	})

	for _, pick := range picks {
		pickString += fmt.Sprintf("**%d :star:**  %s\n", pick.rarity, pick.name)
	}

	s.InteractionResponseEdit(s.State.User.ID, i.Interaction, &discordgo.WebhookEdit{
		Embeds: []*discordgo.MessageEmbed{
			{
				Title:       "Results",
				Description: pickString,
			},
		},
	})
}

type Pick struct {
	name   string
	rarity int
}
