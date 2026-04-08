package cmd

import (
	"github.com/spf13/cobra"
)

// fetch
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch asset data from third party resources.",
}

// get
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Fetch assets from the database.",
}

// save
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save assets to the database.",
}

// delete
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete assets from the database",
}

func init() {
	fetchCmd.AddCommand(fetchCryptoCmd)

	getCmd.AddCommand(getCryptoCmd)

	saveCmd.AddCommand(saveCryptoCmd)

	deleteCmd.AddCommand(deleteCryptoCmd)

	RootCmd.AddCommand(
		fetchCmd,
		getCmd,
		saveCmd,
		deleteCmd,
	)
}
