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
	dynamicMiddleware := alice.New(app.session.Enable,app.authenticate)
	mux := pat.New()

	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	mux.Get("/snippet/create", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.createSnippetForm))
	mux.Post("/snippet/create", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.createSnippet))
	mux.Get("/user/signup", dynamicMiddleware.ThenFunc(app.signupUserForm))
	mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.signupUser))
	mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm))
	mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser))
	mux.Post("/user/logout", dynamicMiddleware.ThenFunc(app.logoutUser))
	mux.Get("/snippet/:id",dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.showSnippet))

	//For loading the css and other static data in html script templates we send these data over HTTP using fileServer
	// and below is its handler
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
