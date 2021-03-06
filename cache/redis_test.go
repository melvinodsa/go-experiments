package cache

import (
	"sync"
	"testing"
)

/*
 * This file contains the benchmark for the redis cache
 */

func BenchmarkRedis(b *testing.B) {
	//Setup the data
	RedisSetup()
	var wg sync.WaitGroup
	for n := 0; n < b.N; n++ {
		wg.Add(1)
	}
	for n := 0; n < b.N; n++ {
		go wr(&wg)
	}
	wg.Wait()
}

func wr(wg *sync.WaitGroup) {
	req := GetRandom()
	RedisCache(req)
	wg.Done()
}
