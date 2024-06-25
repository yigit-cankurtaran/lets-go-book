package main

// cmd will contain the app-specific code for the execs
// in the project. we'll have 1 exec so far, will live here

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// command line flag named addr, default is ":4000"
	// and some short text to explain what it controls
	// having these values hardcoded isn't ideal
	addr := flag.String("addr", ":4000", "HTTP network address")
	// we use String because if the val can't be an str
	// it will log an error and exit
	flag.Parse()
	// we need to call Parse before we use the addr var

	mux := http.NewServeMux()

	// creating our file server for static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// using Handle() to register the file server as the handler for
	// all URL paths that start with "/static/"
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	// we can serve single files with http.ServeFile()
	// but that's more unsafe and less efficient

	mux.HandleFunc("/", home)
	// HandleFunc transforms a function into a handler
	// and registers it in the same step
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Printf("starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
