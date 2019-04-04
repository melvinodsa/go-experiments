package cache

import (
	"testing"
)

/*
 * This file contains the tests for the cache
 */

func BenchmarkCache(b *testing.B) {
	//Setup the data
	CacheSetup()
	for n := 0; n < b.N; n++ {
		req := GetRandom()
		RequestChannel <- req
		if req.Type == READ {
			<-req.Out
		}
	}
}
