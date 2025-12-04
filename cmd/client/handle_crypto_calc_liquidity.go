package main

import(
        "log"
	"strconv"
	"strings"
	"fmt"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

func handleCryptoCalcLiquidity(cs *crypto.CryptoState, args []string) {
        defer log.Print("> ")
	
        controlCalcLiquidity(cs, args)

        list := crypto.CalcCoinLiquidity(cs.CurrentMinLiquidity, cs.CurrentList)

        baseID := strings.Split(cs.CurrentListID, "_")[0]
	newID := fmt.Sprintf("%s_calc_liquidity_%s", baseID, cs.CurrentMinLiquidity)
	cs.UpdateCurrentList(newID, list)
	
	fields := []string{"TotalVolume", "MarketCap", "MarketCapRank"}
	crypto.PrintCryptoList(cs.CurrentList, cs.CurrentListID, cs.ClientTimeframes, fields)
	log.Println("")

        return
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