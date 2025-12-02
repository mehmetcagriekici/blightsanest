package main

import(
        "log"
	
	"github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

func handleCryptoList(cs *crypto.CryptoState, cc *crypto.CryptoCache) {
       log.Printf("Current Crypto List ID: %s\n", cs.CurrentListID)
       
       if len(cc.Market) == 0 {
               log.Println("Client Cache is empty.")
	       return
       }

       for k := range cc.Market {
               log.Printf("List: %s\n", k)
       }
}