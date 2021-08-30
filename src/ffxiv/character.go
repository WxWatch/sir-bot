package ffxiv

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"wxwatch.dev/bot/pkg/xivapi"
)

func (c *FFXIVCommander) characterHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})

	characterName := ""
	server := ""
	showClasses := false

	for _, option := range i.ApplicationCommandData().Options[0].Options {
		switch option.Name {
		case characterNameOption:
			characterName = option.StringValue()
		case serverOption:
			server = option.StringValue()
		case showClassesOption:
			showClasses = option.BoolValue()
		}
	}

	character, err := c.GetCharacter(characterName, server)
	if err != nil {
		logger.Error(err)
	}
	embed := c.BasicCharacterEmbed(character)

	s.InteractionResponseEdit(s.State.User.ID, i.Interaction, &discordgo.WebhookEdit{
		Embeds: []*discordgo.MessageEmbed{embed},
	})

	character, _ = c.GetCharacterDetails(character.ID)
	detailEmbed, _ := c.CharacterDetailEmbed(character)
	embeds := []*discordgo.MessageEmbed{detailEmbed}

	if showClasses {
		classJobEmbeds, _ := c.CharacterClassJobEmbeds(character.ClassJobs)
		embeds = append(embeds, classJobEmbeds...)
	}

	s.InteractionResponseEdit(s.State.User.ID, i.Interaction, &discordgo.WebhookEdit{
		Embeds: embeds,
	})
}

func (c *FFXIVCommander) GetCharacter(name string, server string) (*xivapi.Character, error) {
	return c.client.CharacterSearch(name, server)
}

func (c *FFXIVCommander) GetCharacterDetails(ID int) (*xivapi.Character, error) {
	return c.client.CharacterDetails(ID)
}

func (c *FFXIVCommander) BasicCharacterEmbed(character *xivapi.Character) *discordgo.MessageEmbed {
	if character == nil {
		return &discordgo.MessageEmbed{
			Title:       "Character Search",
			Description: "Error fetching character",
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

	jobField := &discordgo.MessageEmbedField{
		Name:   "Class",
		Value:  character.ActiveClassJob.Name,
		Inline: true,
	}

	levelField := &discordgo.MessageEmbedField{
		Name:   "Level",
		Value:  fmt.Sprintf("%v", character.ActiveClassJob.Level),
		Inline: true,
	}

	serverField := &discordgo.MessageEmbedField{
		Name:   "Server",
		Value:  fmt.Sprintf("%s (%v)", character.Server, character.DC),
		Inline: true,
	}

	raceField := &discordgo.MessageEmbedField{
		Name:   "Race",
		Value:  character.Race.Name,
		Inline: true,
	}

	townField := &discordgo.MessageEmbedField{
		Name:   "Town",
		Value:  character.Town.Name,
		Inline: true,
	}

	return &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("%s (%v)", character.Name, character.ID),
		Description: character.Title.Name,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: character.Avatar,
		},
		Fields: []*discordgo.MessageEmbedField{jobField, levelField, raceField, serverField, townField},
	}, nil
}

func (c *FFXIVCommander) CharacterClassJobEmbeds(classJobs []*xivapi.ClassJob) ([]*discordgo.MessageEmbed, error) {
	embeds := make([]*discordgo.MessageEmbed, 0)
	classJobFields := make(map[string][]*discordgo.MessageEmbedField)

	for _, classJob := range classJobs {
		fields := classJobFields[classJob.Class.ClassJobCategory.Name]

		name := classJob.Class.Abbreviation
		if classJob.Class.Abbreviation != classJob.Job.Abbreviation {
			name = fmt.Sprintf("%s/%s", classJob.Class.Abbreviation, classJob.Job.Abbreviation)
		}

		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   name,
			Value:  fmt.Sprintf("%v", classJob.Level),
			Inline: true,
		})

		classJobFields[classJob.Class.ClassJobCategory.Name] = fields
	}

	for categoryName, fields := range classJobFields {
		embed := &discordgo.MessageEmbed{
			Title:  categoryName,
			Fields: fields,
		}

		embeds = append(embeds, embed)
	}

	return embeds, nil
}
