package api

import "github.com/google/uuid"

type CollectorApi interface {
	Articles() ArticlesApi
	Sources() SourcesApi
	Outputs() OutputsApi
}

type ArticlesApi interface {
	List() (*[]Article, error)
	Get(ID uuid.UUID) (*Article, error)
	ListBySourceId(ID uuid.UUID) (*[]Article, error)
}

type SourcesApi interface {
	List() (*[]Source, error)
	ListBySource(value string) (*[]Source, error)
	GetById(ID uuid.UUID) (*Source, error)
	NewReddit(name string, sourceUrl string) error
	NewYouTube(name string, url string) error
	NewTwitch(Name string) error
	Delete(ID uuid.UUID)  error
	Disable(ID uuid.UUID) error
	Enable(ID uuid.UUID) error
}

type OutputsApi interface {
	DiscordWebHook() OutputDiscordWebHookApi
}

type OutputDiscordWebHookApi interface {
	List() (*[]Discordwebhook, error)
	Get(id uuid.UUID) (*Discordwebhook, error)
	Delete(id uuid.UUID) error
	Disable(id uuid.UUID) error
	Enable(id uuid.UUID) error
	New(server string, channel string, url string) error

}