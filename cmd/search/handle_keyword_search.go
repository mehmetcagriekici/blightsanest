package main

import(
	"fmt"
	"log"
	"slices"
	"strings"
	
	"github.com/mehmetcagriekici/blightsanest/internal/search"
	"github.com/mehmetcagriekici/blightsanest/internal/clientlogic"
)

type scorePair struct {
	doc_id string;
	score float64;
}

func handle_keyword_search(invertedIndex *search.InvertedIndex, assetType string, queryArray []string) {
	// load the cdocuments from the cache
	if err := invertedIndex.LoadDocuments(); err != nil {
		log.Println("Before start searching, please build the inverted index with the command <create_inverted_index>")
		log.Fatal(err)
	}

	query := strings.Join(queryArray, " ")
	// tokenize the query
	queryTokens := search.Tokenize(query)

	// map document ids to their total bm25 scores
	scores := make(map[string]float64)

	// iterate over the query tokens
	for _, t := range queryTokens {
		// get the relevant doc_ids from the index
		for _, d := range invertedIndex.GetDocuments(t) {
			// calculate the bm25 score
			if _, ok := scores[d]; !ok {
				scores[d] = 0.0
			}
			newScore, err := search.CalcBM25(invertedIndex, d, t)
			if err != nil {
				log.Fatal(err)
			}
			scores[d] += newScore
		}
	}

	// sort the scores
	sortedScores := []scorePair{}
	for k, v := range scores {
		if (v > 0.0) {
			pair := scorePair{
				doc_id: k,
				score: v,
			}
			sortedScores = append(sortedScores, pair)
		}
	}
	slices.SortFunc(sortedScores, func(a, b scorePair) int {
		return int(b.score - a.score)
	})

	// display the search results
	for i, v := range sortedScores {
		if i == 5 {
			break
		}

		if assetType == clientlogic.ASSET_CRYPTO {
			fmt.Printf("Cryto List ID: %s\n Match Score: %f\n", v.doc_id, v.score)
		}
	}
}
