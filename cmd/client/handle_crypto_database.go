package main

import(
	"log"
	"context"
	
	"github.com/mehmetcagriekici/blightsanest/internal/crypto"
	"github.com/mehmetcagriekici/blightsanest/internal/database"
)

func handleCryptoDatabase(ctx context.Context,
	queries *database.Queries,
	cs  *crypto.CryptoState,
	cc *crypto.CryptoCache,
	args []string) {
	if len(args) == 0 {
		log.Println("Invalid use of database command!")
		log.Println("Example: database crypto delete ...")
		return
	}
	
	switch key := args[0]; key {
	case "create":
		handleCryptoDatabaseCreate(ctx, cs, cc, queries, args[1:])
	case "read":
		handleCryptoDatabaseRead(ctx, cs, queries, args[2:])
	case "update":
		handleCryptoDatabaseUpdate(ctx, cs, cc, queries, args[2:])
	case "delete":
		handleCryptoDatabaseDelete(ctx, queries, args[2:])
	default:
		log.Println("Invalid database operation! Available: CREATE/READ/UPDATE/DELETE")
		return
	}
}

// create
func handleCryptoDatabaseCreate(ctx context.Context,
	cs *crypto.CryptoState,
	cc *crypto.CryptoCache,
	queries *database.Queries,
	args []string) {
	cryptoKey := cs.CurrentListID
	if len(args) == 2 {
		log.Printf("Fetching the crypto list with the passed ID %s from the crypto cache, updating the current crypto list on this client's state.\n", args[0])
		cryptoEntry, ok := cc.Get(args[0])
		if !ok {
			log.Printf("Provided crypto list ID %s does not exist on the client cache!\n", args[0])
			return
		}

		cs.UpdateCurrentList(args[0], cryptoEntry.Market)
		cryptoKey = args[1]
	}

	if len(args) == 1 {
		log.Printf("Fetching the crypto list with the passed ID %s from the crypto cache, updating the current crypto list on this client's state.\n", args[0])
		cryptoEntry, ok := cc.Get(args[0])
		if !ok {
			log.Printf("Provided crypto list ID %s does not exist on the client cache!\n", args[0])
			return
		}

		cs.UpdateCurrentList(args[0], cryptoEntry.Market)
	}

	log.Printf("Saving the crypto list on this client with the ID %s to the database.\n", cryptoKey)
	if err := crypto.CreateCryptoRow(ctx, queries, cs.CurrentList, cryptoKey); err != nil {
		log.Fatal(err)
	}
}

// read
func handleCryptoDatabaseRead(ctx context.Context,
	cs *crypto.CryptoState,
	queries *database.Queries,
	args []string) {
	if len(args) == 0 {
		log.Println("Please provide the ID of the crypto list you want to get from the database.")
		return
	}

	log.Printf("Getting the crypto list from the database with the ID %s...\n", args[0])
        list, err := crypto.ReadCryptoRow(ctx, queries, args[0])
	if err != nil {
		log.Fatal(err)
	}
	
	log.Printf("New Client List: %s\n", args[0])
	cs.UpdateCurrentList(args[0], list)

}

// update
func handleCryptoDatabaseUpdate(ctx context.Context,
	cs *crypto.CryptoState,
	cc *crypto.CryptoCache,
	queries *database.Queries,
	args []string) {
	if len(args) == 0 {
		log.Println("Please provide the ID of the database you want to update from the database.")
		return
	}

	var currList []crypto.MarketData
	// check if the list is the current state list
	if cs.CurrentListID == args[0] {
		// update the list at the database with the client state list
		log.Printf("Updating the list %s at the database with the current list.\n", args[0])
		currList = cs.CurrentList
	} else {
		// check if the list exists on the client cache
		if entry, ok := cc.Get(args[0]); ok {
			// update the list at the database with the list from the client cache
			log.Printf("Updating the list %s at the database with an existing list from the client cache", args[0])
			currList = entry.Market
		} else {
			log.Printf("Provided list ID %s is not the current list ID neither exists on the client cache. Exiting the process...\n", args[0])
			return
		}
	}

	updatedList, err := crypto.UpdateCryptoRow(ctx, currList, queries, args[0], args[0])
	if err != nil {
		log.Fatal(err)
	}
	
	cs.UpdateCurrentList(args[0], updatedList)
}

// delete
func handleCryptoDatabaseDelete(ctx context.Context, queries *database.Queries, args []string) {
	if len(args) == 0 {
		log.Println("This is a manual process, please provide the IDs of the crypto lists that you want to delete from the database.")
		return
	}
	for _, k := range args {
		if err := crypto.DeleteCryptoRow(ctx, queries, k); err != nil {
			log.Fatal(err)
		}
	}
}
