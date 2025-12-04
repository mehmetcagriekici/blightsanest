package main

import (
        "log"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

func handleCryptoNewHighPrice(cs *crypto.CryptoState, cc *crypto.CryptoCache) {
        list := findNewPrice(cs, cc, crypto.CoinsNewHigh)
	if len(list) == 0 {
	        return
	}
	newID := "find_new_high_price"
        fields := []string{"High24H"}
	commonCryptoHandler(cs, list, fields, newID)
}

func handleCryptoNewLowPrice(cs *crypto.CryptoState, cc *crypto.CryptoCache) {
        list := findNewPrice(cs, cc, crypto.CoinsNewLow)
	if len(list) == 0 {
	        return
	}
	newID := "find_new_low_price"
        fields := []string{"Low24H"}
	commonCryptoHandler(cs, list, fields, newID)
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