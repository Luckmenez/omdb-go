package api

import (
	"encoding/json"
	Omdb "imdb-api/omdb"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)
type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func NewHandler(apiKey string) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	r.Get("/", handleSearchMovie(apiKey))
	
	return r
}

func sendJSON(w http.ResponseWriter, resp Response, status int) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("Erro ao fazer o marshall do JSON:", "error", err)
		sendJSON(w, Response{Error: "Something went wrong"}, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("Erro ao escrever o JSON:", "error", err)
	}
}

func handleSearchMovie(apiKey string) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		search := r.URL.Query().Get("s")
		res, err := Omdb.Search(apiKey, search)
		if err != nil {
			sendJSON(w, Response{Error: "somenthing wrong with omdb"}, http.StatusBadGateway)
			return 
		}
		sendJSON(w, Response{Data: res}, http.StatusOK)
	}
}