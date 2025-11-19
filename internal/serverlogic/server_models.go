package serverlogic

import(
        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

// server crypto payload
type ServerCryptoPayload struct {
        Key        string
	Coins      []crypto.MarketData
	Strategies []string
}