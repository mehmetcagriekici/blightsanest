package main

import(
        "github.com/mehmetcagriekici/blightsanest/internal/routing"
)

// function to manage async crypto data
func asyncCryptoSubscriptionHandler(handler func(routing.CryptoExchangeBody)) func(routing.CryptoExchangeBody) {
        return func(delivery routing.CryptoExchangeBody) {
	        handler(delivery)
	}
}

