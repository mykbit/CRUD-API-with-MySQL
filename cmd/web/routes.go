package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/vinyl/view", app.vinylView)
	mux.HandleFunc("/vinyl/add", app.vynilAdd)
	mux.HandleFunc("/vinyl/edit", app.vinylUpdate)
	mux.HandleFunc("/vinyl", app.vinylDelete).Methods("DELETE")

	return mux
}
