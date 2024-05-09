package main

import (
	"log"
	"net/http"
)

func main() {
	const port string = ":4000"

	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/item/create", itemCreate)
	mux.HandleFunc("/item/view", itemView)

	log.Print("http://localhost" + port)

	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Pasty"))
}

func itemCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Item / Create"))
}

func itemView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Item / View"))
}
