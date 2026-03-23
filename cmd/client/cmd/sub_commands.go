package cmd

import (
	"github.com/spf13/cobra"
)

// root subcommands
// quit - from handle_quit.go
var quitCmd = &cobra.Command{
	Use:   "quit",
	Short: "End the client session",
	Run:   handleQuit,
}

// database
var databaseCmd = &cobra.Command{
	Use:   "database",
	Short: "Database interaction for assets",
}

// switch
var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Switch between asset instances",
}

// set
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Configure an asset aspect.",
}

// save
var saveCmd = &cobra.Command{
	Use: "save",
	Short: "Publish the asset to other clients.",
}

// list
var listCmd = &cobra.Command{
	Use: "list",
	Short: "Print the list of existing asset ids",
}

// fetch

func init() {
	databaseCmd.AddCommand(createDatabaseCmd,
		readDatabaseCmd,
		updateDatabaseCmd,
		deleteDatabaseCmd,
	)

	switchCmd.AddCommand(switchCryptoCmd)

	setCmd.AddCommand(setCryptoCmd)

	saveCmd.AddCommand(saveCryptoCmd)

	listCmd.AddCommand(listCryptoCmd)

	RootCmd.AddCommand(quitCmd,
		databaseCmd,
o		switchCmd,
		setCmd,
		saveCmd,
		listCmd,
	)
}
