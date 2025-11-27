package main

import(
        "log"
	"strconv"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

func handleCryptoCalcVolatility(cs *crypto.CryptoState) {
        log.Println("Calculating the daily volatility ranges of the coins...")
	log.Println("")

        list := crypto.CalcCoinVolatility(cs.CurrentMinVolatility, cs.CurrentMaxVolatility, cs.CurrentList)
	fields := []string{"High24H", "Low24H"}
	crypto.PrintCryptoList(list, cs.CurrentListID, cs.ClientTimeframes, fields)
	log.Println("")
        log.Println("To update the current client list with the result: mutate calc crypto volatility")
	log.Println("")
}

func controlCalcVolatility(cs *crypto.CryptoState, args []string) {
        switch len(args) {
	case 0:
	        log.Println("")
	}
}