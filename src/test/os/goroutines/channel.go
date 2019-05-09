package goroutines

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)
	go sendData(ch)
	go getData(ch)

	time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "Washington"
	time.Sleep(1000)
	ch <- "Tripoli"
	time.Sleep(1000)
	ch <- "London"
	time.Sleep(1000)
	ch <- "Beijing"
	time.Sleep(1000)
	ch <- "Tokyo"
}

func getData(ch chan string) {
	var input string
	// time.Sleep(2e9)
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
	fmt.Println("")
}
