package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "ip"
	num := 3.2

	ptr := &num
	printVar(name)
	fmt.Println("Name is ", name, " type: ", reflect.TypeOf(name))
	fmt.Println("num: ", num, " type: ", reflect.TypeOf(num))
	printVar(*ptr)
}

func printVar(x interface{}) string {
	fmt.Println("x: ", x, " type: ", reflect.TypeOf(x))
	return "0"
}
