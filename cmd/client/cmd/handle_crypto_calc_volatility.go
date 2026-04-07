package cmd

import(
        "log"
	"strconv"
	"fmt"

	"github.com/spf13/cobra"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

var calcCryptoVolatilityCmd = &cobra.Command{
	Use:   "crypto volatility [args...]",
	Short: "Calculate coin volatilities and get coins in a range of min/max volatility preference.",
	Run:   handleCryptoCalcVolatility,
}

func handleCryptoCalcVolatility(cmd *cobra.Command, args []string) {
        controlCalcVolatility(CryptoState, args)
        list := crypto.CalcCoinVolatility(CryptoState.CurrentMinVolatility,
		CryptoState.CurrentMaxVolatility,
		CryptoState.CurrentList)
	newID := fmt.Sprintf("calc_volatility_%f_%f",
		CryptoState.CurrentMinVolatility,
		CryptoState.CurrentMaxVolatility)
        fields := []string{"High24H", "Low24H"}
	commonCryptoHandler(CryptoState, list, fields, newID)
}

func controlCalcVolatility(cs *crypto.CryptoState, args []string) {
        switch len(args) {
	case 0:
	        log.Println("No arguments are passed. Using the current client preferences for min and max volatility values.")
		log.Println("volatility = (high_24h - lwo_24h) / current_price")
		log.Printf("Current Min Volatility: %f\n", cs.CurrentMinVolatility)
		log.Printf("Current Max Volatility: %f\n", cs.CurrentMaxVolatility)
	case 1:
	        log.Println("Only one argument is passed. Updating the client min volatility preference...")
 		log.Printf("Current Max Volatility: %f\n", cs.CurrentMaxVolatility)
 		minVolatility, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
		        log.Fatal(err)
		}
		cs.UpdateCurrentVolatility(minVolatility, cs.CurrentMaxVolatility)
	case 2:
	        log.Println("Updating the client min and max volatility preferences...")
		minVolatility, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
		        log.Fatal(err)
		}
		maxVolatility, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
		        log.Fatal(err)
		}
		cs.UpdateCurrentVolatility(minVolatility, maxVolatility)
	default:
	        log.Println("Invalid use of command: calc crypto volatility <min_volatility float64> <max_volatility float64>")
	}
}
