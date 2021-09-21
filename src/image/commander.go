package image

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"wxwatch.dev/bot/pkg/clarifai"
	"wxwatch.dev/bot/src/discord"
)

var applicationCommands = []*discordgo.ApplicationCommand{
	{
		Name: "Image Prediction",
		Type: discordgo.MessageApplicationCommand,
	},
}

type ImageCommander struct {
	client *clarifai.Client
}

func NewImageCommander() *ImageCommander {
	apiKey := os.Getenv("CLARIFAI_KEY")

	if apiKey == "" {
		log.Fatal("CLARIFAI_KEY is required")
	}

	return &ImageCommander{
		client: clarifai.NewClarifaiClient(apiKey),
	}
}

func (c *ImageCommander) Setup(bot *discord.Bot) {
	applicationHandler := func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.ApplicationCommandData().Name != "Image Prediction" {
			return
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		})

		if len(i.ApplicationCommandData().Resolved.Messages) == 0 {
			s.InteractionResponseEdit(s.State.User.ID, i.Interaction, &discordgo.WebhookEdit{
				Content: "I couldn't find an image in that message :(",
			})
			return
		}

		for _, message := range i.ApplicationCommandData().Resolved.Messages {
			if len(message.Attachments) == 0 {
				s.InteractionResponseEdit(s.State.User.ID, i.Interaction, &discordgo.WebhookEdit{
					Content: "I couldn't find an image in that message :(",
				})
				return
			}

			for _, attachment := range message.Attachments {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
				})

				predictions, err := c.client.Predict(attachment.ProxyURL)
				if err != nil {
					respondWithError(s, i.Interaction, err)
					return
				}

				description := ""
				for _, prediction := range predictions[:5] {
					description += fmt.Sprintf("**%s:** %.1f%%\n", prediction.Name, prediction.Value*100)
				}

				embed := &discordgo.MessageEmbed{
					Title:       "Here's what I think this image is",
					Description: description,
					Thumbnail: &discordgo.MessageEmbedThumbnail{
						URL:      attachment.URL,
						ProxyURL: attachment.ProxyURL,
					},
					Footer: &discordgo.MessageEmbedFooter{
						Text: "Powered by https://www.clarifai.com/",
					},
				}

				_, err = s.InteractionResponseEdit(s.State.User.ID, i.Interaction, &discordgo.WebhookEdit{
					Embeds: []*discordgo.MessageEmbed{embed},
				})
				if err != nil {
					respondWithError(s, i.Interaction, err)
					return
				}
			}

		}

	}

	bot.SetupApplicationCommands(applicationCommands, applicationHandler)
}

func respondWithError(s *discordgo.Session, interaction *discordgo.Interaction, err error) {
	s.InteractionResponseEdit(s.State.User.ID, interaction, &discordgo.WebhookEdit{
		Content: fmt.Sprintf("unformatted error: %s", err.Error()),
	})
}
