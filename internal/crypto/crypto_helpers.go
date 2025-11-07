package crypto

import(
        "fmt"
	"reflect"
)

// function to get a fieldname vy its name - price_percentage_change_{TIMEFRAME} multiple possible fields depending on the input
func GetPriceChange(coin MarketData, timeframe AvailableTimeframes) float64 {
        field := fmt.Sprintf("PriceChangePercentage%s", timeframe)
	r := reflect.ValueOf(coin)
	return reflet.Indirect(r).FieldByName(field)
}
