package api

import (
	"time"

	"github.com/google/uuid"
)

type NullString struct {
	String string `json:"string"`
	Valid  bool   `json:"valid"`
}

// This type contains the Article and source details together.
type ArticleDetails struct {
	Article Article
	Source  Source
}

type Article struct {
	ID          uuid.UUID
	Sourceid    uuid.UUID
	Tags        []string
	Title       string
	Url         string
	Pubdate     time.Time
	Video       string
	Videoheight int32
	Videowidth  int32
	Thumbnail   string
	Description string
	Authorname  string
	Authorimage string
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
	ID      uuid.UUID
	Site    string
	Name    string
	Source  string
	Type    string
	Value   string
	Enabled bool
	Url     string
	Tags    []string
	Deleted bool
}

type Subscription struct {
	ID               uuid.UUID
	Discordwebhookid uuid.UUID
	Sourceid         uuid.UUID
}
