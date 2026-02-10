package crypto

import(
        "sync"
	"math"
)

type PriceRange struct {
        Min float64
	Max float64
}

type CryptoState struct {
	CurrentList                     []MarketData
	CurrentListID                   string
	CurrentOrder                    AvailableOrders
        CurrentSortingField             string
	ClientTimeframes                []AvailableTimeframes
	CurrentTimeframe                AvailableTimeframes
	CurrentMinRank                  int
	CurrentMaxRank                  int
	CurrentMinVolume                float64
	CurrentMaxVolume                float64
	CurrentMinCirculatingSupply     float64
	CurrentMaxATHChangePercentage   float64
	CurrentMinMarketCap             float64
	CurrentMaxMarketCap             float64
	CurrentMaxPriceChangePercentage float64
	CurrentMinPriceChangePercentage float64
	CurrentMinSwingScore            float64
	CurrentMaxSwingScore            float64
	CurrentMinSupply                float64
	CurrentIgnoredCoins             []string
	CurrentMinVolatility            float64
	CurrentMaxVolatility            float64
	CurrentMinGrowthPotential       float64
	CurrentMaxGrowthPotential       float64
	CurrentMinLiquidity             float64
	CurrentMaxLiquidity             float64
	mu                              sync.RWMutex
}

// function to create a new crypt state with default values
func CreateCryptoState() *CryptoState {
        return &CryptoState{
	        CurrentList:                     []MarketData{},
		CurrentListID:                   "",
		CurrentSortingField:             "CurrentPrice",
		ClientTimeframes:                []AvailableTimeframes{},
		CurrentTimeframe:                PCP_DAY,
		CurrentMinRank:                  0,
		CurrentMaxRank:                  250,
		CurrentMinVolume:                math.Inf(-1),
		CurrentMaxVolume:                math.Inf(+1),
		CurrentMinCirculatingSupply:     math.Inf(-1),
		CurrentMaxATHChangePercentage:   math.Inf(+1),
		CurrentMinMarketCap:             math.Inf(-1),
		CurrentMaxMarketCap:             math.Inf(+1),
		CurrentMinPriceChangePercentage: math.Inf(-1),
		CurrentMaxPriceChangePercentage: math.Inf(+1),
		CurrentOrder:                    CRYPTO_ASC,
		CurrentMinSwingScore:            math.Inf(-1),
		CurrentMaxSwingScore:            math.Inf(+1),
		CurrentMinSupply:                math.Inf(-1),
		CurrentIgnoredCoins:             []string{},
		CurrentMinVolatility:            math.Inf(-1),
		CurrentMaxVolatility:            math.Inf(+1),
		CurrentMinGrowthPotential:       math.Inf(-1),
		CurrentMinLiquidity:             math.Inf(-1),
	}
}

// update sorting field
func (cs *CryptoState) UpdateCurrentSortingField(field string) {
        cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.CurrentSortingField = field
}

// update current liquidity
func (cs *CryptoState) UpdateCurrentLiquidity(minLiquidity float64) {
        cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.CurrentMinLiquidity = minLiquidity
}

// update current growth potential
func (cs *CryptoState) UpdateGrowthPotential(minPotential float64) {
        cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.CurrentMinGrowthPotential = minPotential
}

// update current volatility
func (cs *CryptoState) UpdateCurrentVolatility(minVolatility, maxVolatility float64) {
        cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.CurrentMinVolatility = minVolatility
	cs.CurrentMaxVolatility = maxVolatility
}

// update current ignored coins
func (cs *CryptoState) UpdateIgnoredCoins(ignoredCoins []string) {
        cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.CurrentIgnoredCoins = ignoredCoins
}

// update current potential coin supply
func (cs *CryptoState) UpdateSupply(minSupply float64) {
        cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.CurrentMinSupply = minSupply
}

// update current swing score
func (cs *CryptoState) UpdateCurrentSwingScore(minScore, maxScore float64) {
        cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.CurrentMinSwingScore = minScore
	cs.CurrentMaxSwingScore = maxScore
}

// update current timeframe
func (cs *CryptoState) UpdateCurrentTimeframe(timeframe AvailableTimeframes) {
        cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.CurrentTimeframe = timeframe
}

// update order
func (cs *CryptoState) UpdateOrder(order AvailableOrders) {
        cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.CurrentOrder = order
}

// update price change percentage
func (cs *CryptoState) UpdatePriceChangePercentage(minChange, maxChange float64) {
        cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.CurrentMinPriceChangePercentage = minChange
	cs.CurrentMaxPriceChangePercentage = maxChange
} 

// update market cap
func (cs *CryptoState) UpdateMarketCap(minCap, maxCap float64) {
        cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.CurrentMinMarketCap = minCap
	cs.CurrentMaxMarketCap = maxCap
}

// update circulating supply
func (cs *CryptoState) UpdateCirculatingSupply(minSupply float64) {
        cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.CurrentMinCirculatingSupply = minSupply
}

// update ath change percentage
func (cs *CryptoState) UpdateAthChangePercentage(maxAth float64) {
        cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.CurrentMaxATHChangePercentage = maxAth
}

// update the coins list
func (cs *CryptoState) UpdateCurrentList(id string, newList []MarketData) {
        cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.CurrentList = newList
	cs.CurrentListID = id
}

// set client's timeframes
func (cs *CryptoState) UpdateClientTimeframes(frames []AvailableTimeframes) {
        cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.ClientTimeframes = frames
}

// update market cap rank preferences
func (cs *CryptoState) UpdateMarketRank(minRank, maxRank int) {
        cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.CurrentMinRank = minRank
	cs.CurrentMaxRank = maxRank
}

// update coin volume preferences
func (cs *CryptoState) UpdateVolume(minVolume, maxVolume float64) {
        cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.CurrentMinVolume = minVolume
	cs.CurrentMaxVolume = maxVolume
}
