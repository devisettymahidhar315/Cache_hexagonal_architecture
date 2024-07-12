package usecase

import "pp/database"

type Backends struct {
	redis    *database.LRUCache_Redis
	inmemory *database.LRUCache_Inmemory
}

func Init(redis1 *database.LRUCache_Redis, inmemory *database.LRUCache_Inmemory) *Backends {
	return &Backends{
		redis:    redis1,
		inmemory: inmemory,
	}
}

func (r *Backends) Print() string {
	output_redis := r.redis.Print()
	output_inmemory := r.inmemory.Print()
	if output_inmemory == output_redis {
		return output_redis
	} else {
		return "data is not same in both backends"
	}

}

func (r *Backends) Del_all() {
	r.redis.Del_all()
	r.inmemory.Del_all()

}

func (r *Backends) Del_key(key string) string {
	output_redis := r.redis.Del_key(key)
	output_inmemory := r.inmemory.Del_key(key)
	if output_inmemory == output_redis {
		return "key is deleted successfully"

	} else {
		return "data in both backends are different"
	}
}

func (r *Backends) Set(key string, value string, length int, time int) {
	r.redis.Set(key, value, length, time)
	r.inmemory.Set(key, value, length, time)

}

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
