package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /item/create", itemCreate)
	mux.HandleFunc("POST /item/create", itemCreatePost)
	mux.HandleFunc("GET /item/view/{id}", itemViewId)

	logger.Info("starting server", slog.String("addr", *addr))
	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
