package api

import (
	"time"

	"github.com/google/uuid"
)



type NullString struct {
	String string `json:"string"`
	Valid  bool   `json:"valid"`
}

type Article struct {
	ID          uuid.UUID  `json:"id"`
	Sourceid    uuid.UUID  `json:"sourceid"`
	Tags        string     `json:"tags"`
	Title       string     `json:"title"`
	Url         string     `json:"url"`
	Pubdate     time.Time  `json:"pubdate"`
	Video       NullString `json:"video"`
	Videoheight int32      `json:"videoheight"`
	Videowidth  int32      `json:"videowidth"`
	Thumbnail   string     `json:"thumbnail"`
	Description string     `json:"description"`
	Authorname  NullString `json:"authorname"`
	Authorimage NullString `json:"authorimage"`
}

type Discordqueue struct {
	ID        uuid.UUID
	Articleid uuid.UUID
}

type Discordwebhook struct {
	ID      uuid.UUID
	Url     string
	Server  string
	Channel string
	Enabled bool
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
	Options NullString
}

type Source struct {
	ID      uuid.UUID  `json:"id"`
	Site    string     `json:"site"`
	Name    string     `json:"name"`
	Source  string     `json:"source"`
	Type    string     `json:"type"`
	Value   NullString `json:"value"`
	Enabled bool       `json:"enabled"`
	Url     string     `json:"url"`
	Tags    string     `json:"tags"`
}

type Subscription struct {
	ID               uuid.UUID
	Discordwebhookid uuid.UUID
	Sourceid         uuid.UUID
}
