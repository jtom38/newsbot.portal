package api

import (
	"time"

	"github.com/google/uuid"
)

// This is the Articles object that comes from the API.
type ArticleDTO struct {
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

type SourceDTO struct {
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
