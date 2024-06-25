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
		http.NotFound(w, r)
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
		app.errorLog.Println(err.Error())
		// passing in errorLog from the struct
		http.Error(w, "internal server error", http.StatusInternalServerError)
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
		app.errorLog.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	w.Write([]byte("hello world from me!"))
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "displaying snippet with id %d", id)

}
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("create a new snippet"))
}
