package goroutines

import (
	"time"
	"fmt"
)

func f2(in chan int) {
	fmt.Println(<-in) //
}

func main() {
	out := make(chan int,2) // 准备发送
	out <- 2
	go f2(out)
	time.Sleep(10000)
}