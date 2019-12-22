// Check this: https://golang.org/doc/code.html
package main

import (
	"fmt"
	"os"
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

	const constString = "Constant String"

	printVar(name)
	printVar(intnum)
	fmt.Println("Name is ", name, " type: ", reflect.TypeOf(name))
	fmt.Println("num: ", num, " type: ", reflect.TypeOf(num))
	printVar(*ptr) // here pointer is dereferenced

	fmt.Println(">>> Passing arguments by value ")
	fmt.Println("\nName before..")
	printVar(name)
	changeName(name)
	fmt.Println("\nName after...")
	printVar(name)

	fmt.Println("\n\n>>> Passing arguments by reference ")

	fmt.Println("\nName before..")
	printVar(name)
	changeNameByRef(&name)
	fmt.Println("\nName after...")
	printVar(name)

	printVar(constString)
	//cannot do that  - asssign to const - verified by compiler
	//constString = "test"
	printEnvVars()

	name = os.Getenv("USER")
	// sets env var for duration of application run
	os.Setenv("TEST_ENV", "test environment variable")
	testEnv := os.Getenv("TEST_ENV")
	printVar(testEnv)
	printEnvVars()
}

// function names that start with lowercase - are private within a module
// function names that start with capital letter - available in other modules
func printVar(x interface{}) string {
	fmt.Println("x: ", x, " type: ", reflect.TypeOf(x))
	return "0"
}

func changeName(name string) string {
	name += ", what a name!"
	fmt.Println("\nName inside...")
	printVar(name)
	return name
}

func changeNameByRef(name *string) string {
	*name = "Another name"
	fmt.Println("name inside: ")
	printVar(*name)
	return *name
}

func printEnvVars() string {
	fmt.Println("\nEnvironment variables:")
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
	return ""
}
