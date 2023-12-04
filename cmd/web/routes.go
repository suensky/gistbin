package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	// Creat router.
	router := httprouter.New()

	// Customized error handler
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	// Static file server
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	// Dynamic routes
	dynamic := alice.New(app.sessionManager.LoadAndSave)

	// Routing
	router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.home))
	router.Handler(http.MethodGet, "/gist/view/:id", dynamic.ThenFunc(app.gistView))
	router.Handler(http.MethodGet, "/gist/create", dynamic.ThenFunc(app.gistCreate))
	router.Handler(http.MethodPost, "/gist/create", dynamic.ThenFunc(app.gistCreatePost))

	// Middleware registry
	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	return standard.Then(router)
}
