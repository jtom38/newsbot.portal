package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type SourcesApiClient struct {
	endpoint string
	client   *http.Client
}

func NewSourcesApiClient(endpoint string, client *http.Client) SourcesApi {
	c := SourcesApiClient{
		endpoint: endpoint,
		client:   client,
	}
	return c
}

func (c SourcesApiClient) List() (*[]Source, error) {
	var result []SourceDTO
	var items []Source

	uri := fmt.Sprintf("%v/api/config/sources", c.endpoint)
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

	for _, i := range result {
		items = append(items, c.convertFromDto(i))
	}
 
	return &items, nil
}

func (c SourcesApiClient) ListBySource(value string) (*[]Source, error) {
	var result []SourceDTO
	var items []Source

	uri := fmt.Sprintf("%v/api/config/sources/by/source?source=%v", c.endpoint, value)
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

	for _, i := range result {
		items = append(items, c.convertFromDto(i))
	}

	return &items, nil
}

func (c SourcesApiClient) GetById(ID uuid.UUID) (*Source, error) {
	var result SourceDTO
	var items Source

	uri := fmt.Sprintf("%v/api/config/sources/%v", c.endpoint, ID)
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

	items = c.convertFromDto(result)

	return &items, nil
}

func (c SourcesApiClient) NewReddit(name string, sourceUrl string) error {


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

func (c SourcesApiClient) NewYouTube(Name string, Url string) error {
	endpoint := fmt.Sprintf("%v/api/config/sources/new/youtube?name=%v&url=%v", c.endpoint, Name, url.QueryEscape(Url))
	req, err := http.NewRequest("POST", endpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New("unexpected status code")
	}

	return nil
}

func (c SourcesApiClient) NewTwitch(Name string) error {
	endpoint := fmt.Sprintf("%v/api/config/sources/new/twitch?name=%v", c.endpoint, Name)
	req, err := http.NewRequest("POST", endpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New("unexpected status code")
	}

	return nil
}

func (c SourcesApiClient) Delete(ID uuid.UUID) error {
	endpoint := fmt.Sprintf("%v/api/config/sources/%v", c.endpoint, ID)
	req, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("invalid status code")
	}

	return nil
}

func (c SourcesApiClient) Disable(ID uuid.UUID) error {
	endpoint := fmt.Sprintf("%v/api/config/sources/%v/disable", c.endpoint, ID)
	req, err := http.NewRequest("POST", endpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("invalid status code")
	}

	return nil
}

func (c SourcesApiClient) Enable(ID uuid.UUID) error {
	endpoint := fmt.Sprintf("%v/api/config/sources/%v/enable", c.endpoint, ID)

	req, err := http.NewRequest("POST", endpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("invalid status code")
	}

	return nil
}

func (c SourcesApiClient) convertFromDto(item SourceDTO) Source {
	i := Source{
		ID:      item.ID,
		Site:    item.Site,
		Name:    item.Name,
		Source:  item.Source,
		Type:    item.Type,
		Value:   item.Value.String,
		Enabled: item.Enabled,
		Url:     item.Url,
		Tags:    splitTags(item.Tags),
	}
	return i
}
