package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	logFmt := flag.String("logFmt", "text", "\"text\" or \"json\" logs")
	flag.Parse()

	logger := createLogger(logFmt)

	app := &application{
		logger: logger,
	}

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /item/create", app.itemCreate)
	mux.HandleFunc("POST /item/create", app.itemCreatePost)
	mux.HandleFunc("GET /item/view/{id}", app.itemViewId)

	logger.Info("starting server", slog.String("addr", *addr))
	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
