package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// writes an error message and stack trace to the errorLog
// sends 500 internal server error to the user
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	// we use debug.Stack to get a stack trace
	// helpful when trying to debug errors
	app.errorLog.Println(trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// specific status code and message to the user
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// basic notFound helper
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) render(w http.ResponseWriter, status int, page string, data *templateData) {
	// retrieve appropriate template set from cache
	// if no provided name, send server error
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, err)
		return
	}

	w.WriteHeader(status)
	// 200 ok, 400 bad request etc.

	// execute the template set and write to w
	err := ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, err)
	}
}
