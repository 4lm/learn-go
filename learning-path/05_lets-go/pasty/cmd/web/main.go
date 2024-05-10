package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	logFmt := flag.String("logFmt", "text", "\"text\" or \"json\" logs")
	flag.Parse()

	logger := createLogger(logFmt)

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

func createLogger(logFmt *string) *slog.Logger {
	var loggerHandler slog.Handler
	if *logFmt == "json" {
		loggerHandler = slog.NewJSONHandler(os.Stdout, nil)
	} else {
		loggerHandler = slog.NewTextHandler(os.Stdout, nil)
	}
	logger := slog.New(loggerHandler)
	return logger
}
