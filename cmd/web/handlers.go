package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// defining this as a method for our application struct
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	// template.ParseFiles() to read template file into a set
	// we can pass it as a variadic parameter
	ts, err := template.ParseFiles(files...)
	// either relative to current directory or an absolute path
	// this is relative obviously
	if err != nil {
		app.serverError(w, err)
		return
	}

	// then Execute() the template set to write it as response body
	// the last parameter is any dynamic data we want to pass in
	// for now we'll leave it as nil
	err = ts.ExecuteTemplate(w, "base", nil)
	// assigning to err to check if there's an error
	// using ExecuteTemplate because we have multiple templates
	// that template invokes other templates
	if err != nil {
		app.serverError(w, err)
	}

	w.Write([]byte("hello world from me!"))
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	fmt.Fprintf(w, "displaying snippet with id %d", id)

}
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	// some vars holding dummy data
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := 7

	// pass the data to the SnippetModel.Insert() method
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// redirect the user to the relevant page for the snippet
	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)

	// w.Write([]byte("create a new snippet"))
}
