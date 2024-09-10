package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/oliversabler/goth-blueprint/handlers"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stderr, nil))

	defaultHandler := handlers.NewDefaultHandler(log)
	defaultHandler.Handle()

	server := &http.Server{
		Addr:         "localhost:3000",
		Handler:      nil,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Error("Error listening: %v.", slog.Any("error", err))
	}
}
