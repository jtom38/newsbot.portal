package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type OutputApiClient struct {
	endpoint string

	discordWebHooks OutputDiscordWebHookApi
}

func NewOutputsApiClient(endpoint string) OutputsApi {
	c := OutputApiClient{
		endpoint: endpoint,

		discordWebHooks: NewDiscordWebHooksClient(endpoint),
	}
	return c
}

func (c OutputApiClient) DiscordWebHook() OutputDiscordWebHookApi {
	return c.discordWebHooks
}

type DiscordWebHooksClient struct {
	endpoint string
	client   RestClient
}

func NewDiscordWebHooksClient(endpoint string) DiscordWebHooksClient {
	c := DiscordWebHooksClient{
		endpoint: endpoint,
		client:   *NewRestClient(),
	}
	return c

}

type discordWebHooksListResult struct {
	RestPayload
	Payload []DiscordWebHooks `json:"payload"`
}

// Returns all the WebHooks known to the API.
func (c DiscordWebHooksClient) List(ctx context.Context) (*[]DiscordWebHooks, error) {
	var items discordWebHooksListResult
	uri := fmt.Sprintf("%v/api/discord/webhooks", c.endpoint)

	body, err := c.client.Get(ctx, RestArgs{
		Url:         uri,
		StatusCode:  http.StatusOK,
		ContentType: ContentTypeJson,
	})
	if err != nil {
		return &items.Payload, err
	}

	err = json.Unmarshal(body, &items)
	if err != nil {
		return &items.Payload, err
	}

	return &items.Payload, err
}

type discordWebHookGetResult struct {
	RestPayload
	Payload DiscordWebHooks `json:"payload"`
}

// Returns a single Webhook based on its ID value.
func (c DiscordWebHooksClient) Get(ctx context.Context, id uuid.UUID) (*DiscordWebHooks, error) {
	var item discordWebHookGetResult
	uri := fmt.Sprintf("%v/api/discord/webhooks/%v", c.endpoint, id)

	body, err := c.client.Get(ctx, RestArgs{
		Url:         uri,
		StatusCode:  http.StatusOK,
		ContentType: ContentTypeJson,
	})
	if err != nil {
		return &item.Payload, err
	}

	err = json.Unmarshal(body, &item)
	if err != nil {
		return &item.Payload, err
	}

	return &item.Payload, err
}

func (c DiscordWebHooksClient) Delete(ctx context.Context, id uuid.UUID) error {
	uri := fmt.Sprintf("%v/api/discord/webhooks/%v", c.endpoint, id)

	c.client.Delete(ctx, RestArgs{
		Url:         uri,
		StatusCode:  http.StatusOK,
		ContentType: ContentTypeJson,
	})

	return nil
}

func (c DiscordWebHooksClient) Disable(ctx context.Context, id uuid.UUID) error {
	uri := fmt.Sprintf("%v/api/discord/webhooks/%v/disable", c.endpoint, id)

	_, err := c.client.Post(ctx, RestArgs{
		Url:         uri,
		StatusCode:  http.StatusOK,
		ContentType: ContentTypeJson,
	})
	return err
}

func (c DiscordWebHooksClient) Enable(ctx context.Context, id uuid.UUID) error {
	uri := fmt.Sprintf("%v/api/discord/webhooks/%v/enable", c.endpoint, id)

	_, err := c.client.Post(ctx, RestArgs{
		Url:         uri,
		StatusCode:  http.StatusOK,
		ContentType: ContentTypeJson,
	})
	return err
}

func (c DiscordWebHooksClient) New(ctx context.Context, server string, channel string, url string) error {
	uri := fmt.Sprintf("%v/api/discord/webhooks/new?url=%v&server=%v&channel=%v", c.endpoint, url, server, channel)

	_, err := c.client.Post(ctx, RestArgs{
		Url:         uri,
		StatusCode:  http.StatusOK,
		ContentType: ContentTypeJson,
	})
	return err
}

func (c DiscordWebHooksClient) GetByServerAndChannel(ctx context.Context, server string, channel string) ([]DiscordWebHooks, error) {
	var items discordWebHooksListResult
	uri := fmt.Sprintf("%v/api/discord/webhooks/by/serverAndChannel?server=%v&channel=%v", c.endpoint, server, channel)

	resp, err := c.client.Get(context.Background(), RestArgs{
		Url:         uri,
		StatusCode:  200,
		ContentType: ContentTypeJson,
	})
	if err != nil {
		return items.Payload, err
	}

	err = json.Unmarshal(resp, &items)
	if err != nil {
		return items.Payload, err
	}

	return items.Payload, nil
}
