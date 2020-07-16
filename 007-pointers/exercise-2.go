package main

import (
	"fmt"
)

type person struct {
	name string
}

func main() {
	// passing pointers as reference
	p1 := person{
		name: "James Bond",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

func changeMe(p *person) {
	p.name = "Miss Moneypenny"
	// (*p).name = "Miss Moneyp"
}
