package main

import(
        "log"
	"strconv"
	"fmt"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

func handleCryptoCalcLiquidity(cs *crypto.CryptoState, args []string) {
        controlCalcLiquidity(cs, args)
        list := crypto.CalcCoinLiquidity(cs.CurrentMinLiquidity, cs.CurrentList)
	newID := fmt.Sprintf("calc_liquidity_%f", cs.CurrentMinLiquidity)
	fields := []string{"TotalVolume", "MarketCap", "MarketCapRank"}
	commonCryptoHandler(cs, list, fields, newID)
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
		cs.UpdateCurrentLiquidity(minLiquidity, cs.CurrentMaxLiquidity)
	default:
	        log.Println("Invalid use of command: calc crypto liquidity <min_liquidity float64>")
	}
}