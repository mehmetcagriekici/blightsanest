package crypto

import (
        "time"
        "sync"
)

// Crypto Cache Entry
type cryptoEntry struct {
        CreatedAt time.Time
	Market []MarketData
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
		stopCh:   make(chan struct{}),
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
	        CreatedAt: time.Now(),
		Market: append([]MarketData(nil), market...),
	}

        // unlock the cache
	c.mu.Unlock()
}

// Get an entry from the cache
func (c *CryptoCache) Get(key string) (cryptoEntry, bool) {
        // lock the cache
	c.mu.RLock()

        // check if the key exists
	entry, ok := c.Market[key]
	if !ok {
	        c.mu.RUnlock()
	        return cryptoEntry{}, false
	}

        // check if the entry is stale
        isStale := time.Since(entry.CreatedAt) > c.interval
	c.mu.RUnlock()

        if isStale {
	        // lazily remove under write lock
		c.mu.Lock()

                // re-check in case another goroutine updated it
		if e2, ok2 := c.Market[key]; ok2 && time.Since(e2.CreatedAt) > c.interval {
		        delete(c.Market, key)
		}
		c.mu.Unlock()
		return cryptoEntry{}, false
	}

        // return the entry
	return entry, true
}

// remove old entries
func (c *CryptoCache) reapLoop() {
        ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

        for {
	        select {
		case <-ticker.C:
		        c.mu.Lock()
			for k, m := range c.Market {
			        if time.Since(m.CreatedAt) > c.interval {
				        delete(c.Market, k)
				}
			}
			c.mu.Unlock()
		case <-c.stopCh:
		        return
		}
	}
}

// stop ticker and goroutine
func (c *CryptoCache) Close() {
        close(c.stopCh)
}