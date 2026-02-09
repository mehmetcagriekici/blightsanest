package search

// fastapi types

type EmbeddingDoc struct {
	ID   string `json:"id"`
	Data string `json:"data"`
}

type EmbeddingsRequest struct {
	Documents []EmbeddingDoc `json:"documents"`
}

type EmbeddingsResponse struct {
	Count int `json:"count"`
}

type SearchRequest struct {
	Query string `json:"query"`
	Limit int    `json:"limit"`
}

type SearchDocument struct {
	ID   string `json:"id"`
	Data string `json:"data"`
}

type SearchResult struct {
	Score    float64        `json:"score"`
	Document SearchDocument `json:"document"`
}

type SearchResponse struct {
	Results []SearchResult `json:"results"`
}
