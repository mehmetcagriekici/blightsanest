package crypto

import(
        "log"
        "fmt"
	"reflect"
	"strings"
	"math"
	"slices"
)

// function to get a fieldname vy its name - price_percentage_change_{TIMEFRAME} multiple possible fields depending on the input
func GetPriceChange(coin MarketData, timeframe AvailableTimeframes) float64 {
        field := fmt.Sprintf("PriceChangePercentage%s", timeframe)
	r := reflect.ValueOf(coin)
	val := reflect.Indirect(r).FieldByName(field)
	if !val.CanFloat() {
	        log.Fatal("An error occured while trying to get the price percentage field.")
	}
	return val.Float()
}


// function to sort coins by their names
func SortCoinNames(coins []MarketData, i, j int) bool {
        return strings.Compare(coins[i].Name, coins[j].Name) < 0
}

// function to get crypto cache key createdAts
func GetCryptoCacheHour(unix int64) float64 {
        d := float64(3600)
        u := float64(unix)

        // get the unix hours and hours with reminder
	hours := math.Floor(u / d)
	fullHours := u / d

        // seconds
	reminder := (fullHours - hours) * d

       // calc hourly unix in seconds and return the hour to be used as a cache key
       return math.Floor(u - reminder) / d
}

// function to create crypto cache key
func CreateCryptoCacheKey(timeframes []string, unix int64) string {
        frames := strings.Join(timeframes, "-")
	createdAt := GetCryptoCacheHour(unix)
	
        return fmt.Sprintf("cryptoFrames_%s__createdAt-%.0f", frames, createdAt)
}

// function to get timeframes array
func GetInputTimeFrames(frames []string) []AvailableTimeframes {
        timeframes := []AvailableTimeframes{}
	for frame := range slices.Values(frames) {
	        switch frame {
		case "1h":
		        timeframes = append(timeframes, PCP_HOUR)
		case "24h":
		        timeframes = append(timeframes, PCP_DAY)
		case "7d":
		        timeframes = append(timeframes, PCP_WEEK)
		case "30d":
		        timeframes = append(timeframes, PCP_MONTH)
		case "200d":
		        timeframes = append(timeframes, PCP_TWO_HUNDRED)
		case "1y":
		        timeframes = append(timeframes, PCP_YEAR)
		default:
		        log.Println("Invalid timeframe! (1h, 24h, 7d, 30d, 200d, 1y)")
			continue
		}
	}
}