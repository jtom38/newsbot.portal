package api

import "net/http"


type ApiClient struct {
	endpoint string

	Articles ArticlesApiClient
	Sources SourcesApiClient
}

func New(Endpoint string) *ApiClient {
	c := ApiClient{
		endpoint: Endpoint,

		Articles: ArticlesApiClient{
			endpoint: Endpoint,
			client: &http.Client{},
		},

		Sources: SourcesApiClient{
			endpoint:  Endpoint,
			client: &http.Client{},
		},
	}

	return &c
}