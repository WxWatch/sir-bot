package ffxiv

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"wxwatch.dev/bot/pkg/xivapi"
)

func (c *FFXIVCommander) GetCharacter(name string) (*xivapi.Character, error) {
	return c.client.CharacterSearch(name)
}

func (c *FFXIVCommander) GetCharacterDetails(ID int) (*xivapi.Character, error) {
	return c.client.CharacterDetails(ID)
}

func (c *FFXIVCommander) BasicCharacterEmbed(character *xivapi.Character) *discordgo.MessageEmbed {
	if character == nil {
		return &discordgo.MessageEmbed{
			Title:       "Character Search",
			Description: fmt.Sprintf("Error fetching character: %s", character.Name),
		}
	}

	return &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("%s (%v)", character.Name, character.ID),
		Description: character.Server,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: character.Avatar,
		},
	}
}

func (c *FFXIVCommander) CharacterDetailEmbed(character *xivapi.Character) (*discordgo.MessageEmbed, error) {
	if character == nil {
		return nil, nil
	}

	var fields []*discordgo.MessageEmbedField
	for _, classJob := range character.ClassJobs {
		name := classJob.Class.Abbreviation
		if classJob.Class.Abbreviation != classJob.Job.Abbreviation {
			name = fmt.Sprintf("%s/%s", classJob.Class.Abbreviation, classJob.Job.Abbreviation)
		}

		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   name,
			Value:  fmt.Sprintf("%v", classJob.Level),
			Inline: true,
		})
	}

	return &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("%s (%v)", character.Name, character.ID),
		Description: character.Server,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: character.Avatar,
		},
		Fields: fields,
	}, nil
}
