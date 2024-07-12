package database

import (
	"container/list"
	"fmt"
	"pp/core"
	"strings"
	"sync"
	"time"
)

type LRUCache_Inmemory struct {
	cache map[string]*list.Element // Map for fast access to cache elements
	list  *list.List               // Doubly linked list to track access order

	cleanupTime time.Duration // Time interval for periodic cleanup of expired entries
	mu          sync.Mutex    // Mutex for concurrent access to cache data structures
}

// NewLRUCache initializes and returns a new LRUCache instance.
func NewLRUCache_Inmemory(cleanupTime time.Duration) *LRUCache_Inmemory {
	c := &LRUCache_Inmemory{
		cache: make(map[string]*list.Element),
		list:  list.New(),

		cleanupTime: cleanupTime,
	}
	go c.startCleanupRoutine() // Start a goroutine for periodic cache cleanup
	return c
}

func (c *LRUCache_Inmemory) startCleanupRoutine() {
	ticker := time.NewTicker(c.cleanupTime)
	defer ticker.Stop()
	for range ticker.C {
		c.cleanup()
	}
}

// cleanup removes expired items from the cache.
func (c *LRUCache_Inmemory) cleanup() {
	c.mu.Lock()
	defer c.mu.Unlock()
	now := time.Now()
	for elem := c.list.Front(); elem != nil; {
		next := elem.Next()
		node := elem.Value.(*core.CacheNode)
		if !node.ExpireAt.IsZero() && node.ExpireAt.Before(now) {
			// Remove expired node from the linked list and delete from map
			c.list.Remove(elem)
			delete(c.cache, node.Key)
		}
		elem = next
	}
}

func (c *LRUCache_Inmemory) Print() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	orderedItems := []string{}
	now := time.Now()
	for elem := c.list.Front(); elem != nil; {
		next := elem.Next()
		node := elem.Value.(*core.CacheNode)
		if node.ExpireAt.IsZero() || node.ExpireAt.After(now) {
			orderedItems = append(orderedItems, fmt.Sprintf("%s:%s", node.Key, node.Value))
		} else {
			c.list.Remove(elem)
			delete(c.cache, node.Key)
		}
		elem = next
	}
	return strings.Join(orderedItems, ", ")

}

func (c *LRUCache_Inmemory) Del_all() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.list.Init()                            // Clear the linked list
	c.cache = make(map[string]*list.Element) // Reset the cache map

}

func (c *LRUCache_Inmemory) Del_key(key string) string {
	c.mu.Lock()
	defer c.mu.Unlock()
	if elem, found := c.cache[key]; found {
		c.list.Remove(elem)                               // Remove element from linked list
		delete(c.cache, elem.Value.(*core.CacheNode).Key) // Delete from cache map
		return "key is deleted successfully"
	}
	return "key is not present"

}

func (c *LRUCache_Inmemory) Set(key string, value string, length int, ttl int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if elem, found := c.cache[key]; found {
		node := elem.Value.(*core.CacheNode)
		node.Value = value
		if ttl > 0 {
			node.ExpireAt = time.Now().Add(time.Duration(ttl) * time.Second)
		} else {
			node.ExpireAt = time.Time{} // Reset expiration if ttl <= 0
		}
		c.list.MoveToFront(elem) // Move existing item to the front
		return
	}
	if c.list.Len() >= length {
		c.evict() // Evict least recently used element if cache is full
	}
	// Add new element to the front of the list
	expireAt := time.Time{}
	if ttl > 0 {
		expireAt = time.Now().Add(time.Duration(ttl) * time.Second)
	}
	newNode := &core.CacheNode{Key: key, Value: value, ExpireAt: expireAt}
	entry := c.list.PushFront(newNode)
	c.cache[key] = entry

}

func (c *LRUCache_Inmemory) evict() {
	if evicted := c.list.Back(); evicted != nil {
		c.list.Remove(evicted)
		delete(c.cache, evicted.Value.(*core.CacheNode).Key)
	}
}

func (c *LRUCache_Inmemory) Get(key string) string {
	c.mu.Lock()
	defer c.mu.Unlock()
	if elem, found := c.cache[key]; found {
		node := elem.Value.(*core.CacheNode)
		if node.ExpireAt.IsZero() || node.ExpireAt.After(time.Now()) {
			c.list.MoveToFront(elem) // Move accessed item to the front of the list
			return node.Value
		}
		// Remove the expired element from both the list and the map
		c.list.Remove(elem)
		delete(c.cache, key)
	}
	return "" // Return empty string if key not found or expired

}
