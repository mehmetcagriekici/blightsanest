package crypto

import (
        "io"
	"net/http"
	"encoding/json"
	"log"
)

// market type
type MarketData struct{
      ID                        string  `json:"id"`
      Symbol                    string  `json:"symbol"`
      Name                      string  `json:"name"`
      CurrentPrice              float64 `json:"current_price"`
      MarketCap                 float64 `json:"market_cap"`
      MarketCapRank             int     `json:"market_cap_rank"`
      MarketCapChangePercentage float64 `json:"marke_cap_change_percentage_24h"`
      TotalVolume               float64 `json:"total_volume"`
      High24H                   float64 `json:"high_24h"`
      Low24H                    float64 `json:"low_24h"`
      ATH                       float64 `json:"ath"`
      PriceChangePercentage1h   float64 `json:"price_change_percentage_1h"`
      PriceChangePercentage24h  float64 `json:"price_change_percentage_24h"`
      PriceChangePercentage7h   float64 `json:"price_change_percentage_7d"`
      PriceChangePercentage30h  float64 `json:"price_change_percentage_30d"`
      PriceChangePercentage200h float64 `json:"price_change_percentage_200d"`
      PriceChangePercentage1y   float64 `json:"price_change_percentage_1y"`
      AthChangePercentage       float64 `json:"ath_change_percentage"`
      MaxSupply                 float64 `json:"max_supply"`
      CirculatingSupply         float64 `json:"circulating_supply"`
}

// Available orders
type AvailableOrders string
const (
        CRYPTO_DESC AvailableOrders = "desc"
	CRYPTO_ASC AvailableOrders  = "asc"
)

// Price Change Percentage Available timeframes
type AvailableTimeframes string
const (
        PCP_HOUR        AvailableTimeframes = "1h"
	PCP_DAY         AvailableTimeframes = "24h"
	PCP_WEEK        AvailableTimeframes = "7d"
	PCP_MONTH       AvailableTimeframes = "30d"
	PCP_TWO_HUNDRED AvailableTimeframes = "200d"
	PCP_YEAR        AvailableTimeframes = "1y"
)

// function to fetch top gainers/losers in any timeframe
// currency=usd
// order=market_cap_{ORDER}
// price_change_percentage={TIMEFRAME}
func CryptoFetchMarket(url, key string) ([]MarketData, error) {
        client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
	        log.Println("An error occured while trying to create a new request.")
	        return []MarketData{}, err
	}

        // api header
	req.Header.Add("x-cg-demo-api-key", key)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
	        log.Println("An error occured while trying to make a request to the CoinGecko API.")
	        return []MarketData{}, err
	}

        bytes, err := io.ReadAll(resp.Body)
	if err != nil{
	        log.Println("An error occured while tryin to read the response body.")
	        return []MarketData{}, err
	}
        
        // decode the response body
	var market []MarketData
	if err := json.Unmarshal(bytes, &market); err != nil{
	        log.Println("An error occured while trying to decode the bytes from the response body.")
	        return []MarketData{}, err
	}

        return market, nil
}
