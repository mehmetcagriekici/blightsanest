package cmd

import(
	"log"
	"context"

	"github.com/spf13/cobra"

	"github.com/mehmetcagriekici/blightsanest/internal/crypto"
	"github.com/mehmetcagriekici/blightsanest/internal/database"
)

// create
var createDatabaseCmd = &cobra.Command{
	Use:   "create crypto",
	Short: "Create a row for a crypto list on the database",
	Run:    handleCryptoDatabaseCreate,
}

func handleCryptoDatabaseCreate(cmd *cobra.Command, args []string) {
	log.Printf("Saving the crypto list on this client with the ID %s to the database.\n", CryptoState.CurrentListID)
	if err := crypto.CreateCryptoRow(Ctx, DbQueries, CryptoState.CurrentList, CryptoState.CurrentListID); err != nil {
		log.Fatal(err)
	}
}

// read
var readDatabaseCmd = &cobra.Command{
	Use:   "read crypto [args...]",
	Short: "Get a crypto list from the database",
	Args:  cobra.MinimumNArgs(1),
	Run:   handleCryptoDatabaseRead
}

func handleCryptoDatabaseRead(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Println("Please provide the ID of the crypto list you want to get from the database.")
		return
	}

	log.Printf("Getting the crypto list from the database with the ID %s...\n", args[0])
        list, err := crypto.ReadCryptoRow(Ctx, DbQueries, args[0])
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("New Client List: %s\n", args[0])
	CryptoState.UpdateCurrentList(args[0], list)

}

// update
var updateDatabaseCmd = &cobra.Command{
	Use:   "update crypto [args...]",
	Short: "Update a crypto list in the database with another crypto list.",
	Args:  cobra.MinimumNArgs(1),
	Run:   handleCryptoDatabaseUpdate,
}

func handleCryptoDatabaseUpdate(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Println("Please provide the ID of the database you want to update from the database.")
		return
	}

	var currList []crypto.MarketData
	// check if the list is the current state list
	if CryptoState.CurrentListID == args[0] {
		// update the list at the database with the client state list
		log.Printf("Updating the list %s at the database with the current list.\n", args[0])
		currList = CryptoState.CurrentList
	} else {
		// check if the list exists on the client cache
		if entry, ok := CryptoCache.Get(args[0]); ok {
			// update the list at the database with the list from the client cache
			log.Printf("Updating the list %s at the database with an existing list from the client cache", args[0])
			currList = entry.Market
		} else {
			log.Printf("Provided list ID %s is not the current list ID neither exists on the client cache. Exiting the process...\n", args[0])
			return
		}
	}

	updatedList, err := crypto.UpdateCryptoRow(Ctx, currList, DbQueries, args[0], args[0])
	if err != nil {
		log.Fatal(err)
	}

	CryptoState.UpdateCurrentList(args[0], updatedList)
}

// delete
var deleteDatabaseCmd = &cobra.Command{
	Use:   "delete crypto [args...]",
	Short: "Delete selected crypto lists from the database",
	Args:  cobra.MinimumNArgs(1),
	Run:   handleCryptoDatabaseDelete,
}

func handleCryptoDatabaseDelete(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Println("This is a manual process, please provide the IDs of the crypto lists that you want to delete from the database.")
		return
	}
	for _, k := range args {
		if err := crypto.DeleteCryptoRow(Ctx, DbQueries, k); err != nil {
			log.Fatal(err)
		}
	}
}
