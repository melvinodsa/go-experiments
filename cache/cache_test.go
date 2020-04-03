package cache

import (
	"sync"
	"testing"
)

/*
 * This file contains the tests for the cache
 */

func BenchmarkCache(b *testing.B) {
	//Setup the data
	CacheSetup()
	var wg sync.WaitGroup
	for n := 0; n < b.N; n++ {
		wg.Add(1)
	}
	for n := 0; n < b.N; n++ {
		go wc(&wg)
	}
	wg.Wait()
}

func wc(wg *sync.WaitGroup) {
	req := GetRandom()
	RequestChannel <- req
	if req.Type == READ {
		<-req.Out
	}
	wg.Done()
}
