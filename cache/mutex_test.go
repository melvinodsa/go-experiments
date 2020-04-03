package cache

import (
	"sync"
	"testing"
)

/*
 * This file contains the tests for the cache
 */

func BenchmarkMutexCache(b *testing.B) {
	//Setup the data
	CacheSetup()
	var wg sync.WaitGroup
	for n := 0; n < b.N; n++ {
		wg.Add(1)
	}
	for n := 0; n < b.N; n++ {
		go wm(&wg)
	}
	wg.Wait()
}

func wm(wg *sync.WaitGroup) {
	req := GetRandom()
	MutexCache(req)
	wg.Done()
}
