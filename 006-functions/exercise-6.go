package main

import (
	"fmt"
)

func main() {
	// anonymous function
	func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
	}()

	fmt.Println("done")
}
