package api

import (
	"context"

	"github.com/google/uuid"
)

type CollectorApi interface {
	Articles() ArticlesApi
	Sources() SourcesApi
	Outputs() OutputsApi
	Subscriptions() SubscriptionsApi
}

type ArticlesApi interface {
	List(ctx context.Context, param ArticlesListParam) ([]Article, error)
	Get(ctx context.Context, ID uuid.UUID) (*Article, error)
	ListBySourceId(ctx context.Context, ID uuid.UUID, page int) (*[]Article, error)
}

type SourcesApi interface {
	List(ctx context.Context) (*[]Source, error)
	ListBySource(ctx context.Context, value string) (*[]Source, error)

	GetById(ctx context.Context, ID uuid.UUID) (*Source, error)
	GetBySourceAndName(ctx context.Context, SourceName string, Name string) (*Source, error)

	NewReddit(ctx context.Context, name string, sourceUrl string) error
	NewYouTube(ctx context.Context, name string, url string) error
	NewTwitch(ctx context.Context, Name string) error
	Disable(ctx context.Context, ID uuid.UUID) error
	Delete(ctx context.Context, ID uuid.UUID) error
	Enable(ctx context.Context, ID uuid.UUID) error
}

type OutputsApi interface {
	DiscordWebHook() OutputDiscordWebHookApi
}

type OutputDiscordWebHookApi interface {
	List(ctx context.Context) (*[]DiscordWebHooks, error)
	Get(ctx context.Context, id uuid.UUID) (*DiscordWebHooks, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Disable(ctx context.Context, id uuid.UUID) error
	Enable(ctx context.Context, id uuid.UUID) error
	New(ctx context.Context, server string, channel string, url string) error
	GetByServerAndChannel(ctx context.Context, server string, channel string) ([]DiscordWebHooks, error)
}

type SubscriptionsApi interface {
	List(ctx context.Context) ([]Subscription, error)
	GetByDiscordID(ctx context.Context, ID uuid.UUID) (*[]Subscription, error)
	GetBySourceID(ctx context.Context, ID uuid.UUID) (*[]Subscription, error)
	New(ctx context.Context, DiscordID uuid.UUID, SourceID uuid.UUID) error
	Delete(ctx context.Context, ID uuid.UUID) error
}

type QueueApi interface {
	ListDiscordWebHooks()
}
