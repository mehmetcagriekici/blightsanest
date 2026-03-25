package cmd

import (
        "log"
	"time"
	"context"

	"github.com/spf13/cobra"
        amqp "github.com/rabbitmq/amqp091-go"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
	"github.com/mehmetcagriekici/blightsanest/internal/pubsub"
	"github.com/mehmetcagriekici/blightsanest/internal/routing"
)

var saveCryptoCmd = &cobra.Command{
	Use: "crypto",
	Short: "Publish the current crypto list to other clients.",
	Run: handleCryptoSave,
}

func handleCryptoSave(cmd *cobra.Command, args []string) {
        defer log.Print("> ")

        if len(CryptoState.CurrentList) == 0 {
	        log.Println("Current client list is empty. Nothing to save...")
		return
	}

        // save current list to the cache
	CryptoCache.Add(CryptoState.CurrentListID, CryptoState.CurrentList)

        // publish current list to the other clients
	data := routing.CryptoExchangeBody{
	        ID:        CryptoState.CurrentListID,
		CreatedAt: time.Now(),
		Payload:   CryptoState.CurrentList,
	}

        log.Printf("Publishing the list: %s\n", CryptoState.CurrentListID)
	if err := pubsub.PublishClientCrypto(Ctx, Conn, data); err != nil {
	        log.Fatal(err)
	}

	log.Println("List is successfully published.")
	return
}
