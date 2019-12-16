package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	go func() { //goroutine
		defer waitGrp.Done()

		time.Sleep(5 * time.Second) // blocks execution for 5 sec
		fmt.Println("Hi there")
	}() // parens make func self executing

	go func() {
		defer waitGrp.Done()

		fmt.Println("Niccce")
	}()

	waitGrp.Wait()

}
