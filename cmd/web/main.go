package main

// cmd will contain the app-specific code for the execs
// in the project. we'll have 1 exec so far, will live here

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// creating our file server for static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// using Handle() to register the file server as the handler for
	// all URL paths that start with "/static/"
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	// we can serve single files with http.ServeFile()
	// but that's more unsafe and less efficient

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
