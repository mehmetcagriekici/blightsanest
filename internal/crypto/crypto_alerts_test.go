package crypto

import(
        "testing"
	"reflect"
)

// tested fields
// High24H
// Name
func TestCoinsNewHigh(t *testing.T) {
        coin1Old := buildCoinsNewHigh("coin1", 1.12345)
	coin1New := buildCoinsNewHigh("coin1", 2.12345)
        coin2Old := buildCoinsNewHigh("coin2", 3.12345)
	coin2New := buildCoinsNewHigh("coin2", 1.12345)
	coin3Old := buildCoinsNewHigh("coin3", 4.12345)
	coin4New := buildCoinsNewHigh("coin4", 5.12345)
   
        oldCoins := []MarketData{coin1Old, coin2Old, coin3Old}
	newCoins := []MarketData{coin1New, coin2New, coin4New}
	expected := []MarketData{coin1New}
        result := CoinsNewHigh(oldCoins, newCoins)
	
        if !reflect.DeepEqual(expected, result) {
	        t.Errorf("expected: %v, got: %v", expected, result)
	}
}

// tested fields
// PriceChange24H TIMEFRAME = PCP_DAY
func TestCoinsHighPriceSpike(t *testing.T) {
        tresholdRate := 2.12345
        coin1 := buildCoinsHighPriceSpike(1.12345)
	coin2 := buildCoinsHighPriceSpike(2.12345)
	coin3 := buildCoinsHighPriceSpike(3.12345)

        expected := []MarketData{coin2, coin3}
	result := CoinsHighPriceSpike(tresholdRate, PCP_DAY, []MarketData{coin1, coin2, coin3})

        if !reflect.DeepEqual(expected, result) {
	        t.Errorf("expected: %v, got: %v", expected, result)
	}
}

// tested fields
// AthChangePercentage
func TestCoinsGetCloseAthChange(t *testing.T) {
        maxAthChange := 2.12345
	coin1 := buildCoinsGetAthChange(1.12345)
	coin2 := buildCoinsGetAthChange(2.12345)
	coin3 := buildCoinsGetAthChange(3.12345)

        expected := []MarketData{coin1, coin2}
	result := CoinsGetCloseAthChange(maxAthChange, []MarketData{coin1, coin2, coin3})

        if !reflect.DeepEqual(expected, result) {
	        t.Errorf("expected: %v, got: %v", expected, result)
	}
         
}

// Tested fields
// MarketCapRank
// CurrentPrice
// CirculatingSupply
// Name
func TestCoinsHighCirculatingSupply(t *testing.T) {
        alertValue := 1.12345 * 2.12345
	alertMarketRank := 2
	ignoreCoins := []string{"coin2"}
	coin1 := buildCoinsCirculatingSupply(1, 1.12345, 2.12345, "coin1")
	coin2 := buildCoinsCirculatingSupply(2, 3.12345, 4.12345, "coin2")
	coin3 := buildCoinsCirculatingSupply(3, 5.12345, 6.12345, "coin3")

        expected := []MarketData{coin1}
	result := CoinsHighCirculatingSupply(alertMarketRank, alertValue, ignoreCoins, []MarketData{coin1, coin2, coin3})

        if !reflect.DeepEqual(expected, result) {
	        t.Errorf("expected: %v, got: %v", expected, result)
	}

}

func buildCoinsNewHigh(name string, high24h float64) MarketData {
        return MarketData{
	        Name: name,
		High24H: high24h,
	}
}

func buildCoinsHighPriceSpike(priceChange float64) MarketData {
        return MarketData{
	        PriceChangePercentage24H: priceChange,
	}
}

func buildCoinsGetAthChange(athChange float64) MarketData {
        return MarketData{
	        AthChangePercentage: athChange,
	}
}

func buildCoinsCirculatingSupply(marketRank int, currentPrice, circulatingSupply float64, name string) MarketData {
        return MarketData{
	        Name: name,
		MarketCapRank: marketRank,
		CirculatingSupply: circulatingSupply,
		CurrentPrice: currentPrice,
	}
}