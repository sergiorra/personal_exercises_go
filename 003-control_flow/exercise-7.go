package main

import (
	"fmt"
)

func main() {
	// switch without expression
	switch {
	case false:
		fmt.Println("should not print")
	case true:
		fmt.Println("should print")
	}
}
