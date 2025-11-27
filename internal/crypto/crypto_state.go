package crypto

import(
        "sync"
	"strings"
	"math"
)

type PriceRange struct {
        Min float64
	Max float64
}

type CryptoState struct {
	CurrentList                     []MarketData
	CurrentListID                   string
	ClientTimeframes                []string
	CurrentTimeframe                AvailableTimeframes
	CurrentMinRank                  int
	CurrentMaxRank                  int
	CurrentMinVolume                float64
	CurrentMaxVolume                float64
	CurrentMinCirculatingSupply     float64
	CurrentMaxCirculatingSupply     float64
	CurrentMaxATHChangePercentage   float64
	CurrentMinATHChangePercentage   float64
	CurrentMinMarketCap             float64
	CurrentMaxMarketCap             float64
	CurrentMaxPriceChangePercentage float64
	CurrentMinPriceChangePercentage float64
	CurrentOrder                    AvailableOrders
	CurrentMinSwingScore            float64
	CurrentMaxSwingScore            float64
	CurrentMinSupply                float64
	CurrentMaxSupply                float64
	CurrentIgnoredCoins             []string
	CurrentMinVolatility            float64
	CurrentMaxVolatility            float64
	mu                              sync.RWMutex
}

// function to create a new crypt state with default values
func CreateCryptoState() *CryptoState {
        return &CryptoState{
	        CurrentList:                     []MarketData{},
		CurrentListID:                   "",
		ClientTimeframes:                []string{},
		CurrentTimeframe:                PCP_DAY,
		CurrentMinRank:                  0,
		CurrentMaxRank:                  250,
		CurrentMinVolume:                math.Inf(-1),
		CurrentMaxVolume:                math.Inf(+1),
		CurrentMinCirculatingSupply:     math.Inf(-1),
		CurrentMaxCirculatingSupply:     math.Inf(+1),
		CurrentMaxATHChangePercentage:   math.Inf(+1),
		CurrentMinATHChangePercentage:   math.Inf(-1),
		CurrentMinMarketCap:             math.Inf(-1),
		CurrentMaxMarketCap:             math.Inf(+1),
		CurrentMinPriceChangePercentage: math.Inf(-1),
		CurrentMaxPriceChangePercentage: math.Inf(+1),
		CurrentOrder:                    CRYPTO_ASC,
		CurrentMinSwingScore:            math.Inf(-1),
		CurrentMaxSwingScore:            math.Inf(+1),
		CurrentMinSupply:                math.Inf(-1),
		CurrentMaxSupply:                math.Inf(+1),
		CurrentIgnoredCoins:             []string{},
		CurrentMinVolatility:            math.Inf(-1),
		CurrentMaxVolatility:            math.Inf(+1),
	}
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
func (cs *CryptoState) UpdateSupply(minSupply, maxSupply float64) {
        cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.CurrentMinSupply = minSupply
	cs.CurrentMaxSupply = maxSupply
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
func (cs *CryptoState) UpdateCirculatingSupply(minSupply, maxSupply float64) {
        cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.CurrentMinCirculatingSupply = minSupply
	cs.CurrentMaxCirculatingSupply = maxSupply
}

// update ath change percentage
func (cs *CryptoState) UpdateAthChangePercentage(minAth, maxAth float64) {
        cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.CurrentMinATHChangePercentage = minAth
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
func (cs *CryptoState) SetTimeframes(key string) {
        cs.mu.Lock()
	defer cs.mu.Unlock()
	// get the timeframes from the key
	keyParts := strings.Split(key, "__")
	frames := strings.Split(keyParts[0], "_")
	cs.ClientTimeframes = strings.Split(frames[1], "-")
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