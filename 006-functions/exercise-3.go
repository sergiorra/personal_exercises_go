package main

import (
	"fmt"
)

func main() {
	// defer functions
	defer foo()
	fmt.Println("Hello, playground")
}

func foo() {
	defer func() {
		fmt.Println("foo DEFER ran")
	}()
	fmt.Println("foo ran")
}
