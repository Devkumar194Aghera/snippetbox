package main

import (
	"fmt"
	"goModule/snippetbox/pkg/models"
	"net/http"
	"runtime/debug"
)

//We do this for easy retrieval of the user details from the request context.

func (app *application) authenticatedUser(r *http.Request) *models.User {
	user, ok := r.Context().Value(contextKeyUser).(*models.User)
	if !ok {
		return nil
	}
	return user
}

// The serverError helper writes an error message and stack trace to the errorLo
// then sends a generic 500 Internal Server Error response to the user.

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())

	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// The clientError helper sends a specific status code and corresponding descri
// to the user. We'll use this later in the book to send responses like 400 "Bad
// Request" when there's a problem with the request that the user sent.

func (app *application) clientError(w http.ResponseWriter, status int) {

	http.Error(w, http.StatusText(status), status)

}

// For consistency, we'll also implement a notFound helper. This is simply a
// convenience wrapper around clientError which sends a 404 Not Found response
// the user.

func (app *application) notFound(w http.ResponseWriter) {

	app.clientError(w, http.StatusNotFound)

}

// func (app *application) isAuthnticated(r *http.Request) *models.User {

// 	return app.session.GetInt(r, "userID")
// }
