package tinysearch

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

type IndexReader struct {
	indexDir      string
	postingsCache map[string]*PostingsList
	docCountCache int
}

func NewIndexReader(path string) *IndexReader {
	cache := make(map[string]*PostingsList)
	return &IndexReader{path, cache, -1}
}

func (r *IndexReader) postingsLists(terms []string) []*PostingsList {
	postingsLists := make([]*PostingsList, 0, len(terms))
	for _, term := range terms {
		if postings := r.postings(term); postings != nil {
			postingsLists = append(postingsLists, postings)
		}
	}
	return postingsLists
}

func (r *IndexReader) postings(term string) *PostingsList {
	if postingsList, ok := r.postingsCache[term]; ok {
		return postingsList
	}
	filename := filepath.Join(r.indexDir, term)
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil
	}
	var postingsList PostingsList
	err = json.Unmarshal(bytes, &postingsList)
	if err != nil {
		return nil
	}

	r.postingsCache[term] = &postingsList
	return &postingsList
}

func (r *IndexReader) totalDocCount() int {
	if r.docCountCache > 0 {
		return r.docCountCache
	}
	filename := filepath.Join(r.indexDir, "_0.dc")
	file, err := os.Open(filename)
	if err != nil {
		return 0
	}
	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil {
		return 0
	}
	count, err := strconv.Atoi(string(bytes))
	if err != nil {
		return 0
	}
	r.docCountCache = count
	return count

}
