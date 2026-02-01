package search

import(
	"math"
	"errors"
)

// Diminishing returns controller constant
var BM25_K1 float64 = 1.5
// length normalization constant
var BM25_B float64 = 0.75

// term frequency - inverse document frequency
// TF * IDF - bm25 score
func CalcBM25(invertedIndex *invertedIndex, docID, token string) (float64, error) {
	idf, err := CalcIDF(invertedIndex, term)
	if err != nil {
		return 0.0, err
	}

	tf, err := invertedIndex.GetTf(docID, token)
	if err != nil {
		return 0.0, err
	}

	// calculate the average doc length and length norm
	avgDocLen := 0.0
	lengthNorm := 1 - BM25_B
	if len(invertedIndex.DocLengths) > 0 {
		for _, v := range invertedIndex.DocLengths {
			avgDocLen += v
		}
		avgDocLen = avgDocLen / len(invertedIndex.DocLengths)
		lengthNorm += BM25_B * (invertedIndex.DocLengths[docID] / avgDocLen) 
	}

	// saturate and normalize the tf
	saturatedTf := (tf * (BM25_K1 + 1)) / (tf + BM25_K1 * lengthNorm)

	// bm25 score
	return float64(saturatedTf) * idf
}

// inverse document frequency
func CalcIDF(invertedIndex *InvertedIndex, token string) (float64, error) {
	// load the index and the docmap from the inverted index from the disk
	if err := invertedIndex.Load(); err != nil {
		return 0.0, err
	}

	// BM25 IDF solution
	totalDocCount := float64(len(invertedIndex.DocMap))
	termMatchDocCount := float64(len(invertedIndex.Index[token]))
	return math.Log((totalDocCount - termMatchDocCount + 0.5) / (termMatchDocCount + 0.5 ) + 1)
}
