package crypto

import(
        "sync"
	"strings"
)

type PriceRange struct {
        Min float64
	Max float64
}

type CryptoState struct {
	CurrentList      []MarketData
	CurrentListID    string
	ClientTimeframes []string
	mu               sync.RWMutex
}

// function to create a new crypt state with default values
func CreateCryptoState() *CryptoState {
        return &CryptoState{
	        CurrentList:   []MarketData{},
		CurrentListID: "",
		ClientTimeframes: []string{},
	}
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
	frames := strings.Split(parts[0], "_"0)
	cs.ClientTimeframes := strings.Split(frames[1], "-")
}