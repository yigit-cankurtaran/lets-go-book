package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	// creating our file server for static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// using Handle() to register the file server as the handler for
	// all URL paths that start with "/static/"
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	// we can serve single files with http.ServeFile()
	// but that's more unsafe and less efficient

	// using the struct methods as handler funcs
	mux.HandleFunc("/", app.home)
	// HandleFunc transforms a function into a handler
	// and registers it in the same step
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return mux
}
