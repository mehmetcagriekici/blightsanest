package cmd

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mehmetcagriekici/blightsanest/internal/search"
)

type pair struct {
	docID string
	score float64
}

func handleSearch(cmd *cobra.Command, args []string) {
	// convert limit to the integer
	limitN, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}

	// join the query array into a query string
	query := strings.Join(args, " ")

	keywordResults, err := helpKeywordSearch(InvertedIndex, query)
	if err != nil {
		log.Fatal(err)
	}

	res, err := SemanticClient.Search(query)
	if err != nil {
		log.Fatal(err)
	}
	semanticResults := res.Results

	// calc rrf scores
	rrfScores := make(map[string]float64)

	rank := 1
	for {
		if len(keywordResults) == 0 && len(semanticResults) == 0 {
			break
		}

		// starts at rank 1
		firstKeyword := keywordResults[0]
		firstSemantic := semanticResults[0]

		// check if a document exists in both
		keywordScore, ok := rrfScores[firstKeyword.docID]
		if ok {
			rrfScores[firstKeyword.docID] = keywordScore + calcRRF(rank)
		} else {
			rrfScores[firstKeyword.docID] = calcRRF(rank)
		}

		// check if a document exists in both
		semanticScore, ok := rrfScores[firstSemantic.Document.ID]
		if ok {
			rrfScores[firstSemantic.Document.ID] = semanticScore + calcRRF(rank)
		} else {
			rrfScores[firstSemantic.Document.ID] = calcRRF(rank)
		}
		// increase the current rank
		rank += 1

		// remove the first elements
		keywordResults = keywordResults[1:]
		semanticResults = semanticResults[1:]
	}

	// sort the rrfScores
	sortedScores := []pair{}
	for k, v := range rrfScores {
		sortedScores = append(sortedScores, pair{
			docID: k,
			score: v,
		})
	}

	slices.SortFunc(sortedScores, func(a, b pair) int {
		return int(b.score - a.score)
	})

	for i, v := range sortedScores {
		if i == limitN {
			break
		}

		fmt.Printf("Search Result Score: %f, Search Result ID: %s", v.score, v.docID)
	}
}

func calcRRF(rank int) float64 {
	return float64(1 / (rank + 60))
}

func helpKeywordSearch(invertedIndex *search.InvertedIndex, query string) ([]pair, error) {
	// load the cdocuments from the cache
	if err := invertedIndex.LoadDocuments(); err != nil {
		return nil, err
	}

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

	sortedScores := []pair{}
	for k, v := range scores {
		sortedScores = append(sortedScores, pair{
			docID: k,
			score: v,
		})
	}

	slices.SortFunc(sortedScores, func(a, b pair) int {
		return int(b.score - a.score)
	})

	return sortedScores, nil
}
