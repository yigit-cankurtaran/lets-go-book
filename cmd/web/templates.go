package main

import (
	"html/template"
	"path/filepath"

	"snippetbox.yigitcankurtaran.net/internal/models"
)

// define a templateData type to act as a holding structure for any dynamic data that we want to pass to our HTML templates
type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.ParseFiles("./ui/html/base.tmpl.html")
		if err != nil {
			return nil, err
		}

		// adding partials
		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl.html")
		if err != nil {
			return nil, err
		}

		// adding page template
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}
		// template set to the map
		cache[name] = ts
	}
	return cache, nil
}
