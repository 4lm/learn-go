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

	logger.Info("starting server", slog.String("addr", *addr))
	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
