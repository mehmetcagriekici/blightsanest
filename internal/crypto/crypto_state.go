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
	CurrentList                   []MarketData
	CurrentListID                 string
	ClientTimeframes              []string
	CurrentMinRank                int
	CurrentMaxRank                int
	CurrentMinVolume              float64
	CurrentMaxVolume              float64
	CurrentMinCirculatingSupply   float64
	CurrentMaxCirculatingSupply   float64
	CurrentMaxATHChangePercentage float64
	CurrentMinATHCHangePercentage float64
	mu                            sync.RWMutex
}

// function to create a new crypt state with default values
func CreateCryptoState() *CryptoState {
        return &CryptoState{
	        CurrentList:                   []MarketData{},
		CurrentListID:                 "",
		ClientTimeframes:              []string{},
		CurrentMinRank:                0,
		CurrentMaxRank:                250,
		CurrentMinVolume:              math.Inf(-1),
		CurrentMaxVolume:              math.Inf(+1),
		CurrentMinCirculatingSupply:   math.Inf(-1),
		CurrentMaxCirculatingSupply:   math.Inf(+1),
		CurrentMaxATHChangePercentage: math.Inf(+1),
		CurrentMinATHCHangePercentage: math.Inf(-1),
	}
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
	parts := strings.Split(key, "__")
	frames := strings.Split(parts[0], "_")
	cs.ClientTimeframes = strings.Split(frames[1], "-")
}

// update market cap preferences
func (cs *CryptoState) UpdateMarketCap(minRank, maxRank int) {
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