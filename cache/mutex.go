package cache

import "sync"

/*
 * Thsi file cotains the definitions for the high performant cache
 */

type interlCache struct {
	cache map[string]interface{}
	mux   sync.Mutex
}

var internalCache = interlCache{cache: LoadPersistant()}

//GetCacheValue gets the cache value
func GetCacheValue(key string) interface{} {
	internalCache.mux.Lock()
	v, _ := internalCache.cache[key]
	internalCache.mux.Unlock()
	return v
}

//PutCacheValue puts in the cache value
func PutCacheValue(key string, value interface{}) {
	internalCache.mux.Lock()
	internalCache.cache[key] = value
	internalCache.mux.Unlock()
}

//MutexCache for processing the mutex based cache
func MutexCache(req Request) Request {
	if req.Type == READ {
		req.Payload = GetCacheValue(req.Key)
	} else if req.Type == WRITE {
		PutCacheValue(req.Key, req.Payload)
	}
	return req
}
