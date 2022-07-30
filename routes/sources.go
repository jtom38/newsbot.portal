package routes

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jtom38/newsbot/portal/api"
)

func (s *HttpServer) sourcesRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", s.SourcesIndex)
	r.Get("/reddit", s.SourcesRedditIndex)
	r.Get("/reddit/new", s.SourcesRedditNewDisplay)
	r.Post("/reddit/new/post", s.SourcesRedditNewPost)

	r.Get("/youtube", s.SourcesyYouTubeIndex)
	r.Get("/twitch", s.SourcesTwitchIndex)
	r.Get("/ffxiv", s.SourcesFfxivIndex)

	return r
}

type ListSourcesParam struct {
	Title string
	Subtitle string

	// Defines what source is currently active.
	// Used for routing
	Source string
	Items *[]api.Source
}

// /settings/sources
func (s *HttpServer) SourcesIndex(w http.ResponseWriter, r *http.Request) {

}

// This displays all the reddit sources known to the app.
//
// /settings/sources/reddit
func (s *HttpServer) SourcesRedditIndex(w http.ResponseWriter, r *http.Request) {
	param := ListSourcesParam{
		Title: "Reddit Sources",
		Subtitle: "Here are the available sources.",
		Source: "reddit",
	}

	items, err  := s.api.ListSourcesBySource("reddit")
	if err != nil {
		panic(err)
	}
	param.Items = items

	w.Header().Add("Content Type", "text/html")
	err = s.templates.ExecuteTemplate(w, "sources.list", param)
	if err != nil {
		panic(err)
	}
}

func (s *HttpServer) SourcesRedditNewDisplay(w http.ResponseWriter, r *http.Request) {
	param := ListSourcesParam{
		Title: "Reddit Sources",
		Subtitle: "Here are the available sources.",
		Source: "reddit",
	}

	items, err  := s.api.ListSourcesBySource("reddit")
	if err != nil {
		panic(err)
	}
	param.Items = items

	w.Header().Add("Content Type", "text/html")
	err = s.templates.ExecuteTemplate(w, "sources.new.reddit", param)
	if err != nil {
		panic(err)
	}
}

// This validates the infomation sent from the form and passes it to the API.
func (s *HttpServer) SourcesRedditNewPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Form Error",
			Code: 500,
			Error: err,
		})
		return 
	}

	name := r.Form.Get("name")
	if name == "" {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Invalid Article ID",
			Code: 500,
			Error: err,
		})
		return 
	}

	uri := fmt.Sprintf("https://reddit.com/r/%v", name)
	err = s.api.SourceNewReddit(name, uri)
	if err != nil {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Failed to add new Reddit source",
			Code: 500,
			Error: err,
		})
		return 
	}

	w.Header().Add("Content Type", "text/html")
	err = s.templates.ExecuteTemplate(w, "sources.list", nil)
	if err != nil {
		panic(err)
	}
}


// This displays all the youtube sources known to the app.
//
// /settings/sources/youtube
func (s *HttpServer) SourcesyYouTubeIndex(w http.ResponseWriter, r *http.Request) {
	param := ListSourcesParam{
		Title: "YouTube Sources",
		Subtitle: "Here are the available sources.",
		Source: "YouTube",
	}

	items, err  := s.api.ListSourcesBySource("youtube")
	if err != nil {
		panic(err)
	}
	param.Items = items

	w.Header().Add("Content Type", "text/html")
	err = s.templates.ExecuteTemplate(w, "sources.list", param)
	if err != nil {
		panic(err)
	}
}

// This displays all the twitch sources known to the app.
//
// /settings/sources/twitch
func (s *HttpServer) SourcesTwitchIndex(w http.ResponseWriter, r *http.Request) {
	param := ListSourcesParam{
		Title: "Twitch Sources",
		Subtitle: "Here are the available sources.",
		Source: "twitch",
	}

	items, err  := s.api.ListSourcesBySource("twitch")
	if err != nil {
		panic(err)
	}
	param.Items = items

	w.Header().Add("Content Type", "text/html")
	err = s.templates.ExecuteTemplate(w, "sources.list", param)
	if err != nil {
		panic(err)
	}
}

// This displays all the ffxiv sources known to the app.
//
// /settings/sources/ffxiv
func (s *HttpServer) SourcesFfxivIndex(w http.ResponseWriter, r *http.Request) {
	param := ListSourcesParam{
		Title: "Final Fantasy XIV Sources",
		Subtitle: "Here are the available sources.",
		Source: "ffxiv",
	}

	items, err  := s.api.ListSourcesBySource("ffxiv")
	if err != nil {
		panic(err)
	}
	param.Items = items

	w.Header().Add("Content Type", "text/html")
	err = s.templates.ExecuteTemplate(w, "sources.list", param)
	if err != nil {
		panic(err)
	}
}


func (s *HttpServer) ListSources(w http.ResponseWriter, r *http.Request) {
	param := ListSourcesParam{}

	items, err  := s.api.ListSources()
	if err != nil {
		panic(err)
	}
	param.Items = items

	w.Header().Add("Content Type", "text/html")
	err = s.templates.ExecuteTemplate(w, "sources.list", param)
	if err != nil {
		panic(err)
	}
}

func (s *HttpServer) GetSourceById(w http.ResponseWriter, r *http.Request) {
	param := ListSourcesParam{}

	items, err  := s.api.ListSources()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	param.Items = items

	w.Header().Add("Content Type", "text/html")
	err = s.templates.ExecuteTemplate(w, "sources.list", param)
	if err != nil {
		panic(err)
	}
}


