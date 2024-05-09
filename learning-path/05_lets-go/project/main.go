package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	const port string = ":4000"

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /item/create", itemCreate)
	mux.HandleFunc("POST /item/create", itemCreatePOST)
	mux.HandleFunc("GET /item/view/{id}", itemView)

	log.Print("http://localhost" + port)

	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home"))
}

func itemCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home / Item / Create"))
}

func itemCreatePOST(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home / Item / Create POST: saving new item..."))
}

func itemView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprintf("Home / Item / View / %d", id)
	w.Write([]byte(msg))
}
