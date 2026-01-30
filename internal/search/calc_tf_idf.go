package search

import(
	"math"
	"errors"
)

// Diminishing returns controller constant
var BM25_K1 float64 = 1.5
// Document length normalization
var BM24_B float64 = 0.75

// term frequency - inverse document frequency
// TF * IDF
func CalcTFIDF(invertedIndex *invertedIndex, docID, term string) (float64, error) {
	idf, err := CalcIDF(invertedIndex, term)
	if err != nil {
		return 0.0, err
	}

	tf, err := invertedIndex.GetTf(docID, term)
	if err != nil {
		return 0.0, err
	}

	// saturate the tf
	saturatedTf := (tf * (BM25_K1 + 1) / (tf + BM25_K1))

	return float64(saturatedTf) * idf
}

// inverse document frequency
func CalcIDF(invertedIndex *InvertedIndex, term string) (float64, error) {
	// load the index and the docmap from the inverted index from the disk
	if err := invertedIndex.Load(); err != nil {
		return 0.0, err
	}

	// tokenize the term - there must be one token
	tokens := Tokenize(term)
	if len(tokens) != 1 {
		return 0.0, errors.New("IDF is calculated for one term.")
	}

	// BM25 IDF solution
	totalDocCount := float64(len(invertedIndex.SearchMap))
	termMatchDocCount := float64(len(invertedIndex.Index[tokens[0]]))
	return math.Log((totalDocCount - termMatchDocCount + 0.5) / (termMatchDocCount + 0.5 ) + 1)
}
