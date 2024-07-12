package core

import (
	"time"
)

type CacheNode struct {
	Key      string    // Key of the cache entry
	Value    string    // Value associated with the key
	ExpireAt time.Time // Expiration time for the cache entry (zero time if no expiration)
}
