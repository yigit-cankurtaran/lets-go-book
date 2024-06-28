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
	// initialize map to act as the cache
	cache := map[string]*template.Template{}

	// use filepath.Glob to get a slice of all file paths in the pages
	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	// loop through the pages slice
	for _, page := range pages {
		// get file name and assign it to name
		name := filepath.Base(page)

		// create slice for filepaths
		files := []string{
			"./ui/html/base.tmpl.html",
			"./ui/html/partials/nav.tmpl.html",
			page,
		}

		// parse files into template set
		ts, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		// add the template set to the map using name as key
		cache[name] = ts
	}

	// return the map
	return cache, nil
}
