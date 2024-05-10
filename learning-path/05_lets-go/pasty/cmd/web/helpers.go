package main

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
		trace  = string(debug.Stack())
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri, "trace", trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func createLogger(logFmt *string) *slog.Logger {
	options := &slog.HandlerOptions{
		AddSource: true,
	}
	var loggerHandler slog.Handler
	if *logFmt == "json" {
		loggerHandler = slog.NewJSONHandler(os.Stdout, options)
	} else {
		loggerHandler = slog.NewTextHandler(os.Stdout, options)
	}
	logger := slog.New(loggerHandler)
	return logger
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
