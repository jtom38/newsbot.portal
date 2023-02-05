package api

type ApiClient struct {
	endpoint string

	_articles      ArticlesApi
	_sources       SourcesApi
	_outputs       OutputsApi
	_subscriptions SubscriptionsApi
}

func New(Endpoint string) ApiClient {
	//client := http.Client{}
	c := ApiClient{
		endpoint: Endpoint,

		_articles:      NewArticlesClient(Endpoint),
		_sources:       NewSourcesApiClient(Endpoint),
		_outputs:       NewOutputsApiClient(Endpoint),
		_subscriptions: NewSubscriptionsClient(Endpoint),
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
