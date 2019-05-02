package atomic

func AddMC(arr []int32, g, m int) {
	n := len(arr) / g
	l := m - (g % m)
	finalout, out := make(chan int32), make(chan int32)
	ins := []chan int32{}
	for i := 0; i < l; i++ {
		in := make(chan int32)
		go merge(in, out, g/l)
		ins = append(ins, in)
	}
	go merge(out, finalout, l)

	for i := 0; i < g; i++ {
		go computeC(arr[n*i:n*(i+1)], ins[i%l])
	}

	<-finalout
}
