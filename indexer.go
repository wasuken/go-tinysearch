package tinysearch

import (
	"bufio"
	"io"
)

type Indexer struct {
	index     *Index
	tokenizer *Tokenizer
}

func NewIndexer(tokenizer *Tokenizer) *Indexer {
	return &Indexer{
		index:     NewIndex(),
		tokenizer: tokenizer,
	}
}

func (idxr *Indexer) update(docID DocumentID, reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(idxr.tokenizer.SplitFunc)
	var position int

	for scanner.Scan() {
		term := scanner.Text()

		if postingsList, ok := idxr.index.Dictionary[term]; !ok {
			idxr.index.Dictionary[term] = NewPostingsList(NewPosting(docID, position))
		} else {
			postingsList.Add(NewPosting(docID, position))
		}
		position++
	}
	idxr.index.TotalDocsCount++
}
