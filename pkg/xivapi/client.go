package xivapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/hashicorp/go-cleanhttp"
)

const baseURL = "https://xivapi.com"

type Client struct {
	client *http.Client
}

func NewClient() *Client {
	return &Client{
		client: cleanhttp.DefaultClient(),
	}
}

func (c *Client) CharacterSearch(name string) (*Character, error) {
	url := fmt.Sprintf("%s/character/search?name=%s", baseURL, url.QueryEscape(name))

	r, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var result CharacterSearchResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Results) <= 0 {
		return nil, nil
	}

	return result.Results[0], nil
}

func (c *Client) CharacterDetails(ID int) (*Character, error) {
	url := fmt.Sprintf("%s/character/%v?extended=true", baseURL, ID)

	r, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var result CharacterDetailsResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result.Character, nil
}
