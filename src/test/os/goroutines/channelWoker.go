package goroutines

import (
	"fmt"
	"time"
)

const N2 = 6

type Task struct {
	code int
	// some state
}

func sendWork(task chan *Task)  {
	for i := 0; i < N2; i++ {   // start N goroutines to do work
		task<-&Task{i}
		time.Sleep(10)
	}
}

func Worker(in, out chan *Task) {
	for {
		t := <-in
		//...
		fmt.Printf("do task %d\n",t.code)
		out <- t
	}
}

func main() {
	pending, done := make(chan *Task), make(chan *Task)
	go sendWork(pending)       // put tasks with work on the channel
	for i := 0; i < N2; i++ {   // start N goroutines to do work
		go Worker(pending, done)
	}
	//consumeWork(done)          // continue with the processed tasks
	time.Sleep(10e8)
}
