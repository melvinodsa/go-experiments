package testdata

var Arr []int32

func Prepare(n int) {
	Arr = make([]int32, n)
	for i := 0; i < n; i++ {
		Arr[i] = int32(i)
	}
}
