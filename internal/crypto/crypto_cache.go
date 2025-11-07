package crypto

import (
        "time"
        "sync"
)

// Crypto Cache Entry
type cryptoEntry struct {
        createdAt time.Time
	market []MarketData
}

// Crypto Cache
type CryptoCache struct {
        Market   map[string]cryptoEntry
	mu       sync.RWMutex
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
func (c *CryptoCache) Add(key string, market []MarketData) {
        // lock the cache
	c.mu.Lock()

        // add a new entry
	c.Market[key] = cryptoEntry{
	        createdAt: time.Now(),
		market: append([]MarketData(nil), market...),
	}

        // unlock the cache
	c.mu.Unlock()
}

// Get an entry from the cache
func (c *CryptoCache) Get(key string) (MarketData[], bool) {
        // lock the cache
	c.mu.RLock()

        // check if the key exists
	entry, ok := c.Market[key]
	if !ok {
	        c.mu.RUnlock()
	        return nil, false
	}

        // check if the entry is stale
        isStale := time.Since(entry.createdAt) > c.interval
	val := entry.market
	c.mu.RUnlock()

        if isStale {
	        // lazily remove under write lock
		c.mu.Lock()

                // re-check in case another goroutine updated it
		if e2, ok2 := c.market[key]; ok2 && time.Since(e2.createdAt) > c.interval {
		        delete(c.market, key)
		}
		c.mu.Unlock()
		return nil, false
	}

        // return the value
	out := append([]MarketData(nil), val...)
	return out, true
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
		        // iterate over the cache entries and delete the expired entries			                    for k, m := range c.Market {
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