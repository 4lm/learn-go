package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	w.Write([]byte("Home / - GET"))
}

func itemCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home / Item / Create - GET"))
}

func itemCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Home / Item / Create - POST"))
}

func itemViewId(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Home / Item / View / %d - GET", id)
}
