package clarifai

import (
	"context"
	"fmt"

	"github.com/Clarifai/clarifai-go-grpc/proto/clarifai/api"
	"github.com/Clarifai/clarifai-go-grpc/proto/clarifai/api/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

type Client struct {
	client  api.V2Client
	context context.Context
}

type Prediction struct {
	Name  string
	Value float32
}

func NewClarifaiClient(apiKey string) *Client {
	conn, err := grpc.Dial(
		"api.clarifai.com:443",
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")),
	)

	if err != nil {
		panic(err)
	}

	client := api.NewV2Client(conn)

	ctx := metadata.AppendToOutgoingContext(
		context.Background(),
		"Authorization", fmt.Sprintf("Key %s", apiKey),
	)

	return &Client{
		client:  client,
		context: ctx,
	}
}

func (c *Client) Predict(url string) ([]Prediction, error) {
	predictions := make([]Prediction, 0)

	// This is a publicly available model ID.
	var GeneralModelId = "aaa03c23b3724a16a56b629203edc62c"
	response, err := c.client.PostModelOutputs(
		c.context,
		&api.PostModelOutputsRequest{
			ModelId: GeneralModelId,
			Inputs: []*api.Input{
				{
					Data: &api.Data{
						Image: &api.Image{
							Url: url,
						},
					},
				},
			},
		},
	)
	if err != nil {
		return nil, err
	}
	if response.Status.Code != status.StatusCode_SUCCESS {
		return nil, fmt.Errorf("failed response: %s", response)
	}

	for _, concept := range response.Outputs[0].Data.Concepts {
		predictions = append(predictions, Prediction{
			Name:  concept.Name,
			Value: concept.Value,
		})
	}

	return predictions, nil
}
