package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	// NumWorkers blah
	NumWorkers = 10
)

// Work is some struct
type Work struct {
	x, y, z int
}

func worker(in <-chan *Work, out chan<- *Work) {
	for w := range in {
		w.z = w.x * w.y
		time.Sleep(time.Duration(w.z) * time.Second)
		out <- w
	}
}

// Run blah
func Run() {
	in, out := make(chan *Work), make(chan *Work)
	for i := 0; i < NumWorkers; i++ {
		go worker(in, out)
	}
	go sendLotsOfWork(in)
	receiveLotsOfResults() <- out
}

// we can potentially create a generator which produces Work items and puts them into in channel
func sendLotsOfWork(in chan<- *Work) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		in <- Work{rand.Intn(10), rand.Intn(10), rand.Intn(10)}
	}
}

func (out chan *Work) receiveLotsOfResults() {
	select {
	case v := <-out:
		fmt.Println(v)
	}

}

func main() {
	Run()
}
