package routes

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jtom38/newsbot/portal/api"
)

func (s *HttpServer) sourcesRouter() http.Handler {
	r := chi.NewRouter()

	//r.Get("/", s.SourcesIndex)
	r.Post("/delete", s.DeleteSourceById)
	r.Post("/disable", s.DisableSourceById)
	r.Post("/enable", s.EnableSourceById)

	r.Get("/reddit", s.SourcesRedditIndex)
	r.Get("/reddit/new", s.SourcesRedditNewDisplay)
	r.Post("/reddit/new", s.SourcesRedditNewPost)

	r.Get("/youtube", s.SourcesYouTubeIndex)
	r.Get("/youtube/new", s.SourcesYouTubeNewForm)
	r.Post("/youtube/new", s.SourcesYouTubeNewPost)

	r.Get("/twitch", s.SourcesTwitchIndex)
	r.Get("/twitch/new", s.SourcesTwitchNewForm)
	r.Post("/twitch/new", s.SourcesTwitchNewFormPost)

	r.Get("/ffxiv", s.SourcesFfxivIndex)

	return r
}

type ListSourcesParam struct {
	Title    string
	Subtitle string

	// Defines what source is currently active.
	// Used for routing
	Source string
	Items  *[]ListSourcesDetailsParam
}

// This struct gives more details on the source so the template can act more.
type ListSourcesDetailsParam struct {
	Enabled  bool
	Disabled bool
	Item     api.Source
}

// /settings/sources/delete?id
func (s *HttpServer) DeleteSourceById(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Form Error",
			Code:  500,
			Error: err,
		})
		return
	}

	id := r.Form.Get("id")
	if id == "" {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Missing Source ID",
			Code:  500,
			Error: err,
		})
		return
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Invalid Source ID",
			Code:  500,
			Error: err,
		})
		return
	}

	err = s.api.Sources.Delete(uid)
	if err != nil {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Failed to delete the Source",
			Code:  500,
			Error: err,
		})
		return
	}
}

// /settings/sources/enable?id
func (s *HttpServer) EnableSourceById(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Form Error",
			Code:  500,
			Error: err,
		})
		return
	}

	id := r.Form.Get("id")
	if id == "" {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Missing Source ID",
			Code:  500,
			Error: err,
		})
		return
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Invalid Source ID",
			Code:  500,
			Error: err,
		})
		return
	}

	err = s.api.Sources.Enable(uid)
	if err != nil {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Failed to delete the Source",
			Code:  500,
			Error: err,
		})
		return
	}

	w.Header().Add("Content Type", "text/html")
	err = s.templates.ExecuteTemplate(w, "sources.posted", HttpParam{
		Title:    "Source Enabled",
		Subtitle: "It will be reviewed on the next collection",
	})
	if err != nil {
		panic(err)
	}
}

// /settings/sources/disable?id
func (s *HttpServer) DisableSourceById(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Form Error",
			Code:  500,
			Error: err,
		})
		return
	}

	id := r.Form.Get("id")
	if id == "" {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Missing Source ID",
			Code:  500,
			Error: err,
		})
		return
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Invalid Source ID",
			Code:  500,
			Error: err,
		})
		return
	}

	err = s.api.Sources.Disable(uid)
	if err != nil {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Failed to delete the Source",
			Code:  500,
			Error: err,
		})
		return
	}

	w.Header().Add("Content Type", "text/html")
	err = s.templates.ExecuteTemplate(w, "sources.posted", HttpParam{
		Title:    "Source Disabled",
		Subtitle: "It will be skipped on the next collection",
	})
	if err != nil {
		panic(err)
	}
}


/* Reddit */
// This displays all the reddit sources known to the app.
//
// /settings/sources/reddit
func (s *HttpServer) SourcesRedditIndex(w http.ResponseWriter, r *http.Request) {
	var details []ListSourcesDetailsParam
	param := ListSourcesParam{
		Title:    "Reddit Sources",
		Subtitle: "Here are the available sources.",
		Source:   "reddit",
	}

	items, err := s.api.Sources.ListBySource("reddit")
	if err != nil {
		panic(err)
	}

	for _, item := range *items {
		var i ListSourcesDetailsParam

		if !item.Enabled {
			i.Disabled = true
			i.Enabled = false
		} else {
			i.Disabled = false
			i.Enabled = true
		}

		i.Item = item

		details = append(details, i)
	}
	param.Items = &details

	w.Header().Add("Content Type", "text/html")
	err = s.templates.ExecuteTemplate(w, "sources.list", param)
	if err != nil {
		panic(err)
	}
}

func (s *HttpServer) SourcesRedditNewDisplay(w http.ResponseWriter, r *http.Request) {
	var details []ListSourcesDetailsParam
	param := ListSourcesParam{
		Title:    "Reddit Sources",
		Subtitle: "Here are the available sources.",
		Source:   "reddit",
	}

	items, err := s.api.Sources.ListBySource("reddit")
	if err != nil {
		panic(err)
	}

	for _, item := range *items {
		var i ListSourcesDetailsParam

		if !item.Enabled {
			i.Disabled = true
			i.Enabled = false
		} else {
			i.Disabled = false
			i.Enabled = true
		}

		i.Item = item

		details = append(details, i)
	}
	param.Items = &details

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
			Code:  500,
			Error: err,
		})
		return
	}

	name := r.Form.Get("name")
	if name == "" {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Invalid Article ID",
			Code:  500,
			Error: err,
		})
		return
	}

	uri := fmt.Sprintf("https://reddit.com/r/%v", name)
	err = s.api.Sources.NewReddit(name, uri)
	if err != nil {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Failed to add new Reddit source",
			Code:  500,
			Error: err,
		})
		return
	}

	w.Header().Add("Content Type", "text/html")
	err = s.templates.ExecuteTemplate(w, "sources.posted", HttpParam{
		Title:    "Source Added",
		Subtitle: "It will be reviewed on the next collection",
	})
	if err != nil {
		panic(err)
	}
}


/* YouTube */
// This displays all the youtube sources known to the app.
//
// /settings/sources/youtube
func (s *HttpServer) SourcesYouTubeIndex(w http.ResponseWriter, r *http.Request) {
	var details []ListSourcesDetailsParam
	param := ListSourcesParam{
		Title:    "YouTube Sources",
		Subtitle: "Here are the available sources.",
		Source:   "youtube",
	}

	items, err := s.api.Sources.ListBySource("youtube")
	if err != nil {
		panic(err)
	}

	for _, item := range *items {
		var i ListSourcesDetailsParam

		if !item.Enabled {
			i.Disabled = true
			i.Enabled = false
		} else {
			i.Disabled = false
			i.Enabled = true
		}

		i.Item = item

		details = append(details, i)
	}
	param.Items = &details

	w.Header().Add("Content Type", "text/html")
	err = s.templates.ExecuteTemplate(w, "sources.list", param)
	if err != nil {
		panic(err)
	}
}

// This is the form that lets you enter a new youtube source into the application
// /settings/sources/youtube/new
func (s *HttpServer) SourcesYouTubeNewForm(w http.ResponseWriter, r *http.Request) {
	var details []ListSourcesDetailsParam
	param := ListSourcesParam{
		Title:    "YouTube Sources",
		Subtitle: "Here are the available sources.",
		Source:   "youtube",
	}

	items, err := s.api.Sources.ListBySource("youtube")
	if err != nil {
		panic(err)
	}

	for _, item := range *items {
		var i ListSourcesDetailsParam

		if !item.Enabled {
			i.Disabled = true
			i.Enabled = false
		} else {
			i.Disabled = false
			i.Enabled = true
		}

		i.Item = item

		details = append(details, i)
	}
	param.Items = &details

	w.Header().Add("Content Type", "text/html")
	err = s.templates.ExecuteTemplate(w, "sources.new.youtube", param)
	if err != nil {
		panic(err)
	}
}

// This validates the infomation sent from the form and passes it to the API.
func (s *HttpServer) SourcesYouTubeNewPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Form Error",
			Code:  500,
			Error: err,
		})
		return
	}

	name := r.Form.Get("name")
	if name == "" {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Missing Name value",
			Code:  500,
			Error: err,
		})
		return
	}

	url := r.Form.Get("url")
	if name == "" {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Missing URL value",
			Code:  500,
			Error: err,
		})
		return
	}

	err = s.api.Sources.NewYouTube(name, url)
	if err != nil {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Failed to add new YouTube source",
			Code:  500,
			Error: err,
		})
		return
	}

	w.Header().Add("Content Type", "text/html")
	err = s.templates.ExecuteTemplate(w, "sources.posted", HttpParam{
		Title:    "Source Added",
		Subtitle: "It will be reviewed on the next collection",
	})
	if err != nil {
		panic(err)
	}
}


/* Twitch */
// This displays all the twitch sources known to the app.
//
// /settings/sources/twitch
func (s *HttpServer) SourcesTwitchIndex(w http.ResponseWriter, r *http.Request) {
	var details []ListSourcesDetailsParam
	param := ListSourcesParam{
		Title:    "Twitch Sources",
		Subtitle: "Here are the available sources.",
		Source:   "twitch",
	}

	items, err := s.api.Sources.ListBySource("twitch")
	if err != nil {
		panic(err)
	}

	for _, item := range *items {
		var i ListSourcesDetailsParam

		if !item.Enabled {
			i.Disabled = true
			i.Enabled = false
		} else {
			i.Disabled = false
			i.Enabled = true
		}

		i.Item = item

		details = append(details, i)
	}
	param.Items = &details

	w.Header().Add("Content Type", "text/html")
	err = s.templates.ExecuteTemplate(w, "sources.list", param)
	if err != nil {
		panic(err)
	}
}

// This is the form that lets you enter a new Twitch source into the application
// /settings/sources/twitch/new
func (s *HttpServer) SourcesTwitchNewForm(w http.ResponseWriter, r *http.Request) {
	var details []ListSourcesDetailsParam
	param := ListSourcesParam{
		Title:    "Twitch Sources",
		Subtitle: "Here are the available sources.",
		Source:   "twitch",
	}

	items, err := s.api.Sources.ListBySource("twitch")
	if err != nil {
		panic(err)
	}

	for _, item := range *items {
		var i ListSourcesDetailsParam

		if !item.Enabled {
			i.Disabled = true
			i.Enabled = false
		} else {
			i.Disabled = false
			i.Enabled = true
		}

		i.Item = item

		details = append(details, i)
	}
	param.Items = &details

	w.Header().Add("Content Type", "text/html")
	err = s.templates.ExecuteTemplate(w, "sources.new.twitch", param)
	if err != nil {
		panic(err)
	}
}

// This validates the infomation sent from the form and passes it to the API.
func (s *HttpServer) SourcesTwitchNewFormPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Form Error",
			Code:  500,
			Error: err,
		})
		return
	}

	name := r.Form.Get("name")
	if name == "" {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Missing Name value",
			Code:  500,
			Error: err,
		})
		return
	}

	err = s.api.Sources.NewTwitch(name)
	if err != nil {
		s.templates.ExecuteTemplate(w, "err", ErrParam{
			Title: "Failed to add new YouTube source",
			Code:  500,
			Error: err,
		})
		return
	}

	w.Header().Add("Content Type", "text/html")
	err = s.templates.ExecuteTemplate(w, "sources.posted", HttpParam{
		Title:    "Source Added",
		Subtitle: "It will be reviewed on the next collection",
	})
	if err != nil {
		panic(err)
	}
}


// This displays all the ffxiv sources known to the app.
//
// /settings/sources/ffxiv
func (s *HttpServer) SourcesFfxivIndex(w http.ResponseWriter, r *http.Request) {
	var details []ListSourcesDetailsParam
	param := ListSourcesParam{
		Title:    "Final Fantasy XIV Sources",
		Subtitle: "Here are the available sources.",
		Source:   "ffxiv",
	}

	items, err := s.api.Sources.ListBySource("ffxiv")
	if err != nil {
		panic(err)
	}

	for _, item := range *items {
		var i ListSourcesDetailsParam

		if !item.Enabled {
			i.Disabled = true
			i.Enabled = false
		} else {
			i.Disabled = false
			i.Enabled = true
		}

		i.Item = item

		details = append(details, i)
	}
	param.Items = &details

	w.Header().Add("Content Type", "text/html")
	err = s.templates.ExecuteTemplate(w, "sources.list", param)
	if err != nil {
		panic(err)
	}
}

func (s *HttpServer) ListSources(w http.ResponseWriter, r *http.Request) {
	param := ListSourcesParam{}
	var details []ListSourcesDetailsParam

	items, err := s.api.Sources.List()
	if err != nil {
		panic(err)
	}

	for _, item := range *items {
		var i ListSourcesDetailsParam

		if !item.Enabled {
			i.Disabled = true
			i.Enabled = false
		} else {
			i.Disabled = false
			i.Enabled = true
		}

		i.Item = item

		details = append(details, i)
	}
	param.Items = &details

	w.Header().Add("Content Type", "text/html")
	err = s.templates.ExecuteTemplate(w, "sources.list", param)
	if err != nil {
		panic(err)
	}
}

func (s *HttpServer) GetSourceById(w http.ResponseWriter, r *http.Request) {
	var details []ListSourcesDetailsParam
	param := ListSourcesParam{}

	items, err := s.api.Sources.List()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	for _, item := range *items {
		var i ListSourcesDetailsParam

		if !item.Enabled {
			i.Disabled = true
			i.Enabled = false
		} else {
			i.Disabled = false
			i.Enabled = true
		}

		i.Item = item

		details = append(details, i)
	}
	param.Items = &details

	w.Header().Add("Content Type", "text/html")
	err = s.templates.ExecuteTemplate(w, "sources.list", param)
	if err != nil {
		panic(err)
	}
}
