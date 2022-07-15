package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
)

type ApiClient struct {
	endpoint string
}

func New(Endpoint string) (*ApiClient) {
	c := ApiClient{
		endpoint:  Endpoint,
	}

	return &c
}

func (c *ApiClient) ListArticles() (*[]Article, error) {
	var items []Article

	uri := fmt.Sprintf("%v/api/articles", c.endpoint)
	res, err := http.Get(uri)
	if err != nil {
		return &items, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &items, err
	}

	err = json.Unmarshal(body, &items)
	if err != nil {
		return &items, err
	}

	return &items, nil
}

// Returns a single article based on its iD
// 
// /api/articles/{id}
func (c *ApiClient) GetArticle(ID uuid.UUID) (*Article, error) {
	var items Article

	uri := fmt.Sprintf("%v/api/articles/%v", c.endpoint, ID)
	res, err := http.Get(uri)
	if err != nil {
		return &items, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &items, err
	}

	err = json.Unmarshal(body, &items)
	if err != nil {
		return &items, err
	}

	return &items, nil
}

func (c *ApiClient) ListArticlesBySourceId(ID uuid.UUID) (*[]Article, error) {
	var items []Article

	uri := fmt.Sprintf("%v/api/articles/by/sourceid?id=%v", c.endpoint, ID)
	res, err := http.Get(uri)
	if err != nil {
		return &items, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &items, err
	}

	err = json.Unmarshal(body, &items)
	if err != nil {
		return &items, err
	}

	return &items, nil
}