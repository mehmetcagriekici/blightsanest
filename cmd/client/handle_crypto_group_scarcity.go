package main

import(
        "log"
	"strconv"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

func handleCryptoGroupScarcity(cs *crypto.CryptoState, args []string) {
        controlScarcityArguments(cs, args)
	
	list := crypto.RankCoinScarcity(cs.CurrentMinCirculatingSupply, cs.CurrentMaxATHChangePercentage, cs.CurrentList)
	log.Println("Successfully identified scarce coins with limited supply and large darwdowns from ATH.")
	log.Println("")
	fields := []string{"ATH", "AthChangePercentage", "CirulatingSupply", "MaxSupply"}
	crypto.PrintCryptoList(list, cs.CurrentListID, cs.ClientTimeframes, fields)
	log.Println("")
	log.Println("")
	log.Println("To update the client state list with the result: mutate group crypto scarcity")
}

// min circulatin supply
// max ath change percentage
func controlScarcityArguments(cs *crypto.CryptoState, args []string) {
        switch len(args) {
	case 0:
	        log.Println("No arguments are passed. Using the values from the current client state.")
		log.Printf("min circulating supply: %f\n", cs.CurrentMinCirculatingSupply)
		log.Printf("max ath change percentage: %f\n", cs.CurrentMaxATHChangePercentage)
	case 1:
	        log.Println("One argument is passed. Using the max ath change percentage value from the current client state.")
                log.Printf("max ath change percentage: %f\n", cs.CurrentMaxATHChangePercentage)
		minSupply, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
		        log.Fatal(err)
		}
		log.Println("Updating the min circulating supply value...")
		cs.UpdateCirculatingSupply(minSupply, cs.CurrentMaxCirculatingSupply)
	case 2:
	        log.Println("Updating the client state with the new values passed as arguments.")
		minSupply, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
		        log.Fatal(err)
		}
		maxATHChange, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
		        log.Fatal(err)
		}
		cs.UpdateCirculatingSupply(minSupply, cs.CurrentMaxCirculatingSupply)
		cs.UpdateAthChangePercentage(cs.CurrentMinATHChangePercentage, maxATHChange)
	default:
	        log.Println("Invalid amount of arguments!")
		log.Println("group crypto scarcity <min_circulating_supply float64> <max_ath_change_percentage float64>")
	}
}