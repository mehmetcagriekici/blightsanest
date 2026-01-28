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
	Index           map[string]map[string]struct{}
	// index path
	PIDX            string
	// doc_id -> doc -> bytes can be unmarshalled
	SearchMap       map[string][]byte
	// docmap path
	PDOC            string
	// term frequencies - how many times each term appears in each document
	// doc_id: {term: count}
	TermFrequencies map[string]map[string]int
	// term frequencies cache path
	PTF             string
}

// initiate inverted index
func NewInvertedIndex() *InvertedIndex {
	return &InvertedIndex{
		Index: make(map[string]map[string]struct{}),
		SearchMap: make(map[string][]byte),
		TermFrequencies: make(map[string]map[string]int),
		PIDX: "cache/db_index.gob",
		PDOC: "cache/db_docmap.gob",
		PTF: "cache/db_termfreq.gob",
	}
}

// load index and the docmap from the disk
func (i *InvertedIndex) LoadDocuments() error {
	// open index file
	fIdx, err := os.Open(i.PIDX)
	if err != nil {
		return err
	}
	defer fIdx.Close()

	// index reader
	rIdx := bufio.NewReader(fIdx)

	// start reading the index into the buf index
	bufIdx := make([]byte, 1024)
	for {
		n, err := rIdx.Read(bufIdx)
		if err != nil {
			return err
		}

		if n == 0 {
			break
		}
	}

	// decode the read index
	decodedIdx, err := pubsub.Decode(bufIdx)
	if err != nil {
		return err
	}
	
	// open docmap file
	fDoc, err := os.Open(i.PDOC)
	if err != nil {
		return err
	}
	defer fDoc.Close()

	// docmap reader
	rDoc := bufio.NewReader(fDoc)
	
	// start reading the docmap
	bufDoc := make([]byte, 1024)
	for {
		n, err := rDoc.Read(bufDoc)
		if err != nil {
			return err
		}

		if n == 0 {
			break
		}
	}

	// decode the read doc map
	decodedDoc, err := pubsub.Decode(bufDoc)
	if err != nil {
		return err
	}

	// open term frequencies file
	fTf, err := os.Open(i.PTF)
	if err != nil {
		return err
	}
	defer fTf.Close()

	// term frequencies reader
	rTf := bufio.NewReader(fTf)

	// start reading the term frequencies
	bufTf := make([]byte, 1024)
	for {
		n, err := rTf.Read(bufTf)
		if err != nil {
			return err
		}

		if n == 0 {
			break
		}
	}

	// decode the term frequencies
	decodedTf, err := pubsub.Decode(bufTf)
	if err != nil {
		return err
	}

	// assign inverted index and docmap
	i.Index = decodedIdx
	i.SearchMap = decodedDoc
	i.TermFrequencies = decodedTf
	
	return nil
}

// save index and the docmap to the disk
func (i *InvertedIndex) SaveDocuments() (int, int, error) {
	// create the cache folder
	if err := os.MkdirAll("cache", 0750); err != nil {
		return 0, 0, err
	}

	// encode index, docmap (searchmap) and term frequencies
	encodedIndex, err := pubsub.Encode(i.Index)
	if err != nil {
		return 0, 0, err
	}

	encodedSearchMap, err := pubsub.Encode(i.SearchMap)
	if err != nil {
		return 0, 0, err
	}

	encodedTermFrequencies, err := pubsub.Encode(i.TermFrequencies)
	if err != nil {
		return 0, 0, err
	}

	// create index file
	fIdx, err := os.Create(i.PIDX)
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
	fDoc, err := os.Create(i.PDOC)
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

	// create term frequencies file
	fTf, err := os.Create(i.PTF)
	if err != nil {
		return err
	}
	defer fTf.Close()

	// create a new writer for the term frequencies and write the encoded tf to the file
	wTf := bufio.NewWriter(fTf)
	nTf, err := wTf.Write(encodedTermFrequencies)
	if err != nil {
		return err
	}

	return nIdx, nDoc, nil
}

// tokenize the input text, add each token to the index with the document ID
func (i *InvertedIndex) AddDocument(docID, text string) {
	tokens := Tokenize(text)
	for _, t := range tokens {
		if _, ok := i.Index[t]; !ok {
			i.Index[t] = make(map[string]struct{})
		}
		var st struct{}
		i.Index[t][docID] = st
	}
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
	for _, crypto := range cryptoData {
		dataKey := crypto.CryptoKey

		// convert json.RawMessage to []byte
		cryptoBytes, err := crypto.CryptoList.MarshalJSON()
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

	return nil
}
