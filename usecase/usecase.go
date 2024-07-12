package usecase

import "pp/database"

// Backends struct holds the Redis and in-memory caches
type Backends struct {
	redis    *database.LRUCache_Redis
	inmemory *database.LRUCache_Inmemory
}

// Init initializes the Backends with Redis and in-memory caches
func Init(redis1 *database.LRUCache_Redis, inmemory *database.LRUCache_Inmemory) *Backends {
	return &Backends{
		redis:    redis1,
		inmemory: inmemory,
	}
}

// Print returns the current state of the caches
// If both caches are the same, it returns the cache content
// Otherwise, it indicates that the data is not the same in both backends
func (r *Backends) Print() string {
	output_redis := r.redis.Print()
	output_inmemory := r.inmemory.Print()
	if output_inmemory == output_redis {
		return output_redis
	} else {
		return "data is not same in both backends"
	}
}

// Del_all deletes all entries from both Redis and in-memory caches
func (r *Backends) Del_all() {
	r.redis.Del_all()
	r.inmemory.Del_all()
}

// Del_key deletes a specified key from both Redis and in-memory caches
// If the key is successfully deleted from both, it returns a success message
// Otherwise, it indicates that the data is different in both backends
func (r *Backends) Del_key(key string) string {
	output_redis := r.redis.Del_key(key)
	output_inmemory := r.inmemory.Del_key(key)
	if output_inmemory == output_redis {
		return "key is deleted successfully"
	} else {
		return "data in both backends are different"
	}
}

// Set adds or updates a key-value pair in both Redis and in-memory caches
func (r *Backends) Set(key string, value string, length int, time int) {
	r.redis.Set(key, value, length, time)
	r.inmemory.Set(key, value, length, time)
}

// Get retrieves a value from both Redis and in-memory caches
// If both caches have the same value, it returns the value
// If the key is not present in both, it returns a message indicating that
// Otherwise, it indicates that the data is different in both backends
func (r *Backends) Get(key string) string {
	output_redis := r.redis.Get(key)
	output_inmemory := r.inmemory.Get(key)
	if output_inmemory == output_redis {
		if output_inmemory == "" {
			return "key is not present"
		}
		return output_redis
	} else {
		return "data is not same in both backends"
	}
}
