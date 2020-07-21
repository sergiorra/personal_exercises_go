package main

import "fmt"

func main() {
	// pulling the numbers off the channel
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}
