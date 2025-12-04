package crypto

import (
	"sort"
        "slices"
)

// function to rank coins by any number field
func RankCoins(fieldName string, order AvailableOrders, coins []MarketData) []MarketData {
        // clone the coins slice
	clone := slices.Clone(coins)

        // sort the clone
	sort.Slice(clone, func(i, j int) bool {
	        if order == CRYPTO_DESC {
	                return GetCoinField(fieldName, clone[i]).Float() >  GetCoinField(fieldName, clone[j]).Float()
		}
		if order == CRYPTO_ASC {
		        return GetCoinField(fieldName, clone[i]).Float() < GetCoinField(fieldName, clone[j]).Float()
		}
		return GetCoinField(fieldName, clone[i]).Equal(GetCoinField(fieldName, clone[j]))
	})

        // return the clone
	return clone
}

// function to group coins by market_cap_rank and combines with total_volume to filter out low-liquidity ones
func GroupHighLiquidityCoins(minRank, maxRank int, minVolume float64, coins []MarketData) []MarketData {
        // clone the coins slice
	clone := slices.Clone(coins)

        // delete the coins which do no meet the credentials -inside the range and above the volume-
	clone = slices.DeleteFunc(clone, func(coin MarketData) bool {
	        return coin.MarketCapRank < minRank || coin.MarketCapRank > maxRank || coin.TotalVolume < minVolume
	})

        // return the filtered clone
	return clone
}

// function to sort the coins by max_supply vs. circulating_supply ratio to identify scarce assets and pair with ath_change_percentage to find undervalued gems near their lows
func RankCoinScarcity(circulatingScore, athChangeScore float64, coins []MarketData) []MarketData {
        // clone the coins slice
	clone := slices.Clone(coins)

        // delete the coins with low scarcity and high ath_change_percentage
	clone = slices.DeleteFunc(clone, func(coin MarketData) bool {
	        return coin.CirculatingSupply / coin.MaxSupply < circulatingScore || coin.AthChangePercentage > athChangeScore
	})

        // return the clone
	return clone
} 
