package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type ArticlesApiClient struct {
	endpoint string
	//client   *http.Client
	rest *RestClient
}

func NewArticlesClient(Endpoint string) ArticlesApiClient {
	c := ArticlesApiClient{
		endpoint: Endpoint,
		rest:     NewRestClient(),
	}
	return c
}

type articlesListResult struct {
	RestPayload
	Payload []Article `json:"payload"`
}

type ArticlesListParam struct {
	Page int32
}

// Returns the top 50 Articles based on the page number, if given.
//
// Route = /api/articles
func (c ArticlesApiClient) List(ctx context.Context, param ArticlesListParam) ([]Article, error) {
	var items articlesListResult

	v := url.Values{}
	if param.Page >= 1 {
		v.Add("page", string(param.Page))
	}

	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}

	uri := fmt.Sprintf("%v/api/articles", c.endpoint)
	if len(keys) >= 1 {
		uri = fmt.Sprintf("%v?%v", uri, v.Encode())
	}

	data, err := c.rest.Get(ctx, RestArgs{
		Url:         uri,
		StatusCode:  http.StatusOK,
		ContentType: ContentTypeJson,
	})
	if err != nil {
		return items.Payload, err
	}

	err = json.Unmarshal(data, &items)
	if err != nil {
		return items.Payload, err
	}

	return items.Payload, err
}

type articleGetResult struct {
	RestPayload
	Payload Article `json:"payload"`
}

// Returns a single article based on its iD
//
// Route =  /api/articles/{id}
func (c ArticlesApiClient) Get(ctx context.Context, ID uuid.UUID) (*Article, error) {
	var item articleGetResult

	uri := fmt.Sprintf("%v/api/articles/%v", c.endpoint, ID)
	data, err := c.rest.Get(ctx, RestArgs{
		Url:         uri,
		StatusCode:  http.StatusOK,
		ContentType: ContentTypeJson,
	})
	if err != nil {
		return &item.Payload, err
	}

	err = json.Unmarshal(data, &item)
	if err != nil {
		return &item.Payload, err
	}

	return &item.Payload, err
}

// Returns the articles that are bound to a source, by ID value.
//
// Route = /api/articles/by/sourceid?id={id}
func (c ArticlesApiClient) ListBySourceId(ctx context.Context, ID uuid.UUID, page int) (*[]Article, error) {
	var items articlesListResult

	uri := fmt.Sprintf("%v/api/articles/by/sourceid?id=%v", c.endpoint, ID)
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

	return &items.Payload, err
}

type articleDetailsResult struct {
	RestPayload
	Payload ArticleDetails `json:"payload"`
}

// This returns the details on a Article with the source details attached.
//
// Route = /api/articles/{ID}/details
func (c ArticlesApiClient) GetDetails(ctx context.Context, ID uuid.UUID) (ArticleDetails, error) {
	var items articleDetailsResult

	uri := fmt.Sprintf("%v/api/articles/%v/details", c.endpoint, ID)
	data, err := c.rest.Get(ctx, RestArgs{
		Url:         uri,
		StatusCode:  http.StatusOK,
		ContentType: ContentTypeJson,
	})
	if err != nil {
		return items.Payload, err
	}

	err = json.Unmarshal(data, &items)
	if err != nil {
		return items.Payload, err
	}

	return items.Payload, err
}
