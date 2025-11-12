package crypto

import(
        "testing"
	"reflect"
)

// test filter coin volume
// total volume
func TestFilterCoinVolume(t *testing.T) {
        minVolume := 2.12345
	coin1 := MarketData{
	        TotalVolume: 1.12345,
	}
	coin2 := MarketData{
	        TotalVolume: 2.12345,
	}
	coin3 := MarketData{
	        TotalVolume: 3.12345,
	}

        expected := []MarketData{coin2, coin3}
	result := FilterCoinVolume(minVolume, []MarketData{coin1, coin2, coin3})

        if !reflect.DeepEqual(expected, result) {
	        t.Errorf("expected: %v, got: %v", expected, result)
	}
}

// test filter coin cap
// market cap
func TestFilterCoinCap(t *testing.T) {
        minCap := 2.12345
	coin1 := MarketData{
	        MarketCap: 1.12345,
	}
	coin2 := MarketData{
	        MarketCap: 2.12345,
	}
	coin3 := MarketData{
	        MarketCap: 3.12345,
	}

        expected := []MarketData{coin2, coin3}
	result := FilterCoinCap(minCap, []MarketData{coin1, coin2, coin3})

        if !reflect.DeepEqual(expected, result) {
	        t.Errorf("expected: %v, got: %v", expected, result)
	}
}

// test filter coin price change
// price change
func TestFilterCoinPriceChange(t *testing.T) {
        minChange := 2.12345
	coin1 := MarketData{
	        PriceChangePercentage24H: 1.12345,
	}
	coin2 := MarketData{
	        PriceChangePercentage24H: 2.12345,
	}
	coin3 := MarketData{
	        PriceChangePercentage24H: 3.12345,
	}

        expected := []MarketData{coin2, coin3}
	result := FilterCoinPriceChange(minChange, PCP_DAY, []MarketData{coin1, coin2, coin3})

        if !reflect.DeepEqual(expected, result) {
	        t.Errorf("expected: %v, got: %v", expected, result)
	}
}

// test find wild swing coins
// high24h
// low24h
func TestFindWildSwingCoins(t *testing.T) {
        minSwing := 3.12345 / 2.12345
	coin1 := MarketData{
	        High24H: 4.12345,
		Low24H: 1.12345,
	}
	coin2 := MarketData{
	        High24H: 3.12345,
		Low24H: 2.12345,
	}
	coin3 := MarketData{
	        High24H: 2.12345,
		Low24H: 2.12345,
	}

        expected := []MarketData{coin1, coin2}
	result := FindWildSwingCoins(minSwing, []MarketData{coin1, coin2, coin3})

        if !reflect.DeepEqual(expected, result) {
	        t.Errorf("expected: %v, got: %v", expected, result)
	}
}

// test search coin
// name
func TestSearchCoinPositive(t *testing.T) {
        name := "btc"
        coin1 := MarketData{
	        Name: "btc",
	}
	coin2 := MarketData{
	        Name: "eth",
	}
	coin3 := MarketData{
	        Name: "",
	}

        res, status := SearchCoin(name, []MarketData{coin1, coin2, coin3})
	if !status {
	        t.Errorf("Expected: %v, got: %v", coin1, res)
	}
	if !reflect.DeepEqual(res, coin1) {
	        t.Errorf("Expected: %v, got: %v", coin1, res)
	}
}

func TestSearchCoinNegative(t *testing.T) {
        name := "coin"
        coin1 := MarketData{
	        Name: "btc",
	}
	coin2 := MarketData{
	        Name: "eth",
	}
	coin3 := MarketData{
	        Name: "",
	}

        res, status := SearchCoin(name, []MarketData{coin1, coin2, coin3})
	if status {
	        t.Errorf("Expected: %v, got: %v", MarketData{}, res)
	}
}

// test flag risk coins
// ath change percentage
// total volume
func TestFlagRiskCoins(t *testing.T) {
        maxAthChange := 2.12345
	minVolume := 3.12345
	coin1 := MarketData{
	        AthChangePercentage: 1.12345,
		TotalVolume: 4.12345,
	}
	coin2 := MarketData{
	        AthChangePercentage: 2.12345,
		TotalVolume: 3.12345,
	}
	coin3 := MarketData{
	        AthChangePercentage: 3.12345,
		TotalVolume: 4.12345,
	}
	coin4 := MarketData{
	        AthChangePercentage: 1.12345,
		TotalVolume: 2.12345,
	}

        expected := []MarketData{coin1, coin2}
        result := FlagRiskCoins(maxAthChange, minVolume, []MarketData{coin1, coin2, coin3, coin4})

        if !reflect.DeepEqual(expected, result) {
	        t.Errorf("expected: %v, got: %v", expected, result)
	}
}

// test flag safe coins
// market cap rank
// price change
func TestFlagSafeCoins(t *testing.T) {
        maxMarketRank := 3
	maxPriceChange := 2.12345
	coin1 := MarketData{
	        MarketCapRank: 1,
		PriceChangePercentage24H: 1.12345,
	}
        coin2 := MarketData{
	        MarketCapRank: 2,
		PriceChangePercentage24H: 3.12345,
	}
	coin3 := MarketData{
	        MarketCapRank: 4,
		PriceChangePercentage24H: 1.12345,
	}

        expected := []MarketData{coin1}
	result := FlagSafeCoins(maxMarketRank, maxPriceChange, PCP_DAY, []MarketData{coin1, coin2, coin3})

        if !reflect.DeepEqual(expected, result) {
	        t.Errorf("expected: %v, got: %v", expected, result)
	}
}