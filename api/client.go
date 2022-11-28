package api

import "net/http"

type ApiClient struct {
	endpoint string

	_articles      ArticlesApi
	_sources       SourcesApi
	_outputs       OutputsApi
	_subscriptions SubscriptionsApi
}

func New(Endpoint string) CollectorApi {
	client := http.Client{}
	c := ApiClient{
		endpoint: Endpoint,

		_articles: NewArticlesClient(Endpoint, &client),
		_sources:  NewSourcesApiClient(Endpoint, &client),
		_outputs:  NewOutputsApiClient(Endpoint, &client),
		_subscriptions: NewSubscriptionsClient(Endpoint, &client),
	}

	return c
}

func (c ApiClient) Articles() ArticlesApi {
	return c._articles
}

func (c ApiClient) Sources() SourcesApi {
	return c._sources
}

func (c ApiClient) Outputs() OutputsApi {
	return c._outputs
}

func (c ApiClient) Subscriptions() SubscriptionsApi {
	return c._subscriptions
}
