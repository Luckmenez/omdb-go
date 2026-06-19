package main

import (
	"imdb-api/api"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	if err := Run(); err != nil {
		slog.Error("failed to run server", "error", err)
		os.Exit(1)
	}
	slog.Info("server started on port 3333")
}

func Run() error {
	apikey := os.Getenv("OMDB_KEY")
	handler := api.NewHandler(apikey)

	s := http.Server{
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: time.Minute,
		Addr: ":3333",
		Handler: handler,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}