package routing

import(
        "time"
	
        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
)

type CryptoExchangeBody struct {
        ID        string
	CreatedAt time.Time
        Payload   []crypto.MarketData
}