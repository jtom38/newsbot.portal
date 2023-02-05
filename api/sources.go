package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type SourcesApiClient struct {
	apiServer string
	routeRoot string

	rest RestClient
}

func NewSourcesApiClient(serverAddress string) SourcesApiClient {
	c := SourcesApiClient{
		apiServer: serverAddress,
		routeRoot: "api/sources",
	}
	return c
}

type listSourcesResult struct {
	Message string   `json:"message"`
	Status  int      `json:"status"`
	Payload []Source `json:"payload"`
}

func (c SourcesApiClient) List(ctx context.Context) (*[]Source, error) {
	var items listSourcesResult

	uri := fmt.Sprintf("%v/%v", c.apiServer, c.routeRoot)
	data, err := c.rest.Get(ctx, RestArgs{
		Url:         uri,
		StatusCode:  http.StatusOK,
		ContentType: ContentTypeJson,
	})
	if err != nil {
		return &items.Payload, err
	}

	err = json.Unmarshal(data, &items)
	if err != nil {
		return &items.Payload, err
	}

	return &items.Payload, nil
}

func (c SourcesApiClient) ListBySource(ctx context.Context, value string) (*[]Source, error) {
	var items listSourcesResult

	uri := fmt.Sprintf("%v/%v/by/source?source=%v", c.apiServer, c.routeRoot, value)

	data, err := c.rest.Get(ctx, RestArgs{
		Url:         uri,
		StatusCode:  http.StatusOK,
		ContentType: ContentTypeJson,
	})
	if err != nil {
		return &items.Payload, err
	}

	err = json.Unmarshal(data, &items)
	if err != nil {
		return &items.Payload, err
	}

	return &items.Payload, nil
}

type singleSourcesResult struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Payload Source `json:"payload"`
}

func (c SourcesApiClient) GetById(ctx context.Context, ID uuid.UUID) (*Source, error) {
	var items singleSourcesResult

	uri := fmt.Sprintf("%v/%v/%v", c.apiServer, c.routeRoot, ID)
	body, err := c.rest.Get(ctx, RestArgs{
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

	return &items.Payload, nil
}

func (c SourcesApiClient) GetBySourceAndName(ctx context.Context, SourceName string, Name string) (*Source, error) {
	var items singleSourcesResult

	uri := fmt.Sprintf("%v/%v/by/sourceAndName?source=%v&name=%v", c.apiServer, c.routeRoot, SourceName, Name)

	body, err := c.rest.Get(ctx, RestArgs{
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

	return &items.Payload, nil
}

func (c SourcesApiClient) NewReddit(ctx context.Context, name string, sourceUrl string) error {
	endpoint := fmt.Sprintf("%v/%v/new/reddit?name=%v&url=%v", c.apiServer, c.routeRoot, name, url.QueryEscape(sourceUrl))
	res, err := c.rest.Post(ctx, RestArgs{
		Url:         endpoint,
		StatusCode:  http.StatusOK,
		ContentType: ContentTypeJson,
	})
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New("unexpected status code")
	}

	return nil
}

func (c SourcesApiClient) NewYouTube(ctx context.Context, Name string, Url string) error {
	endpoint := fmt.Sprintf("%v/%v/new/youtube?name=%v&url=%v", c.apiServer, c.routeRoot, Name, url.QueryEscape(Url))

	res, err := c.rest.Post(ctx, RestArgs{
		Url:         endpoint,
		StatusCode:  http.StatusOK,
		ContentType: ContentTypeJson,
	})
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New("unexpected status code")
	}

	return nil
}

func (c SourcesApiClient) NewTwitch(ctx context.Context, Name string) error {
	endpoint := fmt.Sprintf("%v/%v/new/twitch?name=%v", c.apiServer, c.routeRoot, Name)

	res, err := c.rest.Post(ctx, RestArgs{
		Url:         endpoint,
		StatusCode:  http.StatusOK,
		ContentType: ContentTypeJson,
	})
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New("unexpected status code")
	}

	return nil
}

func (c SourcesApiClient) Delete(ctx context.Context, ID uuid.UUID) error {
	endpoint := fmt.Sprintf("%v/%v/%v", c.apiServer, c.routeRoot, ID)

	resp, err := c.rest.Delete(ctx, RestArgs{
		Url:         endpoint,
		ContentType: ContentTypeJson,
		StatusCode:  http.StatusOK,
	})
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("invalid status code")
	}

	return nil
}

func (c SourcesApiClient) Disable(ctx context.Context, ID uuid.UUID) error {
	endpoint := fmt.Sprintf("%v/%v/%v/disable", c.apiServer, c.routeRoot, ID)

	resp, err := c.rest.Post(ctx, RestArgs{
		Url:         endpoint,
		StatusCode:  http.StatusOK,
		ContentType: ContentTypeJson,
	})
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("invalid status code")
	}

	return nil
}

func (c SourcesApiClient) Enable(ctx context.Context, ID uuid.UUID) error {
	endpoint := fmt.Sprintf("%v/%v/%v/enable", c.apiServer, c.routeRoot, ID)

	resp, err := c.rest.Post(ctx, RestArgs{
		Url:         endpoint,
		StatusCode:  http.StatusOK,
		ContentType: ContentTypeJson,
	})
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("invalid status code")
	}

	return nil
}
