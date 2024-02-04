package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"word-search-in-files/pkg/searcher"
)

type SearchHandler struct {
	Searcher *searcher.Searcher
}

func (s *SearchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	word := r.URL.Query().Get("word")
	if word == "" {
		http.Error(w, "word is empty", http.StatusBadRequest)
		return
	}

	files, err := s.Searcher.Search(word)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error in searching files %v", err), http.StatusInternalServerError)
		return
	}

	resultJSON, err := json.Marshal(files)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error marshaling file %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write(resultJSON)
}
