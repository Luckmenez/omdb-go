package Omdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Result struct {
	Search       []SearchResult `json:"Search"`
	TotalResults string   `json:"totalResults"`
	Response     string   `json:"Response"`
}

type SearchResult struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

func Search(apiKey string, title string) (Result, error) {
	var v url.Values

	v.Set("apiKey", apiKey)
	v.Set("s", title)

	resp, err := http.Get("https://www.omdbapi.com/?" + v.Encode())

	if err != nil {
		return Result{}, fmt.Errorf("failed to get search results: %w", err)
	}
	
	defer resp.Body.Close()

	var result Result

	if err :=json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return Result{}, fmt.Errorf("Error to decode JSOn %w", err)
	}

	return result, nil
}