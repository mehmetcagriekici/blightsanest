package cmd

import (
	"context"
	"encoding/json"
	"log"

	"github.com/spf13/cobra"

	"github.com/mehmetcagriekici/blightsanest/internal/database"
	"github.com/mehmetcagriekici/blightsanest/internal/search"
)

var createCryptoEmbeddingsCmd = &cobra.Command{
	Use:   "crypto",
	Short: "Create embeddings for the crypto assets in the database",
	Run:   handle_create_crypto_embeddings,
}

func handle_create_crypto_embeddings(cmd *cobra.Command, args []string) {
	log.Println("Building the inverted index for the database")
	if err := InvertedIndex.BuildCryptoIndex(ctx, dbQueries); err != nil {
		log.Printf("An error occured while trying to build the inverted index: %v\n", err)
		return

	}

	log.Println("Saving the created inverted index to the local cache folder.")
	if err := InvertedIndex.SaveDocuments(); err != nil {
		log.Printf("Couldn't save the created inverted index to the local machine: %v\n", err)
		return

	}
	// get all the crypto data from the database
	// []model.Crypto
	cryptoData, err := DbQueries.GetAllCrypto(Ctx)
	if err != nil {
		log.Fatal(err)
	}

	// docs to be sent
	docs := []search.EmbeddingDoc{}
	for _, v := range cryptoData {
		doc := search.EmbeddingDoc{
			ID:   v.CryptoKey,
			Data: rawToStr(v.CryptoList),
		}
		docs = append(docs, doc)
	}

	// create semantic index from the client
	n, err := SemanticClient.Index(docs)
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
