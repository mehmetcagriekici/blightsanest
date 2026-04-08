package cmd

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/mehmetcagriekici/blightsanest/internal/crypto"
	"github.com/mehmetcagriekici/blightsanest/internal/pubsub"
	"github.com/mehmetcagriekici/blightsanest/internal/routing"
)

// Available API queries
const (
	GECKO_IDS                     string = "ids"
	GECKO_NAMES                   string = "names"
	GECKO_SYMBOLS                 string = "symbols"
	GECKO_INCLUDE_TOKENS          string = "include_tokens"
	GECKO_CATEGORY                string = "category"
	GECKO_ORDER                   string = "order"
	GECKO_PER_PAGE                string = "per_page"
	GECKO_PAGE                    string = "page"
	GECKO_SPARKLINE               string = "sparkline"
	GECKO_PRICE_CHANGE_PERCENTAGE string = "price_change_percentage"
	GECKO_PERCISION               string = "percision"
)

var queryParameters = []string{GECKO_IDS,
	GECKO_NAMES,
	GECKO_SYMBOLS,
	GECKO_INCLUDE_TOKENS,
	GECKO_CATEGORY,
	GECKO_ORDER,
	GECKO_PER_PAGE,
	GECKO_PAGE,
	GECKO_SPARKLINE,
	GECKO_PRICE_CHANGE_PERCENTAGE,
	GECKO_PERCISION}

var fetchCryptoCmd = &cobra.Command{
	Use:   "crypto [args...]",
	Short: "Fetch crypto data from coingecko api.",
	Run:   handleCryptoFetch,
}

func handleCryptoFetch(cmd *cobra.Command, args []string) {
	// create the request URL
	queries := createCryptoFetchURLQueries(args)
	url := "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd"
	if queries != "" {
		url = fmt.Sprintf("%s&%s", url, queries)
	}

	// control cache - data is cached for each hour based on the queries
	cacheKey := createCryptoCacheKey(time.Now().Unix(), queries)
	if _, ok := CryptoCache.Get(cacheKey); !ok {
		log.Println("Requested crypto list does not exists in the server cache, making a new request to the API.")
		// make the API request
		cryptoList, err := crypto.CryptoFetchMarket(url, CryptoAPIKey)
		if err != nil {
			log.Fatal(err)
		}

		// add list to the cache
		CryptoCache.Add(cacheKey, cryptoList)
	}

	// get the data from the cache
	cacheEntry, ok := CryptoCache.Get(cacheKey)
	if !ok {
		log.Fatal("Requested crypto list could not be fetched.")
	}

	log.Printf("Publishing the requested crypto list with the id: %s\n", cacheKey)
	delivery := routing.CryptoExchangeBody{
		ID:        cacheKey,
		CreatedAt: cacheEntry.CreatedAt,
		Payload:   cacheEntry.Market,
	}

	if err := pubsub.PublishCrypto(Ctx, Conn, delivery); err != nil {
		log.Fatal(err)
	}
}

// server caching
func createCryptoCacheKey(unix int64, queries string) string {
	cacheHour := crypto.GetCryptoCacheHour(unix)
	return fmt.Sprintf("%s_%s", cacheHour, queries)
}

func createCryptoFetchURLQueries(args []string) string {
	queries := []string{}

	if len(args) == 0 {
		return ""
	}

	for i, q := range args {
		if q == "-" {
			continue
		}
		query := fmt.Sprintf("%s=%s", queryParameters[i], q)
		queries = append(queries, query)
	}

	return strings.Join(queries, "&")
}
