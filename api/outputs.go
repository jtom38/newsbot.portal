package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
)

type OutputApiClient struct {
	endpoint string
	client *http.Client
}

// Returns all the WebHooks known to the API.
func (c *OutputApiClient) ListDiscordWebHooks() (*[]Discordwebhook, error) {
	var items []Discordwebhook
	uri := fmt.Sprintf("%v/api/discord/webhooks", c.endpoint)

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return &items, err
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := c.client.Do(req)
	if err != nil {
		return &items, err
	}

	if res.StatusCode != 200 {
		return &items, errors.New("unexpected status code")
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return &items, err
	}

	err = json.Unmarshal(body, &items)
	if err != nil {
		return &items, err
	}

	return &items, nil
}

// Returns a single Webhook based on its ID value.
func (c *OutputApiClient) GetDiscordWebHook(id uuid.UUID) (*Discordwebhook, error) {
	var item Discordwebhook
	uri := fmt.Sprintf("%v/api/discord/webhook/%v", c.endpoint, id)

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return &item, err
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := c.client.Do(req)
	if err != nil {
		return &item, err
	}

	if res.StatusCode != 200 {
		return &item, errors.New("unexpected status code")
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return &item, err
	}

	err = json.Unmarshal(body, &item)
	if err != nil {
		return &item, err
	}

	return &item, nil
}

func (c *OutputApiClient) DeleteDiscordWebHook(id uuid.UUID) error {
	uri := fmt.Sprintf("%v/api/discord/webhook/%v", c.endpoint, id)

	req, err := http.NewRequest("DELETE", uri, nil)
	if err != nil {
		return  err
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := c.client.Do(req)
	if err != nil {
		return  err
	}

	if res.StatusCode != 200 {
		return errors.New("unexpected status code")
	}

	return nil
}

func (c *OutputApiClient) DisableDiscordWebHook(id uuid.UUID) error {
	uri := fmt.Sprintf("%v/api/discord/webhook/%v/disable", c.endpoint, id)

	req, err := http.NewRequest("POST", uri, nil)
	if err != nil {
		return  err
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := c.client.Do(req)
	if err != nil {
		return  err
	}

	if res.StatusCode != 200 {
		return errors.New("unexpected status code")
	}

	return nil
}

func (c *OutputApiClient) EnableDiscordWebHook(id uuid.UUID) error {
	uri := fmt.Sprintf("%v/api/discord/webhook/%v/enable", c.endpoint, id)

	req, err := http.NewRequest("POST", uri, nil)
	if err != nil {
		return  err
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := c.client.Do(req)
	if err != nil {
		return  err
	}

	if res.StatusCode != 200 {
		return errors.New("unexpected status code")
	}

	return nil
}