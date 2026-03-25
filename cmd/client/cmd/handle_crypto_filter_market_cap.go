package cmd

import(
        "log"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

var filterCryptoMarketCapCmd = &cobra.Command{
	Use:   "crypto market_cap [args...]",
	Short: "Filter coins using min and max market cap.",
	Run:   handleCryptoFilterMarketCap,
}

func handleCryptoFilterMarketCap(cmd *cobra.Command, args []string) {
        controlFilterMarketCap(CryptoState, args)
        list := crypto.FilterCoinCap(CryptoState.CurrentMinMarketCap,
		CryptoState.CurrentMaxMarketCap,
		CryptoState.CurrentList)
	newID := fmt.Sprintf("filter_market_cap_%f_%f", CryptoState.CurrentMinMarketCap, CryptoState.CurrentMaxMarketCap)
	fields := []string{"MarketCapRank", "MarketCapChangePercentage"}
	commonCryptoHandler(CryptoState, list, fields, newID)
}

// min market cap
// max market cap
func controlFilterMarketCap(cs *crypto.CryptoState, args []string) {
        switch len(args) {
	case 0:
	        log.Printf("No arguments are passed, using the client state values for min and max market caps: %f, %f\n", cs.CurrentMinMarketCap, cs.CurrentMaxMarketCap)
	case 1:
	        log.Printf("1 argument is passed. Using the user input as the min market cap value. Max market cap value is fetched from the client state. Max market cap: %f\n", cs.CurrentMaxMarketCap)
		minCap, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
		        log.Fatal(err)
		}
		cs.UpdateMarketCap(minCap, cs.CurrentMaxMarketCap)
	case 2:
	        log.Println("All arguments are passed. Using the passed arguments as the min and max market caps...")
	        minCap, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
		        log.Fatal(err)
		}
		maxCap, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
		        log.Fatal(err)
		}
		cs.UpdateMarketCap(minCap, maxCap)
	default:
	        log.Println("Invalid amount of arguments: filter crypto market_cap <min_market_cap float64> <max_market_cap float64>")
	}
}
