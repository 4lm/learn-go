package main

import (
	"log"
	"net/http"
)

func main() {
	const port string = ":4000"

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /item/create", itemCreate)
	mux.HandleFunc("POST /item/create", itemCreatePost)
	mux.HandleFunc("GET /item/view/{id}", itemViewId)

	log.Print("http://localhost" + port)

	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
