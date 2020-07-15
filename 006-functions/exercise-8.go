package main

import "fmt"

func main() {
	// functions returning functions
	f := foo()
	fmt.Println(f())
}

func foo() func() int {
	return func() int {
		return 42
	}
}
