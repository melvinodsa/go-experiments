package atomic

import (
	"sync"
	"sync/atomic"
)

func Add(arr []int32, g int) {
	var ans int32
	wg := sync.WaitGroup{}
	wg.Add(g)
	n := len(arr) / g

	for i := 0; i < g; i++ {
		go compute(arr[n*i:n*(i+1)], &ans, &wg)
	}

	wg.Wait()
}

func compute(arr []int32, ans *int32, wg *sync.WaitGroup) {
	var result int32
	for _, v := range arr {
		result += v
	}
	atomic.AddInt32(ans, result)
	wg.Done()
}
