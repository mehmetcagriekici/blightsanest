package main

import(
        "log"
	"strconv"
	"strings"
	"fmt"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

func handleCryptoFindCoinInflation(cs *crypto.CryptoState, args []string) {
        defer log.Print("> ")
	
        controlFindCoinInflation(cs, args)
	
        list := crypto.CoinsHighCirculatingSupply(cs.CurrentMinRank, cs.CurrentMinSupply, cs.CurrentIgnoredCoins, cs.CurrentList)

        baseID := strings.Split(cs.CurrentListID, "_")[0]
	newID := fmt.Sprintf("%s_find_coin_inflation_%s_%s_%s", baseID, cs.CurrentMinRank, cs.CurrentMinSupply, cs.CurrentIgnoredCoins)
	cs.UpdateCurrentList(newID, list)
	
	fields := []string{"MaxSupply", "CirculatingSupply", "MarketCap", "MarketCapRank"}
	crypto.PrintCryptoList(cs.CurrentList, cs.CurrentListID, cs.ClientTimeframes, fields)
	log.Println("")

        return
}

func controlFindCoinInflation(cs *crypto.CryptoState, args []string) {
        switch len(args) {
	case 0:
	        log.Println("No arguments are passed. Using the client state values.")
		log.Printf("Current Min Market Cap Rank: %f\n", cs.CurrentMinMarketCap)
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