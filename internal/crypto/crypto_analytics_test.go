package crypto

import(
        "testing"
	"reflect"
)

// test calc coin volatility
//current price
// high_24h, low_4h
func TestCalcCoinVolatility(t *testing.T) {
        maxRisk := (3.12345 - 1.12345) / 2.12345
	coin1 := buildCalcCoinVolatility(2.12345, 3.12345, 1.12345)
	coin2 := buildCalcCoinVolatility(0, 0, 0)
	coin3 := buildCalcCoinVolatility(0.12345, 9.12345, 1.12345)

        expected := []MarketData{coin1}
	result := CalcCoinVolatility(maxRisk, []MarketData{coin1, coin2, coin3})

        if !reflect.DeepEqual(expected, result) {
	        t.Errorf("expected: %v, got: %v", expected, result)
	}
}

// test estimate coin upside potential
// current price
// market cap rank
// ath
func TestEstimateCoinUpsidePotential(t *testing.T) {
        minPotential := (5.12345 - 3.12345) / 3.12345 * 100
	maxRank := 2
	coin1 := buildEstimateCoinUpsidePotential(3.12345, 5.12345, 1)
	coin2 := buildEstimateCoinUpsidePotential(3.12345, 4.12345, 2)
	coin3 := buildEstimateCoinUpsidePotential(7.12345, 1.12345, 3)

        expected := []MarketData{coin1}
	result := EstimateCoinUpsidePotential(minPotential, maxRank, []MarketData{coin1, coin2, coin3})

        if !reflect.DeepEqual(expected, result) {
	        t.Errorf("expected: %v, got: %v", expected, result)
	}
}

// test calc coin liquidity
// market cap
// total volume
func TestCalcCoinLiquidity(t *testing.T) {
        // total_volume / market_cap
        minLiquidity := 3.0
	coin1 := MarketData{
	        TotalVolume: 3,
		MarketCap: 1,
	}
	coin2 := MarketData{
	        TotalVolume: 2.12345,
		MarketCap: 1.12345,
	}
	coin3 := MarketData{
	        TotalVolume: 4.12345,
		MarketCap: 0,
	}
	
        expected := []MarketData{coin1}
	result := CalcCoinLiquidity(minLiquidity, []MarketData{coin1, coin2, coin3})

        if !reflect.DeepEqual(expected, result) {
	        t.Errorf("expected: %v, got: %v", expected, result)
	}
}

// test check real trend
// price change percentage
// market cap percentage
func TestCheckRealTrend(t *testing.T) {
        coin1 := buildCheckRealTrend(2.12345, 1.12345)
	coin2 := buildCheckRealTrend(1.12345, 0.0)
	coin3 := buildCheckRealTrend(0.0, 1.12345)

        expected := []MarketData{coin1}
	result, err := CheckRealTrend(PCP_DAY, []MarketData{coin1, coin2, coin3})

        if !reflect.DeepEqual(expected, result) || err != nil {
	        t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func buildCalcCoinVolatility(currentPrice, high24, low24 float64) MarketData {
        return MarketData{
	        CurrentPrice: currentPrice,
		High24H: high24,
		Low24H: low24,
	}
}

func buildEstimateCoinUpsidePotential(currentPrice, ath float64, marketRank int) MarketData {
        return MarketData{
	        CurrentPrice: currentPrice,
		ATH: ath,
		MarketCapRank: marketRank,
	}
}

func buildCheckRealTrend(priceChange, marketCapChange float64) MarketData {
        return MarketData{
	        PriceChangePercentage24h: priceChange,
		MarketCapChangePercentage: marketCapChange,
	}
}
