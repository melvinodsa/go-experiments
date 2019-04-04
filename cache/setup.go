package cache

import (
	"math/rand"
)

//GetRandom will return a random request
func GetRandom() Request {
	requests := []Request{
		NewRequest(READ),
		NewRequest(WRITE),
		NewRequest(DELETE),
	}
	keys := []string{"Foo", "Coolboy", "My Heros"}
	request := requests[rand.Intn(len(requests))]
	request.Key = keys[rand.Intn(len(keys))]
	request.Payload = DataStore[request.Key]
	return request
}

//CacheSetup will setup the cache
func CacheSetup() {
	for k, v := range DataStore {
		req := NewRequest(WRITE)
		req.Key = k
		req.Payload = v
		RequestChannel <- req
	}
}

//RedisSetup will setup the redis data
func RedisSetup() {
	for k, v := range DataStore {
		req := NewRequest(WRITE)
		req.Key = k
		req.Payload = v
		RedisCache(req)
	}
}
