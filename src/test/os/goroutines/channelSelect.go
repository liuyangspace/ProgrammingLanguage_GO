package goroutines

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2)

	time.Sleep(1e9)
}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		fmt.Printf("Request 1 start:%d\n", i)
		ch<-i
	}
}

func pump2(ch chan int) {
	for i := 0; ; i++ {
		fmt.Printf("Request 2 start:%d\n", i)
		ch<-i
	}
}

func suck(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("Request 1 handle:%d\n", v)
			time.Sleep(1000)
			fmt.Printf("Request 1 end:%d\n", v)
		case v := <-ch2:
			fmt.Printf("Request 2 handle:%d\n", v)
			time.Sleep(1000)
			fmt.Printf("Request 2 end:%d\n", v)
		}
	}
}