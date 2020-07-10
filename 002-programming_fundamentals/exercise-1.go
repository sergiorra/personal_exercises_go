package main

import (
	"fmt"
)

func main() {
	// decimal, binary, hexadecimal
	a := 42
	fmt.Printf("%d\t%b\t%#x", a, a, a)
}
