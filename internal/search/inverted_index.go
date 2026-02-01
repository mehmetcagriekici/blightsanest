package search

import(
	"os"
	"fmt"
	"context"
	"errors"

	"github.com/mehmetcagriekici/blightsanest/internal/database"
	"github.com/mehmetcagriekici/blightsanest/internal/readwrite"
)

// Implemented Assets:
//   crypto:
type InvertedIndex struct {
	// token -> doc_id -> {}
	Index           map[string]map[string]struct{}
	// index path
	PIDX            string
	// doc_id -> doc -> bytes can be unmarshalled
	DocMap       map[string][]byte
	// docmap path
	PDOC            string
	// term frequencies - how many times each term appears in each document
	// doc_id: {term: count}
	TermFrequencies map[string]map[string]int
	// term frequencies cache path
	PTF             string
	// Document Lengths
	DocLengths      map[string]int
	// Doc Lengths Cache Path
	PDL             string
}

// initiate inverted index
func NewInvertedIndex() *InvertedIndex {
	return &InvertedIndex{
		Index: make(map[string]map[string]struct{}),
		DocMap: make(map[string][]byte),
		TermFrequencies: make(map[string]map[string]int),
		DocLengths: make(map[string]int),
		PIDX: "../../cache/db_index.gob",
		PDOC: "../../cache/db_docmap.gob",
		PTF: "../../cache/db_termfreq.gob",
		PDL: "../../cache/db_doclens.gob",
	}
}

// function the get the term frequencies of a document for a term
func (i *InvertedIndex) GetTf(docID, term string) (int, error) {
	// tokenize the term
	tokens := Tokenize(term)
	// if there is more than one term return an error
	if len(tokens) != 1 {
		return 0, errors.New("Get term frequencies requires one single term.")
	}

	countObj, ok := i.TermFrequencies[doc_id]
	if !ok {
		return 0, nil
	}

	return i.TermFrequencies[doc_id][tokens[0]], nil
}

// load index and the docmap from the disk
func (i *InvertedIndex) LoadDocuments() error {
	// read index file
	bufIdx, err := readwrite.Read(i.PIDX)
	if err != nil {
		return err
	}
	
	// decode the read index
	decodedIdx, err := readwrite.Decode[map[string]map[string]struct{}](bufIdx)
	if err != nil {
		return err
	}
	
	// read docmap file
	bufDoc, err := readwrite.Read(i.PDOC)
	if err != nil {
		return err
	}
	// decode the read doc map
	decodedDoc, err := readwrite.Decode[map[string][]byte](bufDoc)
	if err != nil {
		return err
	}

	// read term frequencies file
	bufTf, err := readwrite.Read(i.PTF)
	if err != nil {
		return err
	}
	// decode the term frequencies
	decodedTf, err := readwrite.Decode[map[string]map[string]int](bufTf)
	if err != nil {
		return err
	}

	// read doc lengths file
	bufDocl, err := readwrite.Read(i.PDL)
	if err != nil {
		return err
	}
	// decode doc lengths
	decodedDocl, err := readwrite.Decode[map[string]int](bufDocl)
	if err != nil {
		return err
	}

	// assign inverted index
	i.Index = decodedIdx
	i.DocMap = decodedDoc
	i.TermFrequencies = decodedTf
	i.DocLengths = decodedDocl
	
	return nil
}

// save index and the docmap to the disk
func (i *InvertedIndex) SaveDocuments() error {
	// create the cache folder
	if err := os.MkdirAll("../../cache", 0750); err != nil {
		return err
	}

	// encode index, docmap (searchmap) and term frequencies
	encodedIndex, err := readwrite.Encode(i.Index)
	if err != nil {
		return err
	}

	encodedDocMap, err := readwrite.Encode(i.DocMap)
	if err != nil {
		return err
	}

	encodedTermFrequencies, err := readwrite.Encode(i.TermFrequencies)
	if err != nil {
		return err
	}

	encodedDocLengths, err := readwrite.Encode(i.Doclengths)
	if err != nil {
		return err
	}

        // create index file and write the current index
	nIdx, err := readwrite.Write(i.PIDX, encodedIndex)
	if err != nil {
		return err
	}

	fmt.Printf("%d bytes are written for inverted index index.\n", nIdx)
	
	// create docmap file and write the docmap
	nDoc, err := readwrite.Write(i.PDOC, encodedDocMap)
	if err != nil {
		return err
	}

	fmt.Printf("%d bytes are written for inverted index docmap.\n", nIDpc)
	
	// create term frequencies file and write term frequencies
	nTf, err := readwrite.Write(i.PTF, encodedTermFrequencies)
	if err != nil {
		return err
	}

	fmt.Printf("%d bytes are written for inverted index term frequencies.\n", nTf)

	// create document lengths life and write doc lengths
	nDocl, err := readwrite.Write(i.PDL, encodedDocLengths)
	if err != nil {
		return err
	}

	fmt.Printf("%d bytes are written for inverted index document lengths.\n", nDocl)

	return nil
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

		// check if there is a count object for the doc id
		if _, ok := i.TermFrequencies[doc_id]; !ok {
			i.TermFrequencies[doc_id] = make(map[string]int)
		}

		// check if the token exist in the document counter
		if _, ok := i.TermFrequencies[doc_id][t]; !ok {
			i.TermFrequencies[doc_id][t] = 0
		}
		i.TermFrequencies[doc_id][t] += 1
	}

	// store document length
	i.DocLengths[doc_id] = len(tokens)
}

// get the set of document ids for a a given token
func (i *InvertedIndex) GetDocuments(token string) []string {
	// tokenize the query
	tokens := Tokenize(query)

	// array to store doc_ids
	results := []string{}

	// iterate over the index and append the doc_ids that contains the token to the results
	for k, v := range i.Index {
		if _, ok := v[token]; ok {
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
		i.DocMap[dataKey] = cryptoBytes
	}

	return nil
}
