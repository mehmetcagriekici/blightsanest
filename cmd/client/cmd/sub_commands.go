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
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch a published asset from the server.",
}

// get
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a published asset from other clients.",
}

// rank
var rankCmd = &cobra.Command{
	Use:   "rank",
	Short: "Rank assets by existing fields.",
}

// group
var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Group assets using existing features."
}

// filter
var filterCmd = &cobra.Command{
	Use: "filter",
	Short: "Filter assets for existing features."
}

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

	fetchCmd.AddCommand(fetchCryptoCmd)

	getCmd.AddCommand(getCryptoCmd)

	rankCmd.AddCommand(rankCryptoCmd)

	groupCmd.AddCommand(groupCryptoLiquidityCmd, groupCryptoScarcityCmd)

	filterCmd.AddCommand(filterCryptoTotalVolumeCmd,
		filterCryptoMarketCapCmd,
		filterCryptoPriceChangePercentageCmd,
		filterCryptoVolatileCmd,
		filterCryptoHighRiskCmd,
		filterCryptoLowRiskCmd,
	)

	RootCmd.AddCommand(quitCmd,
		databaseCmd,
		switchCmd,
		setCmd,
		saveCmd,
		listCmd,
		fetchCmd,
		getCmd,
		rankCmd,
		groupCmd,
		filterCmd,
	)
}
