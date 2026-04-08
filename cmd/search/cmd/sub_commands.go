package cmd

import (
	"github.com/spf13/cobra"
)

// search
var searchCmd = &cobra.Command{
	Use:   "search [args...]",
	Short: "Search assets from the database",
	Run:   handleSearch,
}

// create embeddings
var embeddingsCmd = &cobra.Command{
	Use:   "embeddings",
	Short: "Create embeddings before searching for assets.",
}

func init() {
	embeddingsCmd.AddCommand(createCryptoEmbeddingsCmd)

	RootCmd.AddCommand(searchCmd, embeddingsCmd)
}
