package crypto

import(
        "testing"
	"reflect"
)

// test rank coins
// price change percentage
func TestRankCoinsASC(t *testing.T) {
        order := MARKET_CAP_ASC
	coin1 := MarketData{
	        PriceChangePercentage24H: 1.12345,
	}
	coin2 := MarketData{
	        PriceChangePercentage24H: 2.12345,
	}
	coin3 := MarketData{
	        PriceChangePercentage24H: 3.12345,
	}

        expected := []MarketData{coin1, coin2, coin3}
	result := RankCoins(PCP_DAY, order, []MarketData{coin1, coin2, coin3})

        if !reflect.DeepEqual(expected, result) {
	        t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestRankCoinsDESC(t *testing.T) {
        order := MARKET_CAP_DESC
	coin1 := MarketData{
	        PriceChangePercentage24H: 1.12345,
	}
	coin2 := MarketData{
	        PriceChangePercentage24H: 2.12345,
	}
	coin3 := MarketData{
	        PriceChangePercentage24H: 3.12345,
	}

        expected := []MarketData{coin3, coin2, coin1}
	result := RankCoins(PCP_DAY, order, []MarketData{coin1, coin2, coin3})

        if !reflect.DeepEqual(expected, result) {
	        t.Errorf("expected: %v, got: %v", expected, result)
	}
}

// test group high liquidity coins
// market cap rank
// total volume
func TestGroupHighLiquidityCoins(t *testing.T) {
        minRank := 2
	maxRank := 3
	minVolume := 2.12345
	coin1 := MarketData{
	        MarketCapRank: 1,
		TotalVolume: 3.12345,
	}
	coin2 := MarketData{
	        MarketCapRank: 2,
		TotalVolume: 1.12345,
	}
	coin3 := MarketData{
	        MarketCapRank: 3,
		TotalVolume: 3.12345,
	}

        expected := []MarketData{coin3}
	result := GroupHighLiquidityCoins(minRank, maxRank, minVolume, []MarketData{coin1, coin2, coin3})

        if !reflect.DeepEqual(expected, result) {
	        t.Errorf("expected: %v, got: %v", expected, result)
	}
}

// test rank coin scarcity
// circulating supply
// max supply
// ath change percentage
func TestRankCoinScarcity(t *testing.T) {
        // circulating_supply / max_supply
        circulatingScore := 2.12345 / 3.12345
	athChangeScore := 2.12345
	coin1 := MarketData{
	        CirculatingSupply: 2.12345,
		MaxSupply: 3.12345,
		AthChangePercentage: 2.12345,
	}
	coin2 := MarketData{
	        CirculatingSupply: 1.12345,
		MaxSupply: 5.12345,
		AthChangePercentage: 1.12345,
	}
	coin3 := MarketData{
	        CirculatingSupply: 4.12345,
		MaxSupply: 7.12345,
		AthChangePercentage: 3.12345,
	}

        expected := []MarketData{coin1}
	result := RankCoinScarcity(circulatingScore, athChangeScore, []MarketData{coin1, coin2, coin3})

        if !reflect.DeepEqual(expected, result) {
	        t.Errorf("expected: %v, got: %v", expected, result)
	}
}