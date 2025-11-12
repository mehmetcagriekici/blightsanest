package main

import(
	"log"
	"os"
	"time"
	"strconv"

        "github.com/joho/godotenv"
	"github.com/mehmetcagriekici/internal/crypto"
	"github.com/mehmetcagriekici/internal/serverlogic"
)

func main() {
        // load environment variables
	if err := godotenv.Load(); err != nil {
	        log.Fatal(err)
	}

        // get the crypto api key
	cryptoAPIKey := os.Getenv("COIN_GECKO_KEY")
	interval := os.Getenv("SERVER_INTERVAL")

        // print server introduction
	serverlogic.PrintServerIntroduction()

        // create the server crypto cache
	cryptoCache := crypto.CreateCryptoCache(interval * time.Hour)
	
	// before fetching a new crypto list check if one with the same frame exists in the cache

        // create the server crypto state
}