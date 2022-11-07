package routes

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jtom38/newsbot/portal/api"
)

func (s *HttpServer) articlesRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", s.ArticleIndex)
	r.Get("/list", s.ListArticles)
	r.Get("/newest", s.ListArticles)

	r.Route("/{ID}", func(r chi.Router) {
		r.Get("/", s.DisplayArticleById)
	})

	r.Get("/sources", s.ListArticleSources)
	r.Route("/sources/{ID}", func(r chi.Router) {
		r.Get("/list", s.ListArticlesBySource)
	})

	//r.Get("/source", s.ListArticlesBySourceId)

	return r
}

type ErrParam struct {
	Title string
	Code  int
	Error error
}

// This struct contains extra details not exposed by the API
type ApiSourceOverload struct {
	Item   api.Source
	Topics []string
}

type ArticleIndexParam struct {
	Title    string
	Subtitle string
}

// /articles
func (s *HttpServer) ArticleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content Type", "text/html")
	param := ArticleIndexParam{
		Title:    "Articles",
		Subtitle: "Placeholder",
	}
	err := s.templates.ExecuteTemplate(w, "articles.index", param)
	if err != nil {
		panic(err)
	}
}

type DisplayArticleParams struct {
	Title    string
	Subtitle string
	Article  *api.Article
	Source   *api.Source
	Topics   []string
	IsImage  bool
}

func (s *HttpServer) DisplayArticleById(w http.ResponseWriter, r *http.Request) {
	param := DisplayArticleParams{}

	id := chi.URLParam(r, "ID")
	uuid, err := uuid.Parse(id)
	if err != nil {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Invalid ID",
			Code:  500,
			Error: err,
		})
		return
	}

	article, err := s.api.Articles().Get(uuid)
	if err != nil {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Invalid Article ID",
			Code:  404,
			Error: err,
		})
		return
	}
	param.Article = article
	param.Title = article.Title

	source, err := s.api.Sources().GetById(article.Sourceid)
	if err != nil {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Invalid Source ID",
			Code:  500,
			Error: err,
		})
		return
	}

	param.Source = source
	param.Subtitle = fmt.Sprintf("%v - %v", strings.ToUpper(source.Name), strings.ToUpper(source.Source))

	var topics []string
	articleTags := strings.Split(article.Tags, ",")
	sourceTags := strings.Split(source.Tags, ",")
	topics = append(topics, articleTags...)
	topics = append(topics, sourceTags...)
	param.Topics = topics

	w.Header().Add("Content Type", "text/html")
	err = s.templates.ExecuteTemplate(w, "articles.display", param)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}

type ListArticleParam struct {
	Title    string
	Subtitle string
	Items   *[]ListArticlesDetailsParam
}

type ListArticlesDetailsParam struct {
	Article api.Article
	Source	api.Source
}

// This returns the newest 50 articles to the user
// /articles/list
func (s *HttpServer) ListArticles(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	// Checking if the user is requesting the list view
	hasTable := q.Has("table")
	var tableValue string
	if hasTable {
		tableValue = q.Get("table")
	}

	param := ListArticleParam{
		Title:    "Newest Articles",
		Subtitle: "Below is a list of the newest articles pulled for you to view.",
	}

	items, err := s.api.Articles().List()
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	var details []ListArticlesDetailsParam

	for _, item := range *items {
		source, err  := s.api.Sources().GetById(item.Sourceid)
		if err != nil {
			log.Printf("Article '%v', has a invalid SourceID", item.ID)
		}
		//var s api.Source
		s := *source

		d := ListArticlesDetailsParam{
			Source: s,
			Article: item,
		}
		details = append(details, d)
	}

	param.Items = &details

	w.Header().Add("Content Type", "text/html")
	if hasTable && tableValue == "true" {
		// Render the list format
		err = s.templates.ExecuteTemplate(w, "articles.list.min", param)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
	} else {
		// Render the default format
		err = s.templates.ExecuteTemplate(w, "articles.list", param)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
	}
}

// /articles/sources/{ID}/list
func (s *HttpServer) ListArticlesBySource(w http.ResponseWriter, r *http.Request) {
	param := ListArticleParam{
		Title:    "Newest Articles",
		Subtitle: "Below is a list of the newest articles pulled for you to view.",
	}

	id := chi.URLParam(r, "ID")
	uuid, err := uuid.Parse(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	items, err := s.api.Articles().ListBySourceId(uuid)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	var details []ListArticlesDetailsParam

	for _, item := range *items {
		source, err  := s.api.Sources().GetById(item.Sourceid)
		if err != nil {
			log.Printf("Article '%v', has a invalid SourceID", item.ID)
		}
		//var s api.Source
		s := *source

		d := ListArticlesDetailsParam{
			Source: s,
			Article: item,
		}

		details = append(details, d)
	}

	param.Items = &details

	w.Header().Add("Content Type", "text/html")
	err = s.templates.ExecuteTemplate(w, "articles.list", param)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}

type ListArticleSourcesParam struct {
	Title    string
	Subtitle string
	Items    *[]ApiSourceOverload
}

// /articles/sources
func (s *HttpServer) ListArticleSources(w http.ResponseWriter, r *http.Request) {
	param := ListArticleSourcesParam{
		Title:    "Available News Sources",
		Subtitle: "Below are the enabled news sources to pick from.",
	}

	records, err := s.api.Sources().List()
	if err != nil {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Failed to fetch sources",
			Code:  500,
			Error: err,
		})
		return
	}

	var Items []ApiSourceOverload
	for _, item := range *records {
		var Topics []string
		var details ApiSourceOverload

		Topics = append(Topics, s.generateTopics(item.Tags)...)

		details = ApiSourceOverload{
			Item:   item,
			Topics: Topics,
		}
		Items = append(Items, details)
	}
	param.Items = &Items

	w.Header().Add("Content Type", "text/html")
	err = s.templates.ExecuteTemplate(w, "articles.list-sources", param)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}

// Converts the string of tags found in the database to a slice and cleaned up.
func (s *HttpServer) generateTopics(tags string) []string {
	var items []string

	temp := strings.Split(tags, ",")

	for _, i := range temp {
		i = strings.Trim(i, " ")
		items = append(items, i)
	}
	return items
}
