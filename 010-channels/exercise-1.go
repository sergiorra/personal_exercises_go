package main

import (
	"fmt"
)

func main() {
	// channels
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
