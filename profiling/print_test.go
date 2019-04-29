package profiling_test

import (
	"testing"

	"github.com/melvinodsa/go-experiments/profiling"
)

var testcases = []profiling.Printer{
	{I: "asd"},
	{I: nil},
	{I: 10},
}

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range testcases {
			v.String()
		}
	}
}

var fmttestcases = []profiling.FmtPrinter{
	{I: "asd"},
	{I: nil},
	{I: 10},
}

func BenchmarkFmtString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range fmttestcases {
			v.String()
		}
	}
}
