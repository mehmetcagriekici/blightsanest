package main

import(
        "log"
	"strconv"
	"strings"
	"fmt"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

func handleCryptoCalcGrowthPotential(cs *crypto.CryptoState, args []string) {
        defer log.Print("> ")
	
        controlCalcGrowthPotential(cs, args)

        list := crypto.EstimateCoinUpsidePotential(cs.CurrentMinGrowthPotential, cs.CurrentMaxRank, cs.CurrentList)

        baseID := strings.Split(cs.CurrentListID, "_")[0]
        newID := fmt.Sprintf("%s_calc_growth_potential_%s_%s", baseID, cs.CurrentMinGrowthPotential, cs.CurrentMaxRank)
	cs.UpdateCurrentList(newID, list)
	
	fields := []string{"ATH", "AthChangePercentage"}
	crypto.PrintCryptoList(cs.CurrentList, cs.CurrentListID, cs.ClientTimeframes, fields)
	log.Println("")

        return
}

func controlCalcGrowthPotential(cs *crypto.CryptoState, args []string) {
        switch len(args) {
	case 0:
	        log.Println("No arguments are passed. Using the client preferences for min growth potential and max market cap rank values.")
		log.Printf("Min Growth Potential: %f\n", cs.CurrentMinGrowthPotential)
		log.Printf("Max Market Cap Rank: %d\n", cs.CurrentMaxRank)
	case 1:
	        log.Println("One argument is passed. Updating the client min growth potential preference...")
		log.Printf("Max Market Cap Rank: %d\n", cs.CurrentMaxRank)
		minPotential, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
		        log.Fatal(err)
		}
		cs.UpdateGrowthPotential(minPotential, cs.CurrentMaxGrowthPotential)
	case 2:
	        log.Println("Updating the client min growth potential and max market cap rank preferences...")
		minPotential, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
		        log.Fatal(err)
		}
		maxRank, err := strconv.Atoi(args[1])
		if err != nil {
		        log.Fatal(err)
		}
		cs.UpdateGrowthPotential(minPotential, cs.CurrentMaxGrowthPotential)
		cs.UpdateMarketRank(cs.CurrentMinRank, maxRank)
	default:
	        log.Println("Invalid use of command: calc crypto growth_potential <min_potential float64> <max_market_cap_rank int>")
	}
}