package main

import (
	"log"
	"net/http"
)

// home handler function writing a byte slice
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// if we don't want / to be a catch all
		http.NotFound(w, r)
		// requires both of this for some reason
		return
	}
	w.Write([]byte("Hello from me!"))
	// byte slice because it only accepts byte slice
	// fundamental type in Go, directly represents binary data

	// regular Go func with 2 parameters
	// responsewriter provides methods to construct a response
	// *http.Request is a pointer to a struct that holds all the information about the request

	// this will be used for our home route
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("display a specific snippet"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		// if method is not post
		// first is the header name second is the header value
		w.Header().Set("Allow", "POST")
		// send 405 and write it out
		w.WriteHeader(405)
		// we can only call this once, more logs a warning
		// we have to put this first because anything else sends a 200 OK code
		w.Write([]byte("method not allowed"))
		return
	}

	w.Write([]byte("create a new snippet"))
}

func main() {
	// fmt.Println("Hello, World!")
	// basic hello world

	// we need a handler
	// a handler is a function that handles a request
	// then a router
	// a router stores mapping between a request and a handler
	// then a server
	// in go we can establish a server and listen to a port in application

	mux := http.NewServeMux()
	// mux is a router
	// creating this explicitly is good
	// we don't HAVE TO but it's a lot safer and more secure
	mux.HandleFunc("/", home)
	// servemux treats this as a catch all
	// any path that doesn't match any other handler will be sent to this handler
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	// fixed paths bc they don't end with a /
	// the URL has to match this exactly
	// "/" is a subtree path because it ends in a slash
	// something like "/static/" would also be an example
	// we can think of it like a wild card. such as "/static/**"

	// ListenAndServe to start a new server
	// 2 parameters, port and our servemux
	// if there's an error use log.Fatal() to log and exit

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	// if we don't specify a port it will listen on all interfaces
	log.Fatal(err)

}
