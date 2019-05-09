package goroutines

import (
	"fmt"
	"time"
)

func f1(in chan int) {
	time.Sleep(1000)
	fmt.Println(<-in) //
}

func main() {
	out := make(chan int) // 准备发送
	go f1(out) // 若是注释此条语句，则通道接收者未准备好，通道阻塞
	out <- 2
	// go f1(out)
	time.Sleep(2000)
}