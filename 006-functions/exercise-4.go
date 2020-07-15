package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "and I am", p.age, "years old.")
}

func main() {
	// methods
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p1.speak()
}
