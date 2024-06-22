package main

import (
	"log"
	"net/http"
)

// home handler function writing a byte slice
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from me!"))
	// byte slice because it only accepts byte slice
	// fundamental type in Go, directly represents binary data

	// regular Go func with 2 parameters
	// responsewriter provides methods to construct a response
	// *http.Request is a pointer to a struct that holds all the information about the request

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
	mux.HandleFunc("/", home)

	// ListenAndServe to start a new server
	// 2 parameters, port and our servemux
	// if there's an error use log.Fatal() to log and exit

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
