package goroutines

import (
	"fmt"
	"time"
)

type Empty interface {}
const N  =  6

func doSomething(i int,x float64) float64 {
	fmt.Printf("Data\t%d\tStart\n",i)
	defer fmt.Printf("Data\t%d\tEnd\n",i)
	time.Sleep(1000)
	return x
}

func main()  {
	var empty Empty
	data := make([]float64, N)
	res := make([]float64, N)
	sem := make(chan Empty, N)
	for i, xi := range data {
		go func (i int, xi float64) {
			res[i] = doSomething(i, xi)
			sem <- empty
		} (i, xi)
	}
	// wait for goroutines to finish
	for i := 0; i < N; i++ {
		<-sem
	}
}