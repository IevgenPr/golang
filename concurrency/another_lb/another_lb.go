package main

import (
	"math/rand"
	"time"
)

// Request ssss
type Request struct {
	fn func() int // operation to perform
	c  chan int   // return channel to return
}

// Worker blah
type Worker struct {
	requests chan Request // work to do  - buffered channel
	pending  int          // count of pending
	index    int          // index of pending job
}

func (w *Worker) work(done chan *Worker) {
	for {
		req := <-w.requests
		req.c <- req.fn()
		done <- w
	}
}

func requestor(work chan<- Request) {
	c := make(chan int)
	for {
		time.Sleep(time.Duration(rand.Int63n()) * time.Millisecond)(Worker * 2) //Int54n
		work <- Request{workFn, c}
		result := <-c
		//furtherProcess(result)
	}
}

func main() {}
