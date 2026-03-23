package cmd

import (
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func handleQuit (cmd *cobra.Command, args []string) {
	log.Println("Ending client session...")
	SubManager.CloseAll()
	time.Sleep(200 * time.Millisecond)
	os.Exit(0)
}
