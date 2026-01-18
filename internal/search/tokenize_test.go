package search

import(
	"log"
	"slices"
	"testing"
)

func TestTokenize(t *testing.T) {
	query := "Show coins with at least 1234.567 dollars current price"

	expected := []string{"show", "coin", "least", "1234.567", "dollar", "current", "price"}
	tokens := Tokenize(query)
	log.Println(tokens)

	if !slices.Equal(tokens, expected) {
		t.Errorf("Expected: %v \n Got: %v \n", expected, tokens)
	}
}
