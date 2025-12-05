package main

import(
        "log"
	"strconv"
	"strings"
	"fmt"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

func handleCryptoFindCoinInflation(cs *crypto.CryptoState, args []string) {
        controlFindCoinInflation(cs, args)
        list := crypto.CoinsHighCirculatingSupply(cs.CurrentMaxRank, cs.CurrentMinSupply, cs.CurrentIgnoredCoins, cs.CurrentList)
	newID := fmt.Sprintf("find_coin_inflation_%d_%f_%s", cs.CurrentMaxRank, cs.CurrentMinSupply, strings.Join(cs.CurrentIgnoredCoins, "_"))
	fields := []string{"MaxSupply", "CirculatingSupply", "MarketCap", "MarketCapRank"}
	commonCryptoHandler(cs, list, fields, newID)
}

func controlFindCoinInflation(cs *crypto.CryptoState, args []string) {
        switch len(args) {
	case 0:
	        log.Println("No arguments are passed. Using the client state values.")
		log.Printf("Current Max Market Cap Rank: %d\n", cs.CurrentMaxRank)
		log.Printf("Current Min Supply Value: %f\n", cs.CurrentMinSupply)
		log.Printf("Current Ignored Coins: %s\n", strings.Join(cs.CurrentIgnoredCoins, " "))
	case 1:
	        log.Println("Updating the client max market cap rank preference with the passed argument. Using current client values for min supply and ignored coins preferences.")
		log.Printf("Current Min Supply Value: %f", cs.CurrentMinSupply)
		log.Printf("Current Ignored Coins: %s\n", strings.Join(cs.CurrentIgnoredCoins, " "))
		maxRank, err := strconv.Atoi(args[0])
		if err != nil {
		        log.Fatal(err)
		}
		cs.UpdateMarketRank(cs.CurrentMinRank, maxRank)
	case 2:
	        log.Println("Updating the client max market cap rank and supply preferences. Using the current client values for ignored coins preferences.")
		log.Printf("Current Ignored Coins: %s\n", strings.Join(cs.CurrentIgnoredCoins, " "))
		maxRank, err := strconv.Atoi(args[0])
		if err != nil {
		        log.Fatal(err)
		}
		minSupply, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
		        log.Fatal(err)
		}
		cs.UpdateMarketRank(cs.CurrentMinRank, maxRank)
		cs.UpdateSupply(minSupply)
	default:
	        log.Println("Updating the client max market cap, supply, and ignored coins preferences. After the min supply argument, all space-separated arguments considered as coin names.")
		maxRank, err := strconv.Atoi(args[0])
		if err != nil {
		        log.Fatal(err)
		}
		minSupply, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
		        log.Fatal(err)
		}
		cs.UpdateMarketRank(cs.CurrentMinRank, maxRank)
		cs.UpdateSupply(minSupply)
		cs.CurrentIgnoredCoins = args[2:]
	}
}