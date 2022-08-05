package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type SourcesApiClient struct {
	endpoint string
}

func (c *SourcesApiClient) List() (*[]Source, error) {
	var items []Source

	uri := fmt.Sprintf("%v/api/config/sources", c.endpoint)
	res, err := http.Get(uri)
	if err != nil {
		return &items, err
	}
	defer res.Body.Close()

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

func (c *SourcesApiClient) ListBySource(value string) (*[]Source, error) {
	var items []Source

	uri := fmt.Sprintf("%v/api/config/sources/by/source?source=%v", c.endpoint, value)
	res, err := http.Get(uri)
	if err != nil {
		return &items, err
	}
	defer res.Body.Close()

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

func (c *SourcesApiClient) GetById(ID uuid.UUID) (*Source, error) {
	var items Source

	uri := fmt.Sprintf("%v/api/config/sources/%v", c.endpoint, ID)
	res, err := http.Get(uri)
	if err != nil {
		return &items, err
	}
	defer res.Body.Close()

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

func (c *SourcesApiClient) NewReddit(name string, sourceUrl string) error {
	endpoint := fmt.Sprintf("%v/api/config/sources/new/reddit?name=%v&url=%v", c.endpoint, name, url.QueryEscape(sourceUrl))
	res, err := http.Post(endpoint, "application/json", nil)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New("unexpected status code")
	}

	return nil
}

func (c *SourcesApiClient) Delete(ID uuid.UUID) error {
	endpoint := fmt.Sprintf("%v/api/config/sources/%v", c.endpoint, ID)
	req, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client {}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("invalid status code")
	}

	return nil
}

func (c *SourcesApiClient) Disable(ID uuid.UUID) error {
	endpoint := fmt.Sprintf("%v/api/config/sources/%v/disable", c.endpoint, ID)
	req, err := http.NewRequest("POST", endpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client {}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("invalid status code")
	}

	return nil
}

func (c *SourcesApiClient) Enable(ID uuid.UUID) error {
	endpoint := fmt.Sprintf("%v/api/config/sources/%v/enable", c.endpoint, ID)
	req, err := http.NewRequest("POST", endpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client {}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("invalid status code")
	}

	return nil
}
