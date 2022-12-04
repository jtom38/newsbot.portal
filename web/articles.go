package web

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/jtom38/newsbot/portal/api"
)

var (
	pageArticlesIndex       = parseArticles("templates/articles/index.html")
	pageArticlesList        = parseArticles("templates/articles/list.html")
	pageArticlesListCards   = parse("templates/articles/list-card-view.html")
	pageArticlesListSources = parseArticles("templates/articles/list-sources.html")
	pageArticlesDisplay     = parseArticles("templates/articles/display.html")
)

func (s *HttpServer) articlesRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", s.ArticleIndex)
	r.Get("/list", s.ArticleList)
	r.Get("/newest", s.ArticleList)
	r.Get("/list/card", s.ArticleListCards)

	r.Route("/{ID}", func(r chi.Router) {
		r.Get("/", s.DisplayArticleById)
	})

	r.Get("/sources", s.ListArticleSources)
	r.Route("/sources/{ID}", func(r chi.Router) {
		r.Get("/list", s.ListArticlesBySource)
		r.Get("/card", s.CardArticlesBySource)
	})
	return r
}

type TitlesParam struct {
	Title    string
	Subtitle string
	Errors   []string
}

type ErrorParam struct {
	Title    string
	Subtitle string
	Code     int
	Error    string
}

// /articles
func (s *HttpServer) ArticleIndex(w http.ResponseWriter, r *http.Request) {
	var err error
	param := TitlesParam{
		Title:    "Articles",
		Subtitle: "Placeholder",
	}

	if pageArticlesIndex.Execute(w, param); err != nil {
		log.Print(err)
	}
}

type ListArticleParam struct {
	Title    string
	Subtitle string
	Errors   []string
	Items    *[]ListArticlesDetailsParam
}

type ListArticlesDetailsParam struct {
	Article api.Article
	Source  api.Source
}

func (s *HttpServer) ArticleList(w http.ResponseWriter, r *http.Request) {
	param := ListArticleParam{
		Title:    "Newest Posts",
		Subtitle: "Placeholder",
	}

	items, err := s.api.Articles().List()
	if err != nil {
		param.Errors = append(param.Errors, err.Error())
		pageArticlesList.Execute(w, param)
		return
	}

	var details []ListArticlesDetailsParam
	for _, item := range *items {
		source, err := s.api.Sources().GetById(item.Sourceid)
		if err != nil {
			//log.Printf("Article '%v', has a invalid SourceID", item.ID)
		}
		//var s api.Source
		s := *source

		d := ListArticlesDetailsParam{
			Source:  s,
			Article: item,
		}
		details = append(details, d)
	}
	param.Items = &details

	pageArticlesList.Execute(w, param)
}

func (s *HttpServer) ArticleListCards(w http.ResponseWriter, r *http.Request) {
	param := ListArticleParam{
		Title:    "Articles",
		Subtitle: "Placeholder",
	}

	items, err := s.api.Articles().List()
	if err != nil {
		errorPage.Execute(w, ErrorParam{
			Title: "This didn't load correctly...",
			Error: err.Error(),
		})
		return
	}

	var details []ListArticlesDetailsParam
	for _, item := range *items {
		source, err := s.api.Sources().GetById(item.Sourceid)
		if err != nil {
			//log.Printf("Article '%v', has a invalid SourceID", item.ID)
		}
		//var s api.Source
		s := *source

		d := ListArticlesDetailsParam{
			Source:  s,
			Article: item,
		}
		details = append(details, d)
	}
	param.Items = &details

	pageArticlesListCards.Execute(w, param)
}

// This struct contains extra details not exposed by the API
//type ApiSourceOverload struct {
//	Item   api.Source
//	Topics []string
//}

type ListArticleSourcesParam struct {
	Title    string
	Subtitle string
	Errors   []string
	Items    *[]api.Source
}

// /articles/sources
func (s *HttpServer) ListArticleSources(w http.ResponseWriter, r *http.Request) {
	param := ListArticleSourcesParam{
		Title:    "Available News Sources",
		Subtitle: "Below are the enabled news sources to pick from.",
	}

	var activeItems []api.Source

	records, err := s.api.Sources().List()
	if err != nil {
		param.Errors = append(param.Errors, err.Error())
		pageArticlesListSources.Execute(w, param)
		return
	}

	for _, item := range *records {
		if !item.Enabled {
			continue
		}
		activeItems = append(activeItems, item)
	}

	param.Items = &activeItems

	pageArticlesListSources.Execute(w, param)
}

func (s *HttpServer) getArticlesBySourceId(ID uuid.UUID) ([]ListArticlesDetailsParam, error) {
	var details []ListArticlesDetailsParam

	items, err := s.api.Articles().ListBySourceId(ID)
	if err != nil {
		return details, err
	}

	for _, item := range *items {
		source, err := s.api.Sources().GetById(item.Sourceid)
		if err != nil {
			log.Printf("Article '%v', has a invalid SourceID", item.ID)
		}
		//var s api.Source
		s := *source

		d := ListArticlesDetailsParam{
			Source:  s,
			Article: item,
		}

		details = append(details, d)
	}

	return details, nil
}

func (s *HttpServer) ListArticlesBySource(w http.ResponseWriter, r *http.Request) {
	param := ListArticleParam{
		Title:    "Newest Articles",
		Subtitle: "Below is a list of the newest articles pulled for you to view.",
	}

	id := chi.URLParam(r, "ID")
	uid, err := uuid.Parse(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	details, err := s.getArticlesBySourceId(uid)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	if len(details) >= 1 {
		param.Title = fmt.Sprintf("Newest posts from %v", details[0].Source.Name)
	}

	param.Items = &details
	pageArticlesList.Execute(w, param)
}

func (s *HttpServer) CardArticlesBySource(w http.ResponseWriter, r *http.Request) {
	param := ListArticleParam{
		Title:    "Newest Articles",
		Subtitle: "Below is a list of the newest articles pulled for you to view.",
	}

	id := chi.URLParam(r, "ID")
	uid, err := uuid.Parse(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	details, err := s.getArticlesBySourceId(uid)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	param.Title = fmt.Sprintf("Newest posts from %v", details[0].Source.Name)

	param.Items = &details
	pageArticlesListCards.Execute(w, param)
}

type DisplayArticleParams struct {
	Title    string
	Subtitle string
	Errors   []string
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
		param.Errors = append(param.Errors, err.Error())
		pageArticlesDisplay.Execute(w, param)
		return
	}

	article, err := s.api.Articles().Get(uuid)
	if err != nil {
		param.Errors = append(param.Errors, err.Error())
		pageArticlesDisplay.Execute(w, param)
		return
	}
	param.Article = article
	param.Title = article.Title

	source, err := s.api.Sources().GetById(article.Sourceid)
	if err != nil {
		param.Errors = append(param.Errors, err.Error())
		pageArticlesDisplay.Execute(w, param)
		return
	}

	param.Source = source
	param.Subtitle = fmt.Sprintf("%v - %v", strings.ToUpper(source.Name), strings.ToUpper(source.Source))
	pageArticlesDisplay.Execute(w, param)
}
