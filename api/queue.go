package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type QueueClient struct {
	apiServer string
	routeRoot string

	rest RestClient
}

func NewQueueClient(serverAddress string) QueueClient {
	c := QueueClient{
		apiServer: serverAddress,
		routeRoot: "api/queue",
	}

	return c
}

type queueListDiscordWebhooks struct {
	RestPayload
	Payload []ArticleDetails `json:"payload"`
}

func (c QueueClient) ListDiscordWebHooks(ctx context.Context) ([]ArticleDetails, error) {
	var items queueListDiscordWebhooks

	uri := fmt.Sprintf("%v/%v/discord/webhooks", c.apiServer, c.routeRoot)
	body, err := c.rest.Get(ctx, RestArgs{
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

	return items.Payload, nil
}
