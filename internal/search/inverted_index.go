package search

import(
	"os"
	"bufio"
	"context"

	"github.com/mehmetcagriekici/blightsanest/internal/pubsub"
	"github.com/mehmetcagriekici/blightsanest/internal/database"
)

// Implemented Assets:
//   crypto:
type InvertedIndex struct {
	// token -> doc_id -> doc
	Index     map[string]map[string]struct{}
	// doc_id -> doc -> bytes can be unmarshalled
	SearchMap map[string][]byte
}

// initiate inverted index
func NewInvertedIndex() *InvertedIndex {
	return &InvertedIndex{
		Index: make(map[string]map[string]struct{}),
		SearchMap: make(map[string][]byte),
	}
}

// save index and the docmap to the disk
func (i *InvertedIndex) SaveDocuments() (int, int, error) {
	// path to index
	pIdx := "cache/db_index.gob"
	// path to docmap
	pDoc := "cache/db_docmap.gob"

	// create the cache folder
	if err := os.MkdirAll("cache", 0750); err != nil {
		return 0, 0, err
	}

	// encode index and docmap (searchmap)
	encodedIndex, err := pubsub.Encode(i.Index)
	if err != nil {
		return 0, 0, err
	}

	encodedSearchMap, err := pubsub.Encode(i.SearchMap)
	if err != nil {
		return 0, 0, err
	}

	// create index file
	fIdx, err := os.Create(pIdx)
	if err != nil {
		return 0, 0, err
	}
	defer fIdx.Close()
	
	// create a new writer for the index and write encoded index to the file
	wIdx := bufio.NewWriter(fIdx)
	nIdx, err := wIdx.Write(encodedIndex)
	if err != nil {
		return 0, 0, err
	}

	// create docmap file
	fDoc, err := os.Create(pDoc)
	if err != nil {
		return 0, 0, err
	}
	defer fDoc.Close()

	// create a new writer for the docmap and write encoded searchmap to the file
	wDoc := bufio.NewWriter(fDoc)
	nDoc, err := wDoc.Write(encodedSearchMap)
	if err != nil {
		return nIdx, 0, err
	}

	return nIdx, nDoc, nil
}

// tokenize the input text, add each token to the index with the document ID
func (i *InvertedIndex) AddDocument(docID, text string) {
	tokens := Tokenize(text)
	i.Index[docID] = tokens
}

// get the set of document ids for a a given query
func (i *InvertedIndex) GetDocuments(query string) []string {
	// tokenize the query
	tokens := Tokenize(query)

	// array to store doc_ids
	results := []string{}

	// iterate over the index and append the doc_ids that contains the token to the results
	for k, v := range i.Index {
		// iterate over the tokens
		for _, t := range tokens {
			if _, ok := v[t]; ok {
				results = append(results, k)
			}
		}
	}

	return results
}

func (i *InvertedIndex) BuildCryptoIndex(ctx context.Context, queries *database.Queries) error {
	// get the entire crypto data from the database
	cryptoData, err := queries.GetAllCrypto(ctx)
	if err != nil {
		return err
	}

	// iterate over the entire data
	for crypto in cryptoData {
		dataKey := crypto.CryptoKey

		// convert json.RawMessage to []byte
		cryptoBytes, err = crypto.CryptoList.MarshalJSON()
		if err != nil {
			return err
		}

		// stringify the bytes
		content := string(cryptoBytes)

		// add docs to index
		i.AddDocument(dataKey, content)

		// add docs to docmap
		i.SearchMap[dataKey] = cryptoBytes
	}
}
