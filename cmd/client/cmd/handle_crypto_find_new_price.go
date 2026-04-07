package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

var searchCryptoNewHighPriceCmd = &cobra.Command{
	Use:   "crypto new_high_price",
	Short: "Find the coins with new high price",
	Run:   handleCryptoNewHighPrice,
}

func handleCryptoNewHighPrice(cmd *cobra.Command, args []string) {
	list := findNewPrice(CryptoState, CrpyptoCache, crypto.CoinsNewHigh)
	if len(list) == 0 {
		return
	}
	newID := "find_new_high_price"
	fields := []string{"High24H"}
	commonCryptoHandler(CryptoState, list, fields, newID)
}

var searchCryptoNewLowPriceCmd = &cobra.Command{
	Use:   "crypto new_low_price",
	Short: "Find the coins with new low prices",
	Run:   handleCryptoNewLowPrice,
}

func handleCryptoNewLowPrice(cmd *cobra.Command, args []string) {
	list := findNewPrice(CryptoState, CryptoCache, crypto.CoinsNewLow)
	if len(list) == 0 {
		return
	}
	newID := "find_new_low_price"
	fields := []string{"Low24H"}
	commonCryptoHandler(CryptoState, list, fields, newID)
}

func findNewPrice(cs *crypto.CryptoState, cc *crypto.CryptoCache, foo func(oldCoins, newCoins []crypto.MarketData) []crypto.MarketData) []crypto.MarketData {
	if len(cc.Market) == 0 {
		log.Println("There are no lists in the cache to compare!")
		return []crypto.MarketData{}
	}

	// cache holds lists hourly which are fetched within last 24 hours
	// if the current new high/low is above/below all of the new highs/lows in the cache
	compared := cs.CurrentList
	for k := range cc.Market {
		log.Printf("---- Comparing list %s with current list %s\n", k, cs.CurrentListID)
		cryptoEntry, ok := cc.Get(k)
		if !ok {
			continue
		}
		compared = foo(cryptoEntry.Market, compared)
	}
	return compared
}
