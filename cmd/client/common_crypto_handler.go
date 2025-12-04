package main

import(
	"fmt"
	"strings"
	
        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

func commonCryptoHandler(cs         *crypto.CryptoState,
                         list       []crypto.MarketData,
			 fields     []string,
			 id         string) {
	baseID := strings.Split(cs.CurrentListID, "_")[0]
	newID := fmt.Sprintf("%s_%s", baseID, id)
	
	cs.UpdateCurrentList(newID, list)
	
	crypto.PrintCryptoList(cs.CurrentList, cs.CurrentListID, cs.ClientTimeframes, fields)
}