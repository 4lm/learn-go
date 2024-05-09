package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /item/create", itemCreate)
	mux.HandleFunc("POST /item/create", itemCreatePost)
	mux.HandleFunc("GET /item/view/{id}", itemViewId)

	const port string = ":4000"
	log.Print("http://localhost" + port)
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
