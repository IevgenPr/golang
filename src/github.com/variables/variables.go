package main

import (
	"fmt"
	"reflect"
)

var (
	// "module" variable is unused in code, and compiler does not complain!
	module = "golang module var"
	intnum = 10
)

func main() {
	name := "ip"
	num := 3.2

	ptr := &num //pointer to address of num

	printVar(name)
	printVar(intnum)
	fmt.Println("Name is ", name, " type: ", reflect.TypeOf(name))
	fmt.Println("num: ", num, " type: ", reflect.TypeOf(num))
	printVar(*ptr) // here pointer is dereferenced
}

// function names that start with lowercase - are private within a module
// function names that start with capital letter - available in other modules
func printVar(x interface{}) string {
	fmt.Println("x: ", x, " type: ", reflect.TypeOf(x))
	return "0"
}
