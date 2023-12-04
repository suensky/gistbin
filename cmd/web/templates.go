package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/suensky/gistbin/internal/models"
)

type templateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
	Form        any
	Flash       string
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:06")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		// V1: hardcoded base files
		// files := []string{
		// 	"./ui/html/base.html",
		// 	"./ui/html/partials/nav.html",
		// 	page,
		// }

		// V2: parse all files under partials
		t, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.html")
		if err != nil {
			return nil, err
		}
		t, err = t.ParseGlob("./ui/html/partials/*.html")
		if err != nil {
			return nil, err
		}

		t, err = t.ParseFiles(page)
		if err != nil {
			return nil, err
		}
		cache[name] = t
	}

	return cache, nil
}
