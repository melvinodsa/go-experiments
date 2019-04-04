package cache

import (
	"testing"
)

/*
 * This file contains the benchmark for the redis cache
 */

func BenchmarkRedis(b *testing.B) {
	//Setup the data
	RedisSetup()
	for n := 0; n < b.N; n++ {
		req := GetRandom()
		RedisCache(req)
	}
}
