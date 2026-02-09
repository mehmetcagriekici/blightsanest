package main

import(
	"log"
	"context"
	"encoding/json"
	
	"github.com/mehmetcagriekici/blightsanest/internal/search"
	"github.com/mehmetcagriekici/blightsanest/internal/database"
)

func handle_create_semantic_index(ctx context.Context, queries *database.Queries, client *search.Client) {
	// get all the crypto data from the database
	// []model.Crypto 
	cryptoData, err := queries.GetAllCrypto(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// docs to be sent
	docs := []search.EmbeddingDoc{}
	for _, v := range cryptoData {
		doc := search.EmbeddingDoc{
			ID: v.CryptoKey,
			Data: rawToStr(v.CryptoList),
		}
		docs = append(docs, doc)
	}

	// create semantic index from the client
	n, err := client.Index(docs);
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Semantic Index created with %d embeddings", n.Count)
}

func rawToStr(raw json.RawMessage) string {
	bytes, err := raw.MarshalJSON()
	if err != nil {
		log.Fatal(err)
	}
	
	return string(bytes[:])
}
