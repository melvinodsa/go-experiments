package atomic_test

import (
	"runtime"
	"strconv"
	"testing"

	"github.com/melvinodsa/go-experiments/atomic"
	"github.com/melvinodsa/go-experiments/atomic/testdata"
)

var tcs = []int{1000, 10000, 100000, 1000000, 10000000, 100000000}
var nos = []int{1, 2, 3, 10, 100, 1000}

func BenchmarkAdd(b *testing.B) {
	for _, v := range tcs {
		for _, no := range nos {
			b.Run(strconv.Itoa(v)+"routines"+strconv.Itoa(no), func(b *testing.B) {
				testdata.Prepare(v)
				for i := 0; i < b.N; i++ {
					atomic.Add(testdata.Arr, no*runtime.NumCPU())
				}
			})
		}
	}
}
