package crypto

import(
        "errors"
	"slices"
)
	
// function to watch daily price swing
func CalcCoinVolatility(minVolatility, maxVolatility float64, coins []MarketData) []MarketData {
        // slice to store low risk coins
        lowRiskCoins := []MarketData{}

        // iterate over the coins and calc volatility if lower append it to the low risk coins
        for coin := range slices.Values(coins) {
	        if coin.CurrentPrice == 0 {
		        continue
		}
		
	        volatility := (coin.High24H - coin.Low24H) / coin.CurrentPrice
		if volatility <= maxVolatility && volatility >= minVolatility {
		        lowRiskCoins = append(lowRiskCoins, coin)
		}
	}

        // return the low risk coins
	return lowRiskCoins
}

// function to calculate potential growth
func EstimateCoinUpsidePotential(minPotential float64, maxMarketRank int, coins []MarketData) []MarketData {
        // slice to store high potential coins
        highPotentialCoins := []MarketData{}

        // iterate over the coins and get the high potential ones
	for coin := range slices.Values(coins) {
	        if coin.CurrentPrice == 0 {
		        continue
		}
		
                // potential growth score
                pgs := (coin.ATH - coin.CurrentPrice) / coin.CurrentPrice * float64(100)
	        if coin.MarketCapRank <= maxMarketRank && pgs >= minPotential {
		        highPotentialCoins = append(highPotentialCoins, coin)
		}
	}

        // return the high potential coins
	return highPotentialCoins
}

// function to calculate liquidty score
func CalcCoinLiquidity(minLiquidity float64, coins []MarketData) []MarketData {
        // slice to store high liquidity coins
	highLiquidityCoins := []MarketData{}

        // iterate over the coins and get the coins with high liquidity score
	for coin := range slices.Values(coins) {
	        if coin.MarketCap == 0 {
		        continue
		}

                if coin.TotalVolume / coin.MarketCap >= minLiquidity {
		        highLiquidityCoins = append(highLiquidityCoins, coin)
		}
	}

        // return the high liquidity coins
	return highLiquidityCoins
}

// function to see if a coind is real trend or fake pump
func CheckRealTrend(timeframe AvailableTimeframes, coins []MarketData) ([]MarketData, error) {
        // slice to store real trend coins
	realTrendCoins := []MarketData{}

        // only available for 24H timeframe
	if timeframe != PCP_DAY {
	        return realTrendCoins, errors.New("This feature is available only for the 24H timeframe")
	}
        
        // iterate over the coins and find the real trending ones
	for coin := range slices.Values(coins) {
	        if coin.PriceChangePercentage24h > 0 && coin.MarketCapChangePercentage > 0 {
		        realTrendCoins = append(realTrendCoins, coin)
		}
	}

        // return the real trend coins
	return realTrendCoins, nil
}