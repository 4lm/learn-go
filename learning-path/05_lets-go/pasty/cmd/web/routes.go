package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /item/create", app.itemCreate)
	mux.HandleFunc("POST /item/create", app.itemCreatePost)
	mux.HandleFunc("GET /item/view/{id}", app.itemViewId)

	return mux
}
