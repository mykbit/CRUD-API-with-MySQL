package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodGet, "/vinyl/view/:id", app.vinylView)
	router.HandlerFunc(http.MethodGet, "/vinyl/add", app.vinylCreate)
	router.HandlerFunc(http.MethodPost, "/vinyl/add", app.vinylCreatePost)
	router.HandlerFunc(http.MethodGet, "/vinyl/edit/:id", app.vinylUpdate)
	router.HandlerFunc(http.MethodPost, "/vinyl/edit/:id", app.vinylUpdatePost)
	router.HandlerFunc(http.MethodDelete, "/vinyl/:id", app.vinylDelete)

	return app.logRequest(secureHeaders(router))
}
