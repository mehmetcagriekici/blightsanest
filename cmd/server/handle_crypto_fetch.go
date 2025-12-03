package main

import(
        "log"
	"strings"
	"fmt"
	"time"
	"context"
	
	amqp "github.com/rabbitmq/amqp091-go"

        "github.com/mehmetcagriekici/blightsanest/internal/pubsub"
	"github.com/mehmetcagriekici/blightsanest/internal/crypto"
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

func handleCryptoFetch(ctx context.Context, conn *amqp.Connection, cc *crypto.CryptoCache, apiKey string, args []string) {
        // create the request URL
        queries := createCryptoFetchURLQueries(args)
	url := "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd"
	if queries != "" {
	        url = fmt.Sprintf("%s&%s", url, queries)
	}

        // control cache - data is cached for each hour based on the queries
	cacheKey := createCryptoCacheKey(time.Now().Unix(), queries)
	_, ok := cc.Get(cacheKey)
	if !ok {
	        log.Println("Requested crypto list does not exists in the server cache, making a new request to the API.")
	        // make the API request
		cryptoList, err := crypto.CryptoFetchMarket(url, apiKey)
		if err != nil {
		        log.Fatal(err)
		}
			
		// add list to the cache
		cc.Add(cacheKey, cryptoList)
	}

        // get the data from the cache
        cacheEntry, ok := cc.Get(cacheKey)
	if !ok {
	        log.Fatal("Requested crypto list could not be fetched.")
	}
	
        log.Printf("Publishing the requested crypto list with the id: %s\n", cacheKey)
	delivery := routing.CryptoExchangeBody{
	        ID:        cacheKey,
		CreatedAt: cacheEntry.CreatedAt,
		Payload:   cacheEntry.Market,
	}
	
	if err := pubsub.PublishCrypto(ctx, conn, delivery); err != nil {
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

func printCryptoFetchHelp() {
        log.Println("All the text for Available Queries below is copied from the documentation link below.")
        log.Println("Documentation Link: https://docs.coingecko.com/reference/coins-markets")
	log.Println("vs_currency is usd and locale is not specified - default: en -")
	log.Println("")
	log.Println("")
	log.Println("Available Queries:")

        printQueryExplanation(GECKO_IDS,
	                      "bitcoin",
			      "coins' IDs",
			      "refers to /coins/list",
			      "comma-separated if querying more than 1 coin.")
	printQueryExplanation(GECKO_NAMES,
	                      "Bitcoin",
			      "coins' names",
			      "",
			      "comma-separated if querying more than 1 coin.")
	printQueryExplanation(GECKO_SYMBOLS,
	                      "btc",
			      "coins' symbols",
			      "",
			      "comma-separated if querying more than 1 coin.")
	printQueryExplanation(GECKO_INCLUDE_TOKENS,
	                      "top",
			      "when specified top returns top-ranked tokens (by market cap or volume)",
			      "for symbols lookups, specify all to include all matching tokens",
			      "Available options: top, all")
	printQueryExplanation(GECKO_CATEGORY,
	                      "layer-1",
			      "filter based on coins' category",
			      "refers to /coins/categories/list",
			      "Example: layer-1")
	printQueryExplanation(GECKO_ORDER,
	                      "market_cap_desc",
			      "sort result by field",
			      "",
			      "Available options: market_cap_asc, market_cap_desc, volume_asc, volume_desc, id_asc, id_desc")
	printQueryExplanation(GECKO_PER_PAGE,
	                      "100",
			      "total results per page",
			      "nuber",
			      "Valid values: 1...250")
	printQueryExplanation(GECKO_PAGE,
	                      "1",
			      "page through results",
			      "",
			      "number")
	printQueryExplanation(GECKO_SPARKLINE,
	                      "false",
			      "include sparkline 7 days data",
			      "",
			      "boolean")
	printQueryExplanation(GECKO_PRICE_CHANGE_PERCENTAGE,
	                      "1h",
			      "include price change percentage timeframe",
			      "Valid values: 1h, 24h, 7d, 14d, 30d, 200d, 1y",
			      "comma-separated if query more than 1 timeframe")
	printQueryExplanation(GECKO_PERCISION,
	                      "",
			      "decimal place for currency price value",
			      "",
			      "Available options: full, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18")
}

func printQueryExplanation(query, defaultValue, definition, refersTo, details string) {
        log.Println("")
        log.Printf("%s:\n", query)
	log.Printf("    default:     %s\n", defaultValue)
	log.Printf("    definition:  %s\n", definition)
	log.Printf("    referrings:  %s\n", refersTo)
	log.Printf("    details:     %s\n", details)
	log.Println("")
}