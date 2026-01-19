package search

import(
	"context"

	"github.com/mehmetcagriekici/blightsanest/internal/crypto"
	"github.com/mehmetcagriekici/blightsanest/internal/database"
)

// Implemented Assets:
//   crypto
type SearchData{
}

type InvertedIndex struct {
	Index     map[string]map[int]struct{}
	SearchMap map[int]SearchData
}
