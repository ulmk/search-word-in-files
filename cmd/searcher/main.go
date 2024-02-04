package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"word-search-in-files/pkg/handlers"
	"word-search-in-files/pkg/searcher"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting cwd:", err)
		return
	}
	examplesDir := filepath.Join(cwd, "examples")
	fs := os.DirFS(examplesDir)
	searcher := &searcher.Searcher{FS: fs}
	//files, _ := searcher.Search("Палашка")
	//fmt.Println(files)
	searchHandler := &handlers.SearchHandler{
		Searcher: searcher,
	}

	http.Handle("/files/search", searchHandler)
	fmt.Println("Server is listening on :8080")
	http.ListenAndServe(":8080", nil)
}
