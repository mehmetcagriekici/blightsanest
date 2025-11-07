package crypto

import(
        "slices"
	"strings"
)

// functions to let users filter by tresholds like total_colume, market_cap, price_change_percentage
func FilterCoinVolume(minVolume float64, coins []MarketData) []MarketData {
        // copy coins
	copy := slices.Clone(coins)

        // delete the undesired coins
	copy = slices.DeleteFunc(copy, func(coin MarketData) bool {
	        return coin.TotalVolume < minVolume
	})

        // return the copy
	return copy
}

func FilterCoinCap(minCap float64, coins []MarketData) []MarketData {
        // copy coins
	copy := slices.Clone(coins)

        // delete the undesired coins
	copy = slices.DeleteFunc(copy, func(coin MarketData) bool {
	        return coin.MarketCap < minCap
	})

        // return the copy
	return copy
}

func FilterCoinPriceChange(minChange float64, timeframe AvailableTimeframes, coins []MarketData) []MarketData {
        // copy coins
	copy := slices.Clone(coins)

        // delete the undesired coins
	copy = slices.DeleteFunc(copy, func(coin MarketData) bool {
	        return GetPriceChange(coin, timeframe) < minChange
	})

        // return the copy
	return copy
}

func FindWildSwingCoins(minRate float64, coins []MarketData) []MarketData {
        // copy coins
	copy := slices.Clone(coins)

        // delete the lower swinging coins
	copy = slices.DeleteFunc(copy, func(coin MarketData) bool {
	        return coin.High24H / coin.Low24H < minRate
	})

        // return the copy
	return copy
}

// function for quick lookups by name
func SearchCoin(name string, coins []MarketData) (MarketData, bool) {
        // copy coins
	copy := slices.Clone(coins)

        // sort the copy by name
	slices.SortFunc(copy, func(c1, c2 MarketData) int {
	        return strings.Compare(strings.ToLower(c1.Name), strings.ToLower(c2.Name))
	})

        // get the names
	names := []string{}
	for c := range slices.Values(copy) {
	        names = append(names, strings.ToLower(c.Name))
	}

        // search for the coin
	coin, found := slices.BinarySearch(names, strings.ToLower(name))
	if !found {
	        return MarketData{}, false 
	}
	return coin, true
}

// function to flag high-risk with ath_change_percentage near 0% or low total_volume
func FlagRiskCoins(maxAthChange, minVolume float64, coins []MarketData) []MarketData {
        // copy coins
	copy := slices.Clone(coins)

        // delete the low risk coins
        copy = slices.DeleteFunc(copy, func(coin MarketData) bool {
                return coin.AthChangePercentage > maxAthChange || coin.TotalVolume < minVolume
        }) 

        // return the copy
	return copy
}

// function to flag safe coins
func FlagSafeCoins(minMarketRank int, maxPriceChange float64, timeframe AvailableTimeframes, coins []MarketData) []MarketData {
        // copy coins
	copy := slices.Clone(coins)

        // delete the high risk coins
	copy = slices.DeleteFunc(copy, func(coin MarketData) bool {
	        return coin.MarketCapRank < minMarketRank || GetPriceChange(coin, timeframe) > maxPriceChange
	})

        // return the copy
	return copy
} 
