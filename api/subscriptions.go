package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
)

type SubscriptionsApiClient struct {
	endpoint string
	client   RestClient
}

func NewSubscriptionsClient(endpoint string, client *http.Client) SubscriptionsApiClient {
	c := SubscriptionsApiClient{
		endpoint: endpoint,
		client:   NewRestClient(),
	}
	return c
}

func (c SubscriptionsApiClient) List() (*[]Subscription, error) {
	var items []Subscription

	uri := fmt.Sprintf("%v/api/subscriptions", c.endpoint)
	res, err := http.Get(uri)
	if err != nil {
		return &items, err
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

func (c SubscriptionsApiClient) GetByDiscordID(ID uuid.UUID) (*[]Subscription, error) {
	var items []Subscription

	uri := fmt.Sprintf("%v/api/subscriptions/byDiscordId?id=%v", c.endpoint, ID.String())
	res, err := c.client.Get(context.Background(), RestArgs{
		Url: uri,
		StatusCode: 200,
		Body: nil,
	})
	if err != nil {
		return &items, err
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

func (c SubscriptionsApiClient) GetBySourceID(ID uuid.UUID) (*[]Subscription, error) {
	var items []Subscription

	uri := fmt.Sprintf("%v/api/subscriptions/bySourceId?id=%v", c.endpoint, ID.String())
	res, err := c.client.Get(context.Background(), RestArgs{
		Url: uri,
		StatusCode: 200,
		Body: nil,
	})
	if err != nil {
		return &items, err
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

func (c SubscriptionsApiClient) New(DiscordID uuid.UUID, SourceID uuid.UUID) error {
	uri := fmt.Sprintf("%v/api/subscriptions/new/discord/webhook?discordWebHookId=%v&sourceId=%v", c.endpoint, DiscordID.String(), SourceID.String())

	res, err := c.client.Post(context.Background(), RestArgs{
		Url: uri,
		StatusCode: 200,
		ContentType: ContentTypeJson,
		Body: nil,
	})
	if err != nil {
		return err
	}
	
	defer res.Body.Close()

	return nil
}

func (c SubscriptionsApiClient) Delete(ID uuid.UUID) error {
	uri := fmt.Sprintf("%v/api/subscriptions/discord/webhook/delete?id=%v", c.endpoint, ID.String())

	res, err := c.client.Delete(context.Background(), RestArgs{
		Url: uri,
		StatusCode: 200,
		Body: nil,
	})
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return nil
}
