package main

import (
	"net/http"

	"github.com/justinas/alice" // New import

	"github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {

	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	// dynamicMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	//new Servemux instead of Default servermux to avoid any commplication might
	//create create if 3rd party pakage use DefaultServeMux
	mux := pat.New()
	mux.Get("/", app.session.Enable(http.HandlerFunc(app.home)))
	mux.Get("/snippet", app.session.Enable(http.HandlerFunc(app.showSnippet)))
	mux.Get("/snippet/create", app.session.Enable(http.HandlerFunc(app.createSnippetForm)))
	mux.Post("/snippet/create", app.session.Enable(http.HandlerFunc(app.createSnippet)))
	mux.Get("/snippet/:id", app.session.Enable(http.HandlerFunc(app.showSnippet)))
	// to prevent from making create = :id

	//For loading the css and other static data in html script templates we send these data over HTTP using fileServer
	// and below is its handler
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
