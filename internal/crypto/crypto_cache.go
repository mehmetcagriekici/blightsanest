package crypto

import (
        "time"
        "sync"
)

// Crypto Cache Entry
type cryptoEntry struct {
        createdAt time.Time
	market MarketData[]
}

// Crypto Cache
type CryptoCache struct {
        Market   map[string]cryptoEntry
	mu       sync.Mutex
	interval time.Duration
	stopCh   chan struct{}
}

// create a new cache
func CreateCryptoCache(interval time.Duration) *CryptoCache {
        c := &CryptoCache{
	        Market:   make(map[string]cryptoEntry),
		interval: interval,
		stopCh:   make(chan struct{})
	}

        // start a goroutine to delete the old entries
	go c.reapLoop()
	
	return c
}

// Add a new entry to the cache
func (c *CryptoCache) Add(key string, market MarketData[]) {
        // lock the cache
	c.mu.Lock()

        // add a new entry
	c.Market[key] = cryptoEntry{
	        createdAt: time.Now(),
		market: market,
	}

        // unlock the cache
	c.mu.Unlock()
}

// Get an entry from the cache
func (c *CryptoCache) Get(key string) (MarketData[], bool) {
        // lock the cache
	c.mu.Lock()
	// unlock the cache on returns
        defer c.mu.Unlock()

        // check if the key exists
	entry, ok := c.Market[key]
	if !ok {
	        return nil, false
	}

        // check if the entry is stale
	if time.Since(entry.createdAt) > c.interval {
	        // delete the entry
		delete(c.Market, key)
		return nil, false
	}

        // return the value
	return entry.market, true
}

// remove old entries
func (c *CryptoCache) reapLoop() {
        // set a ticker
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	
        // go
	for {
	        select {
		        case <-ticker.C:
			        // lock the cache
			        c.mu.Lock()
			        // iterate over the cache entries and delete the expired entries
			        for k, m := range c.Market {
			                if time.Since(m.createdAt) > c.interval {
				                delete(c.Market, k)
				        }
			        }
			        // unlock the cache
			        c.mu.Unlock()
			case <-c.stopCh:
			        // exit the loop
			        return
		}
	}
}

// stop ticker and goroutine
func (c *CryptoCache) Close() {
        close(c.stopCh)
}