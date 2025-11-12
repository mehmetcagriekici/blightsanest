package crypto

import(
        "slices"
	"strings"
)

// functions to let users filter by tresholds like total_colume, market_cap, price_change_percentage
func FilterCoinVolume(minVolume float64, coins []MarketData) []MarketData {
        // clone coins
	clone := slices.Clone(coins)

        // delete the undesired coins
	clone = slices.DeleteFunc(clone, func(coin MarketData) bool {
	        return coin.TotalVolume < minVolume
	})

        // return the clone
	return clone
}

func FilterCoinCap(minCap float64, coins []MarketData) []MarketData {
        // clone coins
	clone := slices.Clone(coins)

        // delete the undesired coins
	clone = slices.DeleteFunc(clone, func(coin MarketData) bool {
	        return coin.MarketCap < minCap
	})

        // return the clone
	return clone
}

func FilterCoinPriceChange(minChange float64, timeframe AvailableTimeframes, coins []MarketData) []MarketData {
        // clone coins
	clone := slices.Clone(coins)

        // delete the undesired coins
	clone = slices.DeleteFunc(clone, func(coin MarketData) bool {
	        return GetPriceChange(coin, timeframe) < minChange
	})

        // return the clone
	return clone
}

func FindWildSwingCoins(minRate float64, coins []MarketData) []MarketData {
        // clone coins
	clone := slices.Clone(coins)

        // delete the lower swinging coins
	clone = slices.DeleteFunc(clone, func(coin MarketData) bool {
	        return coin.High24H / coin.Low24H < minRate
	})

        // return the clone
	return clone
}

// function for quick lookups by name
func SearchCoin(name string, coins []MarketData) (MarketData, bool) {
        // clone coins
	clone := slices.Clone(coins)

        // sort the clone by name
	slices.SortFunc(clone, func(c1, c2 MarketData) int {
	        return strings.Compare(strings.ToLower(c1.Name), strings.ToLower(c2.Name))
	})

        // get the names
	names := []string{}
	for c := range slices.Values(clone) {
	        names = append(names, strings.ToLower(c.Name))
	}

        // search for the coin
	i, found := slices.BinarySearch(names, strings.ToLower(name))
	if !found {
	        return MarketData{}, false 
	}
	return clone[i], true
}

// function to flag high-risk with ath_change_percentage near 0% or low total_volume
func FlagRiskCoins(maxAthChange, minVolume float64, coins []MarketData) []MarketData {
        // clone coins
	clone := slices.Clone(coins)

        // delete the low risk coins
        clone = slices.DeleteFunc(clone, func(coin MarketData) bool {
                return coin.AthChangePercentage > maxAthChange || coin.TotalVolume < minVolume
        }) 

        // return the clone
	return clone
}

// function to flag safe coins
func FlagSafeCoins(maxMarketRank int, maxPriceChange float64, timeframe AvailableTimeframes, coins []MarketData) []MarketData {
        // clone coins
	clone := slices.Clone(coins)

        // delete the high risk coins
	clone = slices.DeleteFunc(clone, func(coin MarketData) bool {
	        return coin.MarketCapRank > maxMarketRank || GetPriceChange(coin, timeframe) > maxPriceChange
	})

        // return the clone
	return clone
} 
