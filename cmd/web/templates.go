package main

import "snippetbox.yigitcankurtaran.net/internal/models"

// define a templateData type to act as a holding structure for any dynamic data that we want to pass to our HTML templates
type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
