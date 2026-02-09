package search

import (
	"fmt"
	"bytes"
	"net/http"
	"encoding/json"
)

type Client struct {
	BaseURL string
	http    *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
		http:    &http.Client{},
	}
}

func (c *Client) Index(docs []EmbeddingDoc) (*EmbeddingsResponse, error) {
	reqBody, err := json.Marshal(EmbeddingsRequest{Documents: docs})
	if err != nil {
		return nil, err
	}

	resp, err := c.http.Post(c.BaseURL+"/embeddings", "application/json", bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("index: bad status %s", resp.Status)
	}

	var out EmbeddingsResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	
	return &out, nil
}

func (c *Client) Search(query string, limit int) (*SearchResponse, error) {
	reqBody, err := json.Marshal(SearchRequest{Query: query, Limit: limit})
	if err != nil {
		return nil, err
	}

	resp, err := c.http.Post(c.BaseURL+"/search", "application/json", bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search: bad status %s", resp.Status)
	}

	var out SearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}
