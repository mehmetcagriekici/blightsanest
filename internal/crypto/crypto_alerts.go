package crypto

import(
        "sort"
        "slices"
)

// functions to get the coins that hit a new high/low of the day
func CoinsNewHigh(oldCoins, newCoins []MarketData) []MarketData {
        // copy the slices
	cloneOld := slices.Clone(oldCoins)
	cloneNew := slices.Clone(newCoins)
	
	// sort the slices by the current prices descending order
	sort.Slice(cloneOld, func(i, j int) bool {
	        return SortCoinNames(cloneOld, i, j)
	})
	sort.Slice(cloneNew, func(i, j int) bool {
	        return SortCoinNames(cloneNew, i, j)
	})
	
	// compare the slices
	filtered := []MarketData{}
	for i, coin := range cloneNew {
	        if coin.Name != cloneOld[i].Name {
		        continue
		}
	        if coin.High24H > cloneOld[i].High24H {
		        filtered = append(filtered, coin)
		}
	}
	
	// return the new slice
	return filtered
}

func CoinsNewLow(oldCoins, newCoins []MarketData) []MarketData {
        // copy the slices
	cloneOld := slices.Clone(oldCoins)
	cloneNew := slices.Clone(newCoins)
	
	// sort the slices by the current prices descending order
	sort.Slice(cloneOld, func(i, j int) bool {
	        return SortCoinNames(cloneOld, i, j)
	})
	sort.Slice(cloneNew, func(i, j int) bool {
	        return SortCoinNames(cloneNew, i, j)
	})
	
	// compare the slices
	filtered := []MarketData{}
	for i, coin := range cloneNew {
	        if coin.Name != cloneOld[i].Name {
		        continue
		}
	        if coin.Low24H < cloneOld[i].Low24H {
		        filtered = append(filtered, coin)
		}
	}
	
	// return the new slice
	return filtered
}

// function to get the coins that have a treshold price spike
func CoinsHighPriceSpike(tresholdRate float64, timeframe AvailableTimeframes, coins []MarketData) []MarketData {
        // copy the coins
	clone := slices.Clone(coins)

        // delete the coins under the treshold rate - filter higher or equal
	clone = slices.DeleteFunc(clone, func(coin MarketData) bool {
	        priceChange := GetPriceChange(coin, timeframe)
		return priceChange < tresholdRate
	})

        // return the copy
	return clone
}

// function to get the coins that reached a percentage of their aths
func CoinsGetCloseAthChange(maxAthChange float64, coins []MarketData) []MarketData {
        // copy the coins
	clone := slices.Clone(coins)

        // delete the coins with high ath changes
	clone = slices.DeleteFunc(clone, func(coin MarketData) bool {
	        return coin.AthChangePercentage > maxAthChange
	})

        // return the copy
	return clone
}

// function to get the coins with high circulating_supply
func CoinsHighCirculatingSupply(alertMarketRank int, alertValue float64, ignoreCoins []string, coins []MarketData) []MarketData {
        // copy the coins
	clone := slices.Clone(coins)

       // delete the ignored coins
       clone = slices.DeleteFunc(clone, func(coin MarketData) bool {
               return coin.MarketCapRank > alertMarketRank || coin.CurrentPrice * coin.CirculatingSupply < alertValue || slices.Contains(ignoreCoins, coin.Name)
       })

       // return the copy
       return clone
}