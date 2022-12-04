package web

import (
	"embed"
	"html/template"
)

//go:embed *
var files embed.FS

func parse(file string) *template.Template {
	temp := template.Must(template.New("layout.html").ParseFS(files, "layout.html", file))
	return temp
}

// This will load layout, requested template, and Articles menu
func parseArticles(file string) *template.Template {
	temp := template.Must(template.New("layout.html").ParseFS(files, "layout.html", "templates/articles/menu.html", file))
	return temp
}

func parseSettings(file string) *template.Template {
	temp := template.Must(template.New("layout.html").ParseFS(files, "layout.html", "templates/settings/menu.html", file))
	return temp
}
