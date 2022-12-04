package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
)

type ArticlesApiClient struct {
	endpoint string
	client   *http.Client
}

func NewArticlesClient(Endpoint string, Client *http.Client) ArticlesApi {
	c := ArticlesApiClient{
		endpoint: Endpoint,
		client:   Client,
	}
	return c
}

func (c ArticlesApiClient) List() (*[]Article, error) {
	var items []Article
	var results []ArticleDTO

	uri := fmt.Sprintf("%v/api/articles", c.endpoint)
	res, err := http.Get(uri)
	if err != nil {
		return &items, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return &items, err
	}
	defer res.Body.Close()

	err = json.Unmarshal(body, &results)
	if err != nil {
		return &items, err
	}

	for _, i := range results {
		items = append(items, c.convertDtoObject(i))
	}

	return &items, nil
}

// Returns a single article based on its iD
//
// /api/articles/{id}
func (c ArticlesApiClient) Get(ID uuid.UUID) (*Article, error) {
	var result ArticleDTO
	var items Article

	uri := fmt.Sprintf("%v/api/articles/%v", c.endpoint, ID)
	res, err := http.Get(uri)
	if err != nil {
		return &items, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return &items, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return &items, err
	}

	items = c.convertDtoObject(result)

	return &items, nil
}

// Returns the articles that are bound to a source, by ID value.
//
// /api/articles/by/sourceid?id={id}
func (c ArticlesApiClient) ListBySourceId(ID uuid.UUID) (*[]Article, error) {
	var result []ArticleDTO
	var items []Article

	uri := fmt.Sprintf("%v/api/articles/by/sourceid?id=%v", c.endpoint, ID)
	res, err := http.Get(uri)
	if err != nil {
		return &items, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return &items, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return &items, err
	}

	for _, i := range result {
		items = append(items, c.convertDtoObject(i))
	}

	return &items, nil
}

func (c *ArticlesApiClient) convertDtoObject(i ArticleDTO) Article {
	n := Article{
		ID:          i.ID,
		Sourceid:    i.Sourceid,
		Tags:        splitTags(i.Tags),
		Title:       i.Title,
		Url:         i.Url,
		Pubdate:     i.Pubdate,
		Video:       i.Video.String,
		Videoheight: i.Videoheight,
		Videowidth:  i.Videowidth,
		Thumbnail:   i.Thumbnail,
		Description: i.Description,
		Authorname:  i.Authorname.String,
		Authorimage: i.Authorimage.String,
	}
	return n
}
