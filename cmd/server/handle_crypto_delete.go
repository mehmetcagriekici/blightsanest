package main

import(
	"log"
	"context"

	"github.com/mehmetcagriekici/blightsanest/internal/database"
)

func handleCryptoDelete(ctx context.Context, args []string, queries *database.Queries) {
	if len(args) == 0 {
		log.Println("To delete a crypto list from the database you need to provide the ID of the crypto list...")
		return
	}

	deleted, err := queries.DeleteCryptoList(ctx, args[0])
	if err != nil {
		log.Fatal(err)
	}

	for _, l := range deleted {
		log.Printf("Crypto list %s is successfully deleted.\n", l.CryptoKey)
	}
}
