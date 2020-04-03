package cache

import "sync"

/*
 * Thsi file cotains the definitions for the high performant cache
 */

type mutexCache struct {
	cache map[string]interface{}
	mux   sync.Mutex
}

func NewMutexCache() *mutexCache {
	return &mutexCache{
		cache: map[string]interface{}{},
	}

}

//GetCacheValue gets the cache value
func (c *mutexCache) Get(key string) (interface{}, bool) {
	c.mux.Lock()
	v, ok := internalCache.cache[key]
	c.mux.Unlock()
	return v, ok
}

//PutCacheValue puts in the cache value
func (c *mutexCache) Put(key string, value interface{}) {
	c.mux.Lock()
	c.cache[key] = value
	c.mux.Unlock()
}
