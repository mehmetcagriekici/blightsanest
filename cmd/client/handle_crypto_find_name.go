package main

import (
        "log"
	
        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

func handleCryptoFindName(cs *crypto.CryptoState, name string) {
        if name == "" {
	        log.Fatal("Invalid usage: find crypto name <coin_name string>")
	}
	coin, ok := crypto.SearchCoin(name, cs.CurrentList)
	if !ok {
	        log.Fatal("Couldn't find the coin.")
	}

        fields := []string{"ID", "MarketCap", "MarketCapChangePercentage", "TotalVolume", "ATH", "AthChangePercentage", "MaxSupply", "CirculatingSupply"}
	crypto.PrintCryptoList([]crypto.MarketData{coin}, cs.CurrentListID, cs.ClientTimeframes, fields)
}