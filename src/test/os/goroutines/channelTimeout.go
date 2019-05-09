package goroutines

import (
	"time"
)

func main()  {
	const timeoutNs  = 1
	ch := make(chan int, 1)
	go func() { for { ch <- 1 } } ()
	//outime := time.NewTimer(timeoutNs)
	outime := time.After(timeoutNs)
L:
	for {
		select {
		case <-ch:
			// do something
			println("task do ...")
			//case <-outime.C:
		case <-outime:
			// call timed out
			println("task timeout ...")
			break L
			//return
		}
	}
}