package atomic

func AddC(arr []int32, g int) {
	n := len(arr) / g
	in, out := make(chan int32), make(chan int32)
	go merge(in, out, g)

	for i := 0; i < g; i++ {
		go computeC(arr[n*i:n*(i+1)], in)
	}

	<-out
}

func computeC(arr []int32, ch chan int32) {
	var result int32
	for _, v := range arr {
		result += v
	}
	ch <- result
}

func merge(in, out chan int32, g int) {
	var ans int32
	for {
		v := <-in
		ans += v
		g--
		if g == 0 {
			out <- ans
			return
		}
	}
}
