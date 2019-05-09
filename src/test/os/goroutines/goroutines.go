package goroutines

import (
	"flag"
	"runtime"
	"fmt"
	"time"
)

func longWait() {
	fmt.Println("Beginning longWait()")
	time.Sleep(5 * 1e9) // sleep for 5 seconds
	fmt.Println("End of longWait()")
}

func shortWait() {
	fmt.Println("Beginning shortWait()")
	time.Sleep(2 * 1e9) // sleep for 2 seconds
	fmt.Println("End of shortWait()")
}

func main()  {
	// gc 编译器下设置 GOMAXPROCS 为一个大于默认值 1 的数值来允许运行时支持使用多于 1 个的操作系统线程
	var numCores = flag.Int("n", 2, "number of CPU cores to use")
	flag.Parse()
	runtime.GOMAXPROCS(*numCores)
	//
	fmt.Println("main start:")
	go longWait()
	go shortWait()
	fmt.Println("main sleep:")
	// sleep works with a Duration in nanoseconds (ns) !
	time.Sleep(10 * 1e9)
	fmt.Println("main end:")
}


