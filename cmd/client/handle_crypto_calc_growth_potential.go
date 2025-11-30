package main

import(
        "log"
	"strconv"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

func handleCryptoCalcGrowthPotential(cs *crypto.CryptoState, args []string) {
        controlCalcGrowthPotential(cs, args)
	
        log.Println("Calculating the estimate growth potential from current price to ATH withing rank constraints...")
	log.Println("growth_potential = (ATH - current_price) / current_price * 100")
	log.Println("")

        list := crypto.EstimateCoinUpsidePotential(cs.CurrentMinGrowthPotential, cs.CurrentMaxRank, cs.CurrentList)
	fields := []string{"ATH", "AthChangePercentage"}
	crypto.PrintCryptoList(list, cs.CurrentListID, cs.ClientTimeframes, fields)
	log.Println("")
	log.Println("To update the current client list with the result: mutate calc crypto growth_potential")
}

func controlCalcGrowthPotential(cs *crypto.CryptoState, args []string) {
        switch len(args) {
	case 0:
	        log.Println("No arguments are passed. Using the client preferences for min growth potential and max market cap rank values.")
		log.Printf("Min Growth Potential: %f\n", cs.CurrentMinGrowthPotential)
		log.Printf("Max Market Cap Rank: %f\n", cs.CurrentMaxRank)
	case 1:
	        log.Println("One argument is passed. Updating the client min growth potential preference...")
		log.Printf("Max Market Cap Rank: %f\n", cs.CurrentMaxRank)
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