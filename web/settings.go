package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jtom38/newsbot/portal/api"
)

const (
	FFXIVSourceName   = "ffxiv"
	RedditSourceName  = "reddit"
	TwitchSourceName  = "twitch"
	YoutubeSourceName = "youtube"
)

var (
	pageError = parse("templates/err.html")

	pageSettingIndex           = parseSettings("templates/settings/index.html")
	pageSourceUpdated          = parseSettings("templates/settings/posted.html")
	pageSettingSourcesList     = parseSettings("templates/settings/sources/list.html")
	pageSettingsNewRedditForm  = parseSettings("templates/settings/sources/new-reddit.html")
	pageSettingsNewTwitchForm  = parseSettings("templates/settings/sources/new-twitch.html")
	pageSettingsNewYouTubeForm = parseSettings("templates/settings/sources/new-youtube.html")

	pageSettingsDiscordWebhooksList = parseSettings("templates/settings/outputs/discordwebhooks/list.html")
	pageSettingsDiscordWebhooksForm = parseSettings("templates/settings/outputs/discordwebhooks/new.html")
)

type SettingsRouter struct {
	_api api.CollectorApi
}

func NewSettingsRouter(api *api.CollectorApi) SettingsRouter {
	c := SettingsRouter{
		_api: *api,
	}
	return c
}

func (s *SettingsRouter) GetRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", s.SettingsIndex)

	r.Post("/sources/disable", s.DisableSourceById)
	r.Post("/sources/enable", s.EnableSourceById)

	r.Get("/sources/reddit", s.ListReddit)
	r.Get("/sources/reddit/new", s.NewRedditForm)
	r.Post("/sources/reddit/new", s.NewRedditPost)

	r.Get("/sources/youtube", s.ListYoutube)
	r.Get("/sources/youtube/new", s.NewYouTubeForm)
	r.Post("/sources/youtube/new", s.NewYouTubePost)

	r.Get("/sources/twitch", s.ListTwitch)
	r.Get("/sources/twitch/new", s.NewTwitchForm)
	r.Post("/sources/twitch/new", s.NewTwitchPost)

	r.Get("/sources/ffxiv", s.ListFfxiv)

	r.Get("/outputs/discord/webhooks", s.ListDiscordWebHooks)
	r.Get("/outputs/discord/webhooks/new", s.NewDiscordWebHooksForm)
	r.Post("/outputs/discord/webhooks/new", s.NewDiscordWebhookPost)

	return r
}

func (s SettingsRouter) SettingsIndex(w http.ResponseWriter, r *http.Request) {
	param := TitlesParam{
		Title:    "Configuration",
		Subtitle: "It doesn't do anything on its own",
	}
	pageSettingIndex.Execute(w, param)
}

type UpdateSourceParam struct {
	Title    string
	Subtitle string
	IsError  bool
	Message  string
	Code     int
}

// /settings/sources/enable?id
func (s *SettingsRouter) EnableSourceById(w http.ResponseWriter, r *http.Request) {
	param := UpdateSourceParam{
		Title:    "Source was not enabled",
		Subtitle: "See error for details.",
		IsError:  true,
		Code:     500,
	}

	err := r.ParseForm()
	if err != nil {
		param.Message = err.Error()
		pageError.Execute(w, param)
		return
	}

	id := r.Form.Get("id")
	if id == "" {
		param.Message = "Missing Source ID"
		pageError.Execute(w, param)
		return
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		param.Message = err.Error()
		pageError.Execute(w, param)
		return
	}

	err = s._api.Sources().Enable(uid)
	if err != nil {
		param.Message = err.Error()
		pageError.Execute(w, err)
		return
	}

	param = UpdateSourceParam{
		Title:    "Source was enabled",
		Subtitle: "Head on back to see the change.",
	}
	if pageSourceUpdated.Execute(w, param); err != nil {
		log.Print(err)
	}
}

// /settings/sources/disable?id
func (s *SettingsRouter) DisableSourceById(w http.ResponseWriter, r *http.Request) {
	param := UpdateSourceParam{
		Title:    "Source was not disabled",
		Subtitle: "See error for details.",
		IsError:  true,
		Code:     500,
		Message:  "",
	}

	err := r.ParseForm()
	if err != nil {
		param.Message = err.Error()
		pageError.Execute(w, param)
		return
	}

	id := r.Form.Get("id")
	if id == "" {
		param.Message = err.Error()
		pageError.Execute(w, param)
		return
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		param.Message = err.Error()
		pageError.Execute(w, param)
		return
	}

	err = s._api.Sources().Disable(uid)
	if err != nil {
		param.Message = err.Error()
		pageError.Execute(w, param)
		return
	}

	param = UpdateSourceParam{
		Title:    "Source was disabled",
		Subtitle: "Head on back to see the change",
	}
	if pageSourceUpdated.Execute(w, param); err != nil {
		log.Print(err)
	}
}

type ListSettingsParam struct {
	Title      string
	Subtitle   string
	Items      *[]api.Source
	SourceName string
	IsError    bool
	Error      string
}

func (s SettingsRouter) ListReddit(w http.ResponseWriter, r *http.Request) {
	param := ListSettingsParam{
		Title:      "Known Subreddits",
		Subtitle:   "Here you can see the available sources to pick from ",
		SourceName: RedditSourceName,
		IsError:    false,
		Error:      "",
	}

	items, err := s._api.Sources().ListBySource(RedditSourceName)
	if err != nil {
		param.IsError = true
		param.Error = err.Error()
	}

	param.Items = items

	if pageSettingSourcesList.Execute(w, param); err != nil {
		log.Print(err)
	}
}

func (s SettingsRouter) ListYoutube(w http.ResponseWriter, r *http.Request) {
	param := ListSettingsParam{
		Title:      "Known YouTube Channels",
		Subtitle:   "Here you can see the available sources to pick from ",
		SourceName: YoutubeSourceName,
		IsError:    false,
		Error:      "",
	}

	items, err := s._api.Sources().ListBySource(YoutubeSourceName)
	if err != nil {
		param.IsError = true
		param.Error = err.Error()
	}

	param.Items = items

	if pageSettingSourcesList.Execute(w, param); err != nil {
		log.Print(err)
	}
}

func (s SettingsRouter) ListTwitch(w http.ResponseWriter, r *http.Request) {
	param := ListSettingsParam{
		Title:      "Known Twitch Streamers",
		Subtitle:   "Here you can see the available sources to pick from ",
		SourceName: TwitchSourceName,
		IsError:    false,
		Error:      "",
	}

	items, err := s._api.Sources().ListBySource(TwitchSourceName)
	if err != nil {
		param.IsError = true
		param.Error = err.Error()
	}

	param.Items = items

	if pageSettingSourcesList.Execute(w, param); err != nil {
		log.Print(err)
	}
}

func (s SettingsRouter) ListFfxiv(w http.ResponseWriter, r *http.Request) {
	param := ListSettingsParam{
		Title:      "Known Final Fantasy XIV regions",
		Subtitle:   "Here you can see the available sources to pick from",
		SourceName: FFXIVSourceName,
		IsError:    false,
		Error:      "",
	}

	items, err := s._api.Sources().ListBySource(FFXIVSourceName)
	if err != nil {
		param.IsError = true
		param.Error = err.Error()
	}

	param.Items = items

	if pageSettingSourcesList.Execute(w, param); err != nil {
		log.Print(err)
	}
}

type NewSourceParam struct {
	Title      string
	Subtitle   string
	SourceName string
}

func (s SettingsRouter) NewRedditForm(w http.ResponseWriter, r *http.Request) {
	var err error
	param := NewSourceParam{
		Title:      "Create a new Reddit monitor",
		Subtitle:   "",
		SourceName: RedditSourceName,
	}
	if pageSettingsNewRedditForm.Execute(w, param); err != nil {
		log.Print(err)
	}
}

// This handles the data from the form and sends it to the API
func (s SettingsRouter) NewRedditPost(w http.ResponseWriter, r *http.Request) {
	param := ErrorParam{
		Title:    "Failed to add a new Reddit source",
		Subtitle: "See the error for details.",
		Code:     500,
	}
	err := r.ParseForm()
	if err != nil {
		param.Error = err.Error()
		pageError.Execute(w, param)
		return
	}

	name := r.Form.Get("name")
	if name == "" {
		param.Error = "Subreddit name was missing from the form"
		pageError.Execute(w, param)
		return
	}

	uri := fmt.Sprintf("https://reddit.com/r/%v", name)
	err = s._api.Sources().NewReddit(name, uri)
	if err != nil {
		param.Error = err.Error()
		pageError.Execute(w, param)
		return
	}

	p := TitlesParam{
		Title:    "New Reddit source was added",
		Subtitle: "Head on back to see the update",
	}

	if pageSourceUpdated.Execute(w, p); err != nil {
		log.Print(err)
	}
}

func (s SettingsRouter) NewTwitchForm(w http.ResponseWriter, r *http.Request) {
	var err error
	param := NewSourceParam{
		Title:      "Create a new Twitch monitor",
		Subtitle:   "",
		SourceName: TwitchSourceName,
	}
	if pageSettingsNewTwitchForm.Execute(w, param); err != nil {
		log.Print(err)
	}
}

func (s SettingsRouter) NewTwitchPost(w http.ResponseWriter, r *http.Request) {
	param := ErrorParam{
		Title:    "Failed to add a new Twitch source",
		Subtitle: "See the error for details.",
		Code:     500,
	}
	err := r.ParseForm()
	if err != nil {
		param.Error = err.Error()
		pageError.Execute(w, param)
		return
	}

	name := r.Form.Get("name")
	if name == "" {
		param.Error = "Subreddit name was missing from the form"
		pageError.Execute(w, param)
		return
	}

	err = s._api.Sources().NewTwitch(name)
	if err != nil {
		param.Error = err.Error()
		pageError.Execute(w, param)
		return
	}

	p := TitlesParam{
		Title:    "New Twitch source was added",
		Subtitle: "Head on back to see the update",
	}

	if pageSourceUpdated.Execute(w, p); err != nil {
		log.Print(err)
	}
}

func (s SettingsRouter) NewYouTubeForm(w http.ResponseWriter, r *http.Request) {
	var err error
	param := NewSourceParam{
		Title:      "Create a new YouTube monitor",
		Subtitle:   "",
		SourceName: YoutubeSourceName,
	}
	if pageSettingsNewYouTubeForm.Execute(w, param); err != nil {
		log.Print(err)
	}
}

func (s SettingsRouter) NewYouTubePost(w http.ResponseWriter, r *http.Request) {
	param := ErrorParam{
		Title:    "Failed to add a new YouTube source",
		Subtitle: "See the error for details.",
		Code:     500,
	}
	err := r.ParseForm()
	if err != nil {
		param.Error = err.Error()
		pageError.Execute(w, param)
		return
	}

	name := r.Form.Get("name")
	if name == "" {
		param.Error = "Channel name was missing from the form."
		pageError.Execute(w, param)
		return
	}

	url := r.Form.Get("url")
	if url == "" {
		param.Error = "URL name was missing from the form."
		pageError.Execute(w, param)
		return
	}

	err = s._api.Sources().NewYouTube(name, url)
	if err != nil {
		param.Error = err.Error()
		pageError.Execute(w, param)
		return
	}

	p := TitlesParam{
		Title:    "New YouTube source was added",
		Subtitle: "Head on back to see the update",
	}

	if pageSourceUpdated.Execute(w, p); err != nil {
		log.Print(err)
	}
}

type ListOutputDiscordWebHooks struct {
	Title string
	Subtitle string
	Items *[]api.Discordwebhook
}

func (s SettingsRouter) ListDiscordWebHooks(w http.ResponseWriter, r *http.Request) {
	param := ListOutputDiscordWebHooks{
		Title:      "Discord WebHooks",
		Subtitle:   "Here you can see the available sources to pick from",
	}

	items, err := s._api.Outputs().DiscordWebHook().List()
	if err != nil {
		e := ErrorParam {
			Title: "Failed to collect Discord Web Hooks",
			Subtitle: "See error for details",
			Code: 500,
			Error: err.Error(),
		}
		errorPage.Execute(w, e)
		return
	}

	param.Items = items

	if pageSettingsDiscordWebhooksList.Execute(w, param); err != nil {
		log.Print(err)
	}
}

func (s SettingsRouter) NewDiscordWebHooksForm(w http.ResponseWriter, r *http.Request) {
	var err error
	param := TitlesParam {
		Title: "New Discord Webhook",
	}
	if pageSettingsDiscordWebhooksForm.Execute(w, param); err != nil {
		log.Print(err)
	}
}

func (s SettingsRouter) NewDiscordWebhookPost(w http.ResponseWriter, r *http.Request) {
	param := ErrorParam{
		Title:    "Failed to add a new YouTube source",
		Subtitle: "See the error for details.",
		Code:     500,
	}
	err := r.ParseForm()
	if err != nil {
		param.Error = err.Error()
		pageError.Execute(w, param)
		return
	}

	server := r.Form.Get("server")
	if server == "" {
		param.Error = "Server name was missing from the form."
		pageError.Execute(w, param)
		return
	}

	url := r.Form.Get("url")
	if url == "" {
		param.Error = "URL name was missing from the form."
		pageError.Execute(w, param)
		return
	}

	channel := r.Form.Get("channel")
	if channel == "" {
		param.Error = "Channel was missing from the form."
		pageError.Execute(w, param)
		return
	}

	err = s._api.Outputs().DiscordWebHook().New(server, channel, url)
	if err != nil {
		param.Error = err.Error()
		pageError.Execute(w, param)
		return
	}

	p := TitlesParam{
		Title:    "New Discord Web Hook",
		Subtitle: "Head on back to see the update",
	}

	if pageSourceUpdated.Execute(w, p); err != nil {
		log.Print(err)
	}
}