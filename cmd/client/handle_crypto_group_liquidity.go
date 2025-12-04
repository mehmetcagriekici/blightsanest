package main

import(
        "log"
	"strconv"
	"fmt"
	"strings"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

func handleCryptoGroupLiquidity(cs *crypto.CryptoState, args []string) {
        defer log.Print("> ")
	
        controlLiquidityArguments(cs, args)
	
        highCoins := crypto.GroupHighLiquidityCoins(cs.CurrentMinRank, cs.CurrentMaxRank, cs.CurrentMinVolume, cs.CurrentList)

        baseID := strings.Split(cs.CurrentListID, "_")[0]
	newID := fmt.Sprintf("%s_group_liquidity_%s_%s_%s", baseID, cs.CurrentMinRank, cs.CurrentMaxRank, cs.CurrentMinVolume)
	cs.UpdateCurrentList(newID, highCoins)
	
	fields := []string{"MarketCapRank", "TotalVolume"}
	crypto.PrintCryptoList(cs.CurrentList, cs.CurrentListID, cs.ClientTimeframes, fields)
        log.Println("")
	
	return
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
