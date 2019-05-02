package atomic_test

import (
	"runtime"
	"strconv"
	"testing"

	"github.com/melvinodsa/go-experiments/atomic"
	"github.com/melvinodsa/go-experiments/atomic/testdata"
)

func BenchmarkAddC(b *testing.B) {
	for _, v := range tcs {
		for _, no := range nos {
			b.Run(strconv.Itoa(v)+"routines"+strconv.Itoa(no), func(b *testing.B) {
				testdata.Prepare(v)
				for i := 0; i < b.N; i++ {
					atomic.AddC(testdata.Arr, no*runtime.NumCPU())
				}
			})
		}
	}
}
