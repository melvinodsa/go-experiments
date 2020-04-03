package cache

import (
	"math/rand"
	"sync"
	"testing"
)

/*
 * This file contains the tests for the cache
 */

const NUM_GOROUTINES = 10
const NUM_DATAPOINTS = 10000

const ALPHA = "abcdefghijklmnopqrstunwxyz"

func sampleData() map[string]int64 {
	var sampleData = map[string]int64{}
	for i, c := range ALPHA {
		sampleData[string(c)] = int64(i)
	}
	return sampleData
}
func randKey() string {
	key := ALPHA[rand.Intn(len(ALPHA))]
	return string(key)
}

func runCacheGetter(wg *sync.WaitGroup, cacheRunner CacheLike) {
	sum := int64(0)
	for n := 0; n < NUM_DATAPOINTS; n++ {
		key := randKey()
		val, ok := cacheRunner.Get(key)
		if !ok {
			continue
		}
		intval, ok := val.(int64)
		if !ok {
			continue
		}
		sum = sum + intval
	}
	wg.Done()
	// return sum
}

func genRandomReq() Request {
	request := NewRequest(READ)
	request.Key = randKey()
	return request
}
func runCacheChan(wg *sync.WaitGroup) {
	sum := int64(0)
	for n := 0; n < NUM_DATAPOINTS; n++ {
		req := genRandomReq()
		RequestChannel <- req
		<-req.Out
		intval, ok := req.Payload.(int64)
		if !ok {
			continue
		}
		sum = sum + intval
	}
	wg.Done()
	// return sum
}

func setupMutexCache() *mutexCache {
	mcache := NewMutexCache()
	for k, v := range sampleData() {
		mcache.Put(k, v)
	}
	return mcache
}

func BenchmarkCacheWithMutex(b *testing.B) {
	b.ReportAllocs()
	//Setup the data
	mcache := setupMutexCache()
	var wg sync.WaitGroup
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for n := 0; n < NUM_GOROUTINES; n++ {
			wg.Add(1)
			go runCacheGetter(&wg, mcache)
		}
		wg.Wait()
	}
}

func setupRWMutexCache() *rwMutexCache {
	mcache := NewRWMutexCache()
	for k, v := range sampleData() {
		mcache.Put(k, v)
	}
	return mcache
}

func BenchmarkCacheWithRWMutex(b *testing.B) {
	b.ReportAllocs()
	//Setup the data
	mcache := setupRWMutexCache()
	var wg sync.WaitGroup
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for n := 0; n < NUM_GOROUTINES; n++ {
			wg.Add(1)
			go runCacheGetter(&wg, mcache)
		}
		wg.Wait()
	}
}

func BenchmarkCacheWithChan(b *testing.B) {
	b.ReportAllocs()
	//Setup the data
	CacheSetup()
	var wg sync.WaitGroup
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for n := 0; n < NUM_GOROUTINES; n++ {
			wg.Add(1)
			go runCacheChan(&wg)
		}
		wg.Wait()
	}
}
