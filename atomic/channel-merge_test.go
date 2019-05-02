package atomic_test

import (
	"runtime"
	"strconv"
	"testing"

	"github.com/melvinodsa/go-experiments/atomic"
	"github.com/melvinodsa/go-experiments/atomic/testdata"
)

var m = 8

func BenchmarkAddMC(b *testing.B) {
	for _, v := range tcs {
		for _, no := range nos {
			b.Run(strconv.Itoa(v)+"routines"+strconv.Itoa(no), func(b *testing.B) {
				testdata.Prepare(v)
				for i := 0; i < b.N; i++ {
					mc := no / m
					if mc == 0 {
						mc = 1
					}
					atomic.AddMC(testdata.Arr, no*runtime.NumCPU(), mc)
				}
			})
		}
	}
}
