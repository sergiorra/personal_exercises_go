package main

import (
	"fmt"
)

func main() {
	// functions
	n := foo()
	x, s := bar()

	fmt.Println(n, x, s)
}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 1984, "Big Brother is Watching"
}
