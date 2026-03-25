package cmd

import(
	"log"

	"github.com/spf13/cobra"
	amqp "github.com/rabbitmq/amqp091-go"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
	"github.com/mehmetcagriekici/blightsanest/internal/pubsub"
	"github.com/mehmetcagriekici/blightsanest/internal/routing"
)

var fetchCryptoCmd = &cobra.Command{
	Use:   "crypto [args...]",
	Short: "Fetch a crypto list from the server."
	Args:  cobra.MinimumNArgs(1),
	Run:   handleCryptoFetch,
}

// gets data from the server
func handleCryptoFetch(cmd *cobra.Command, args []string) {
	defer log.Print("> ")

        // control args
	if len(args) != 1 {
	        log.Println("<fetch crypto> command requires a key of a published crypto list as an argument.")
		log.Println("    fetch crypto <id_of_a_published_crypto_list_from_the_server>")
		return
	}

	key := args[0]

        // check if the requested list is the current list
	if key == CryptoState.CurrentListID {
	        log.Println("Requested crypto list is already the list on the current client. Didn't perform the fetch request to the server.")
		return
	}

        // check client cache if the crypto list exists
	if _, ok := CryptoCache.Get(key); ok {
	        log.Println("Requested crypto list already exists in the client cache.")
		log.Println("To make the requested list the current client list:")
		log.Printf("    switch crypto %s\n", key)
		return
	}

	cancel, err := pubsub.SubscribeCrypto(Conn, func(delivery routing.CryptoExchangeBody) routing.AckType {
	        log.Println("Subscribing to the server crypto channel to get the requested list...")

	        list := delivery.Payload
		id := delivery.ID

                if id == "" || len(list) == 0 {
		        log.Println("No crypto list is delivered to the client subscriber.")
			return routing.NACK_DISCARD
		}

                CryptoCache.Add(id, list)
		CryptoState.UpdateCurrentList(id, list)

		log.Printf("New crypto list <%s> is successfully added to the client cache and the state.\n", id)
		return routing.ACK
	})

	if err != nil {
	        log.Fatal(err)
	}

        SubManager.Add(cancel)
	return
}
