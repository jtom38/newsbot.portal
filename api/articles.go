package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
