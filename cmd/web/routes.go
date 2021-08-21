package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	mux := http.NewServeMux()     // create new routing thing
	mux.HandleFunc("/", app.home) // add routes
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))         // get files in static location
	mux.Handle("/static/", http.StripPrefix("/static", fileServer)) // StripPrefix for safety. Register new handler with ServerMux

	return standardMiddleware.Then(mux)
}
