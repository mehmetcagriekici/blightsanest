package search

import(
	"math"
)

// Diminishing returns controller constant
var BM25_K1 float64 = 1.5
// length normalization constant
var BM25_B float64 = 0.75

// term frequency - inverse document frequency
// TF * IDF - bm25 score
func CalcBM25(invertedIndex *InvertedIndex, docID, token string) (float64, error) {
	idf, err := CalcIDF(invertedIndex, token)
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
			avgDocLen += float64(v)
		}
		avgDocLen = avgDocLen / float64(len(invertedIndex.DocLengths))
		lengthNorm += BM25_B * (float64(invertedIndex.DocLengths[docID]) / avgDocLen) 
	}

	// saturate and normalize the tf
	saturatedTf := (float64(tf) * (BM25_K1 + 1)) / (float64(tf) + BM25_K1 * lengthNorm)

	// bm25 score
	return float64(saturatedTf) * idf, nil
}

// inverse document frequency
func CalcIDF(invertedIndex *InvertedIndex, token string) (float64, error) {
	// load the index and the docmap from the inverted index from the disk
	if err := invertedIndex.LoadDocuments(); err != nil {
		return 0.0, err
	}

	// BM25 IDF solution
	totalDocCount := float64(len(invertedIndex.DocMap))
	termMatchDocCount := float64(len(invertedIndex.Index[token]))
	return math.Log((totalDocCount - termMatchDocCount + 0.5) / (termMatchDocCount + 0.5 ) + 1), nil
}
