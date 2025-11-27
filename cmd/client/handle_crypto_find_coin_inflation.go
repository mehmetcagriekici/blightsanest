package main

import(
        "log"
	"strconv"
	"slices"
	"strings"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

func handleCryptoFindCoinInflation(cs *crypto.CryptoState) {
        log.Println("Starting to find coins with large supply value and relatively low market cap rank with client preferences.")
	log.Println("")
	
        list := crypto.CoinsHighCirculatingSupply(cs.CurrentMinRank, cs.CurrentMinSupply, cs.CurrentIgnoredCoins, cs.CurrentList)
	fields := []string{"MaxSupply", "CirculatingSupply", "MarketCap", "MarketCapRank"}
	crypto.PrintCryptoList(list, cs.CurrentListID, cs.ClientTimeframes, fields)
	log.Println("")
	log.Println("To update the client list with the result: mutate find crypto coin_inflation")
}

func controlFindCoinInflation(cs *crypto.CryptoState, args []string) {
        switch len(args) {
	case 0:
	        log.Println("No arguments are passed. Using the client state values.")
		log.Printf("Current Min Market Cap Rank: %d\n", cs.CurrentMinMarketCap)
		log.Printf("Current Min Supply Value: %f\n", cs.CurrentMinSupply)
		log.Printf("Current Ignored Coins: %s\n", strings.Join(cs.CurrentIgnoredCoins, " "))
	case 1:
	        log.Println("Updating the client min market cap rank preference with the passed argument. Using current client values for min supply and ignored coins preferences.")
		log.Printf("Current Min Supply Value: %f", cs.CurrentMinSupply)
		log.Printf("Current Ignored Coins: %s\n", strings.Join(cs.CurrentIgnoredCoins, " "))
		minRank, err := strconv.Atoi(args[0])
		if err != nil {
		        log.Fatal(err)
		}
		cs.UpdateMarketRank(minRank, cs.CurrentMaxRank)
	case 2:
	        log.Println("Updating the client market cap rank and supply preferences. Using the current client values for ignored coins preferences.")
		log.Printf("Current Ignored Coins: %s\n", strings.Join(cs.CurrentIgnoredCoins, " "))
		minRank, err := strconv.Atoi(args[0])
		if err != nil {
		        log.Fatal(err)
		}
		minSupply, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
		        log.Fatal(err)
		}
		cs.UpdateMarketRank(minRank, cs.CurrentMaxRank)
		cs.UpdateSupply(minSupply, cs.CurrentMaxSupply)
	default:
	        log.Println("Updating the client market cap, supply, and ignored coins preferences. After the min supply argument, all space-separated arguments considered as coin names.")
		minRank, err := strconv.Atoi(args[0])
		if err != nil {
		        log.Fatal(err)
		}
		minSupply, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
		        log.Fatal(err)
		}
		cs.UpdateMarketRank(minRank, cs.CurrentMaxRank)
		cs.UpdateSupply(minSupply, cs.CurrentMaxSupply)
		cs.CurrentIgnoredCoins = args[2:]
	}
}