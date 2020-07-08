package main

import "fmt"

func main () {
	fmt.Println("Hello world!!")
	start()
	for i := 0; i < 20; i++ {
		if i % 2 == 0 {
			println(i)
		}
	}
	end()
}

func start() {
	println("Starting...")
}

func end() {
	println("Ending...")
}
