package main

import (
	"fmt"
)

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	// struct
	p1 := person{
		first: "James",
		last:  "Bond",
		favFlavors: []string{
			"chocolate",
			"martini",
			"rum and coke",
		},
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		favFlavors: []string{
			"strawberry",
			"vanilla",
			"capuccino",
		},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favFlavors {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.favFlavors {
		fmt.Println(i, v)
	}
}
