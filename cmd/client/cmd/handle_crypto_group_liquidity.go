package cmd

import (
        "log"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

var groupCryptoLiquidityCmd = &cobra.Command{
	Use:   "crypto liquidity [args...]",
	Short: "Find crypto liquidity using market rank and total volume",
	Run:   handleCryptoGroupLiquidity,
}

func handleCryptoGroupLiquidity(cmd *cobra.Command, args []string) {
        controlLiquidityArguments(CryptoState, args)
        list := crypto.GroupHighLiquidityCoins(CryptoState.CurrentMinRank,
		CryptoState.CurrentMaxRank,
		CryptoState.CurrentMinVolume,
		CryptoState.CurrentList)

	newID := fmt.Sprintf("group_liquidity_%d_%d_%f", CryptoState.CurrentMinRank, CryptoState.CurrentMaxRank, CryptoState.CurrentMinVolume)
	fields := []string{"MarketCapRank", "TotalVolume"}
	commonCryptoHandler(CryptoState, list, fields, newID)
}

// min market rank int
// max market rank int
// min total volume float64
func controlLiquidityArguments(cs *crypto.CryptoState, args []string) {
        switch len(args) {
	case 0:
	        log.Println("No arguments are passed, using the ones from the client state:")
		log.Printf("min market rank: %d\n", cs.CurrentMinRank)
		log.Printf("max market rank: %d\n", cs.CurrentMaxRank)
		log.Printf("min volume: %f\n", cs.CurrentMinVolume)
		log.Println("")
		log.Println("To use new arguments: group crypto liquidity <min_market_rank int> <max_market_rank int> <min_volume float64>")
	case 1:
	        log.Println("Only one argument is passed. Using the new min market rank value. Getting the rest from the client state.")
		log.Printf("max market rank: %d\n", cs.CurrentMaxRank)
		log.Printf("min volume: %f\n", cs.CurrentMinVolume)
		log.Println("")

		minRank, err := strconv.Atoi(args[0])
		if err != nil {
		         log.Fatal(err)
		}

		log.Println("Updating the client state, min market cap value...")
		cs.UpdateMarketRank(minRank, cs.CurrentMaxRank)
	case 2:
	        log.Println("Two arguments are passed. Using the new min and max market rank values. Using the min volume value from the client state.")
		log.Printf("min volume: %f\n", cs.CurrentMinVolume)
		minRank, err := strconv.Atoi(args[0])
		if err != nil {
		        log.Fatal(err)
		}
		maxRank, err := strconv.Atoi(args[1])
		if err != nil {
		        log.Fatal(err)
		}
		cs.UpdateMarketRank(minRank, maxRank)
	case 3:
	        minRank, err := strconv.Atoi(args[0])
		if err != nil {
		        log.Fatal(err)
		}
		maxRank, err := strconv.Atoi(args[1])
		if err != nil {
		        log.Fatal(err)
		}
		minVolume, err := strconv.ParseFloat(args[2], 64)
		if err != nil {
		        log.Fatal(err)
		}
		cs.UpdateMarketRank(minRank, maxRank)
		cs.UpdateVolume(minVolume, cs.CurrentMaxVolume)
	        log.Println("Client state min and max market ranks, and min volume values are updated with the passed arguments.")
		default:
		        log.Println("Invalid amount of arguments! Please Provide: <min_market_rank int> <max_market_rank int> <min_volume float64>...")
	}
}
