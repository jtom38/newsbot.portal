package api

import (
	"time"

	"github.com/google/uuid"
)

type RestPayload struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Article struct {
	ID          uuid.UUID `json:"id"`
	SourceID    uuid.UUID `json:"sourceid"`
	Tags        []string  `json:"tags"`
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	Pubdate     time.Time `json:"pubdate"`
	Video       string    `json:"video"`
	VideoHeight int32     `json:"videoHeight"`
	VideoWidth  int32     `json:"videoWidth"`
	Thumbnail   string    `json:"thumbnail"`
	Description string    `json:"description"`
	AuthorName  string    `json:"authorName"`
	AuthorImage string    `json:"authorImage"`
}

type ArticleDetails struct {
	ID          uuid.UUID `json:"id"`
	Source      Source    `json:"source"`
	Tags        []string  `json:"tags"`
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	Pubdate     time.Time `json:"pubdate"`
	Video       string    `json:"video"`
	VideoHeight int32     `json:"videoHeight"`
	VideoWidth  int32     `json:"videoWidth"`
	Thumbnail   string    `json:"thumbnail"`
	Description string    `json:"description"`
	AuthorName  string    `json:"authorName"`
	AuthorImage string    `json:"authorImage"`
}

type DiscordWebHooks struct {
	ID      uuid.UUID `json:"ID"`
	Url     string    `json:"url"`
	Server  string    `json:"server"`
	Channel string    `json:"channel"`
	Enabled bool      `json:"enabled"`
}

type Discordqueue struct {
	ID        uuid.UUID
	Articleid uuid.UUID
}

type Icon struct {
	ID       uuid.UUID
	Filename string
	Site     string
}

type Setting struct {
	ID      uuid.UUID
	Key     string
	Value   string
	Options string
}

type Source struct {
	ID      uuid.UUID `json:"id"`
	Site    string    `json:"site"`
	Name    string    `json:"name"`
	Source  string    `json:"source"`
	Type    string    `json:"type"`
	Value   string    `json:"value"`
	Enabled bool      `json:"enabled"`
	Url     string    `json:"url"`
	Tags    []string  `json:"tags"`
	Deleted bool      `json:"deleted"`
}

type Subscription struct {
	ID               uuid.UUID
	DiscordWebhookId uuid.UUID
	SourceId         uuid.UUID
}

type SubscriptionDetails struct {
	ID             uuid.UUID
	Source         Source
	DiscordWebHook DiscordWebHooks
}
