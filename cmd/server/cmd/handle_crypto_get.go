package cmd

import (
	"encoding/json"
	"log"

	"github.com/spf13/cobra"

	"github.com/mehmetcagriekici/blightsanest/internal/crypto"
	"github.com/mehmetcagriekici/blightsanest/internal/pubsub"
	"github.com/mehmetcagriekici/blightsanest/internal/routing"
)

var getCryptoCmd = &cobra.Command{
	Use:   "crypto [args...]",
	Short: "Get crypto data from the database.",
	Run:   handleCryptoGet,
}

func handleCryptoGet(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		log.Printf("Please provide the ID of the crypto list you want to get from the database.")
		return
	}

	// check if requested list already exists in the cache
	if _, ok := CryptoCache.Get(args[0]); ok {
		log.Println("Requested list already exists in the current cache. Quiting the process...")
		return
	}

	data, err := DbQueries.GetCryptoList(Ctx, args[0])
	if err != nil {
		log.Fatal(err)
	}

	// decode the crypto list
	mb, err := data.CryptoList.MarshalJSON()
	if err != nil {
		log.Fatal(err)
	}
	var b []crypto.MarketData
	if err := json.Unmarshal(mb, &b); err != nil {
		log.Fatal(err)
	}

	// add new list to the cache
	CryptoCache.Add(data.CryptoKey, b)

	// publish the new list from the server
	delivery := routing.CryptoExchangeBody{
		ID:        data.CryptoKey,
		CreatedAt: data.CreatedAt,
		Payload:   b,
	}
	if err := pubsub.PublishCrypto(Ctx, Conn, delivery); err != nil {
		log.Fatal(err)
	}
}
