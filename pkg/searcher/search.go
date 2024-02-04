package searcher

import (
	"fmt"
	"io/fs"
	"strings"
	"sync"
	"word-search-in-files/pkg/internal/dir"
)

type Searcher struct {
	FS fs.FS
}

type SearchResult struct {
	FileName string
}

func (s *Searcher) Search(word string) (files []string, err error) {
	filenames, err := dir.FilesFS(s.FS, ".")
	if err != nil {
		return nil, err
	}
	//fmt.Println(filenames)

	var wg sync.WaitGroup

	resultCh := make(chan SearchResult, len(filenames))

	for _, filename := range filenames {
		wg.Add(1)
		go func(filename string) {
			defer wg.Done()
			content, readerr := fs.ReadFile(s.FS, filename)
			if readerr != nil {
				//return nil, readerr
				fmt.Printf("Error reading file %s: %v\n", filename, readerr)
				return
			}

			fileContent := string(content)

			if strings.Contains(fileContent, word) {
				//files = append(files, filename)
				resultCh <- SearchResult{FileName: filename}
			}
		}(filename)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for result := range resultCh {
		files = append(files, result.FileName)
	}

	return files, nil
}
