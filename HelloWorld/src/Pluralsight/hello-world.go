package main

import (
  "fmt" // go standard lib
  "runtime"
  "reflect"
)

var (
  name, course  string
  module        float64
  tests, test2, fl = "tests", "test2", 10.0
)
// entry point
func main() {
  a := 10.0000000
  b := 3
  c := int(a) + b
  ptr := &a

  fmt.Println(ptr, *ptr) //dereferencing
  fmt.Println("c = ", c, " type of c ", reflect.TypeOf(c))
  fmt.Println("Hello from", runtime.GOOS)
  fmt.Println("Name is set to ", name, " type is ", reflect.TypeOf(name))
  fmt.Println("module is set to", module, " type is ", reflect.TypeOf(module))
}
