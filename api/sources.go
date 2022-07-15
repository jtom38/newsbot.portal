package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
)

func (c *ApiClient) ListSources() (*[]Source, error) {
	var items []Source

	uri := fmt.Sprintf("%v/api/config/sources", c.endpoint)
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

func (c *ApiClient) GetSourceById(ID uuid.UUID) (*Source, error) {
	var items Source

	uri := fmt.Sprintf("%v/api/config/sources/%v", c.endpoint, ID)
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