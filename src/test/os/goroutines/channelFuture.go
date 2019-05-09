package goroutines

import "fmt"

type Matrix struct {
	a,b int
}

func InverseFuture(a Matrix) chan int {
	future := make(chan int)
	go func() {
		future <- a.a*a.b
	}()
	return future
}

func main() {
	a := Matrix{1,2}
	b := Matrix{3,4}
	a_inv_future := InverseFuture(a)   // start as a goroutine
	b_inv_future := InverseFuture(b)   // start as a goroutine
	a_inv := <-a_inv_future
	b_inv := <-b_inv_future
	fmt.Println(a_inv+b_inv)
}
