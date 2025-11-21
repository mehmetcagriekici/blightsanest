package crypto

import(
        "sync"
)

type PriceRange struct {
        Min float64
	Max float64
}

type CryptoState struct {
        Order         AvailableOrders
	Timeframes    []AvailableTimeframes
	CoinsList     []MarketData
	MinVolume     float64
	MinChange     float64
	MaxRank       int
	MinMaxPrice   PriceRange
	SupplyCap     bool
	ExcludeStable bool
	mu            sync.RWMutex
}

// function to create a new crypt state with default values
func CreateCryptoState() *CryptoState {
        return &CryptoState{
	        Order:         MARKET_CAP_DESC,
		Timeframes:    []AvailableTimeframes{},
		CoinsList:     []MarketData{},
		MinVolume:     1000000,
		MinChange:     0,
		MaxRank:       250,
		MinMaxPrice:   PriceRange{Min: 1, Max: 100,},
		SupplyCap:     false,
		ExcludeStable: false,
	}
}

// update order
func (cs *CryptoState) UpdateOrder(newOrder AvailableOrders) {
        cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.Order = newOrder
}

// update timeframes
func (cs *CryptoState) UpdateTimeframes(newFrames []AvailableTimeframes) {
        cs.mu.Lock()
	defer cs.mu.Unlock()	
	cs.Timeframes = append([]AvailableTimeframes(nil), newFrames...)
}

// update min volume
func (cs *CryptoState) UpdateMinVolume(newVolume float64) {
        cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.MinVolume = newVolume
}

// update min change
func (cs *CryptoState) UpdateMinChange(newChange float64) {
        cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.MinChange = newChange
}

// update max rank
func (cs *CryptoState) UpdateMaxRank(newRank int) {
        cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.MaxRank = newRank
}

// update price range
func (cs *CryptoState) UpdatePriceRange(min, max float64) {
        cs.mu.Lock()
	defer cs.mu.Unlock()
	newRange := PriceRange{
	        Min: min,
		Max: max,
	}
	cs.MinMaxPrice = newRange
}

// update supply cap preference
func (cs *CryptoState) UpdateSupplyCap(newPreference bool) {
        cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.SupplyCap = newPreference
}
 
// update exclude stable preference
func (cs *CryptoState) UpdateExcludeStable(newPreference bool) {
        cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.ExcludeStable = newPreference
}

// update the coins list
func (cs *CryptoState) UpdateCoinsList(newList []MarketData) {
        cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.CoinsList = newList
}