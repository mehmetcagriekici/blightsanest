package main

import(
        "log"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

func handleCryptoSwitch(cs  *crypto.CryptoState,  cc  *crypto.CryptoCache, args []string) {
	defer log.Print("> ")
	
	if len(args) != 1 {
	        log.Println("Please provide an ID of an existing list.")
		return
	}
	
	key := args[0]

        cryptoEntry, ok := cc.Get(key)
	if !ok {
	        log.Println("Requested list does not exist in the client cache.")
		log.Println("To make a get request to the server:")
		log.Printf("get crypto %s\n", key)
		return
	}

        log.Println("Updating the current list with the requested one...")
	cs.UpdateCurrentList(key, cryptoEntry.Market)
}
