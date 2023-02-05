package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type SubscriptionsApiClient struct {
	endpoint   string
	routeRoute string
	client     *RestClient
}

func NewSubscriptionsClient(endpoint string) SubscriptionsApiClient {
	c := SubscriptionsApiClient{
		endpoint:   endpoint,
		routeRoute: "api/subscriptions",
		client:     NewRestClient(),
	}
	return c
}

type listSubscriptionsResult struct {
	Message string         `json:"message"`
	Status  int            `json:"status"`
	Payload []Subscription `json:"payload"`
}

func (c SubscriptionsApiClient) List(ctx context.Context) ([]Subscription, error) {
	var items listSubscriptionsResult

	uri := fmt.Sprintf("%v/%v", c.endpoint, c.routeRoute)

	body, err := c.client.Get(ctx, RestArgs{
		Url:         uri,
		StatusCode:  http.StatusOK,
		ContentType: ContentTypeJson,
	})
	if err != nil {
		return items.Payload, err
	}

	err = json.Unmarshal(body, &items)
	if err != nil {
		return items.Payload, err
	}

	return items.Payload, err
}

type listSubscriptionsDetailsResult struct {
	RestPayload
	Payload []SubscriptionDetails `json:"payload"`
}

// This will collect subscription details on the webhook and source.
func (c SubscriptionsApiClient) ListDetails(ctx context.Context) ([]SubscriptionDetails, error) {
	var items listSubscriptionsDetailsResult

	uri := fmt.Sprintf("%v/%v/details", c.endpoint, c.routeRoute)

	body, err := c.client.Get(ctx, RestArgs{
		Url:         uri,
		StatusCode:  http.StatusOK,
		ContentType: ContentTypeJson,
	})
	if err != nil {
		return items.Payload, err
	}

	err = json.Unmarshal(body, &items)
	if err != nil {
		return items.Payload, err
	}

	return items.Payload, err
}

func (c SubscriptionsApiClient) GetByDiscordID(ctx context.Context, ID uuid.UUID) (*[]Subscription, error) {
	var items listSubscriptionsResult

	uri := fmt.Sprintf("%v/%v/by/discordId?id=%v", c.endpoint, c.routeRoute, ID.String())
	body, err := c.client.Get(context.Background(), RestArgs{
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

func (c SubscriptionsApiClient) GetBySourceID(ctx context.Context, ID uuid.UUID) (*[]Subscription, error) {
	var items listSubscriptionsResult

	uri := fmt.Sprintf("%v/%v/by/SourceId?id=%v", c.endpoint, c.routeRoute, ID.String())
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

func (c SubscriptionsApiClient) New(ctx context.Context, DiscordID uuid.UUID, SourceID uuid.UUID) error {
	uri := fmt.Sprintf("%v/%v/discord/webhook/new?discordWebHookId=%v&sourceId=%v", c.endpoint, c.routeRoute, DiscordID.String(), SourceID.String())

	_, err := c.client.Post(ctx, RestArgs{
		Url:         uri,
		StatusCode:  http.StatusOK,
		ContentType: ContentTypeJson,
	})

	return err
}

func (c SubscriptionsApiClient) Delete(ctx context.Context, ID uuid.UUID) error {
	uri := fmt.Sprintf("%v/%v/discord/webhook/delete?id=%v", c.endpoint, c.routeRoute, ID.String())

	_, err := c.client.Delete(ctx, RestArgs{
		Url:        uri,
		StatusCode: http.StatusOK,
	})

	return err
}
