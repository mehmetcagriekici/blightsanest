package crypto

import (
        "io"
        "fmt"
	"net/http"
	"encoding/json"
	"strings"
	"slices"
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
      PriceChangePercentage1H   float64 `json:"price_change_percentage_1h"`
      PriceChangePercentage24H  float64 `json:"price_change_percentage_24h"`
      PriceChangePercentage7D   float64 `json:"price_change_percentage_7d"`
      PriceChangePercentage30D  float64 `json:"price_change_percentage_30d"`
      PriceChangePercentage200D float64 `json:"price_change_percentage_200d"`
      PriceChangePercentage1Y   float64 `json:"price_change_percentage_1y"`
      AthChangePercentage       float64 `json:"ath_change_percentage"`
      MaxSupply                 float64 `json:"max_supply"`
      CirculatingSupply         float64 `json:"circulating_supply"`
}

// Available orders
type AvailableOrders string
const (
        MARKET_CAP_DESC AvailableOrders = "market_cap_desc"
	MARKET_CAP_ASC AvailableOrders  = "market_cap_asc"
)

// Price Change Percentage Available timeframes
type AvailableTimeframes string
const (
        PCP_HOUR        AvailableTimeframes = "1H"
	PCP_DAY         AvailableTimeframes = "24H"
	PCP_WEEK        AvailableTimeframes = "7D"
	PCP_MONTH       AvailableTimeframes = "30D"
	PCP_TWO_HUNDRED AvailableTimeframes = "200D"
	PCP_YEAR        AvailableTimeframes = "1Y"
)

// function to fetch top gainers/losers in any timeframe
// currency=usd
// order=market_cap_{ORDER}
// price_change_percentage={TIMEFRAME}
func CryptoFetchMarket(order AvailableOrders, timeframes []AvailableTimeframes, key string) ([]MarketData, error) {
        client := &http.Client{}
	
	// convert timeframes to type string
        urlFrames := []string{}
	for timeframe := range slices.Values(timeframes) {
	        urlFrames = append(urlFrames, string(timeframe))
	}
	
        // api url
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=%s&price_change_percentage=%s", order, strings.Join(urlFrames, ","))
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
	        return []MarketData{}, err
	}

        // api header
	req.Header.Add("x-cg-pro-api-key", key)
	resp, err := client.Do(req)
	if err != nil {
	        return []MarketData{}, err
	}

        bytes, err := io.ReadAll(resp.Body)
	if err != nil{
	        return []MarketData{}, err
	}
        resp.Body.Close()

        // decode the response body
	var market []MarketData
	if err := json.Unmarshal(bytes, &market); err != nil{
	        return []MarketData{}, err
	}

        return market, nil
}
