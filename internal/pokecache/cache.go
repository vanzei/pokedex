package pokecache

import (
    "sync"
    "time"
)

// Cache represents a simple in-memory cache with an eviction mechanism.
type Cache struct {
    mu       sync.Mutex
    entries  map[string]cacheEntry
    interval time.Duration
}

type cacheEntry struct {
    createdAt time.Time
    value     []byte
}

// NewCache creates a new cache with a configurable eviction interval.
func NewCache(interval time.Duration) *Cache {
    c := &Cache{
        entries:  make(map[string]cacheEntry),
        interval: interval,
    }

    // Start the reap loop in a separate goroutine
    go c.reapLoop()

    return c
}

// Add adds a new value to the cache and resets the creation time.
func (c *Cache) Add(key string, val []byte) {
    c.mu.Lock()
    defer c.mu.Unlock()

    c.entries[key] = cacheEntry{
        createdAt: time.Now(),
        value:     val,
    }
}

// Get retrieves the value from the cache. It returns the value and a boolean indicating if the key was found.
func (c *Cache) Get(key string) ([]byte, bool) {
    c.mu.Lock()
    defer c.mu.Unlock()

    entry, found := c.entries[key]
    if !found {
        return nil, false
    }
    return entry.value, true
}

// reapLoop continuously checks if cache entries should be evicted based on the interval.
func (c *Cache) reapLoop() {
    ticker := time.NewTicker(5 * time.Second)
    defer ticker.Stop()

    for range ticker.C {
        c.mu.Lock()
        now := time.Now()
        for key, entry := range c.entries {
            if now.Sub(entry.createdAt) > c.interval {
                delete(c.entries, key)
            }
        }
        c.mu.Unlock()
    }
}