package main

// This is sample from Rob Pike preso
import (
	"fmt"
	"math"
)

// this is closure which takes two functions and returns third function
// and it is important to have opening brace with func keyword on the same line - whaaaat?
// well, you can insert a semicolon - that is fine )))
// well, you can put newline after opening brace )))
func Compose(
	f, g func(
		x float64) float64) func(
	x float64) float64 {
	return func(
		x float64) float64 {
		//here closure returns function as a value
		return f(
			g(x))
	}
}

//f, g func(x float) float
// func Compose(f, g func(x float64) float64) func(x float64) float64 {
// 	return func(x float64) float64 {
// 		return f(g(x))
// 	}
// }

func main() {
	//print(Compose(sin, cos)(0.5))

	fmt.Println("Returning closure ", Compose(math.Sin, math.Cos)(0.5))

}
