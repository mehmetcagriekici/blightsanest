package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

var calcCryptoLiquidityCmd = &cobra.Command{
	Use:   "crypto liquidity [args...]",
	Short: "Calculate coin liquidities and get coins with min liquidity preference.",
	Run:   handleCryptoCalcLiquidity,
}

func handleCryptoCalcLiquidity(cmd *cobra.Command, args []string) {
	controlCalcLiquidity(CryptoState, args)
	list := crypto.CalcCoinLiquidity(CryptoState.CurrentMinLiquidity, CryptoState.CurrentList)
	newID := fmt.Sprintf("calc_liquidity_%f", CryptoState.CurrentMinLiquidity)
	fields := []string{"TotalVolume", "MarketCap", "MarketCapRank"}
	commonCryptoHandler(CryptoState, list, fields, newID)
}

func controlCalcLiquidity(cs *crypto.CryptoState, args []string) {
	switch len(args) {
	case 0:
		log.Println("No arguments are passed. Using the client preferences for min liquidity preference...")
		log.Printf("Min Liquidity: %f\n", cs.CurrentMinLiquidity)
	case 1:
		log.Println("Updating the client min liquidity preference...")
		minLiquidity, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			log.Fatal(err)
		}
		cs.UpdateCurrentLiquidity(minLiquidity)
	default:
		log.Println("Invalid use of command: calc crypto liquidity <min_liquidity float64>")
	}
}
