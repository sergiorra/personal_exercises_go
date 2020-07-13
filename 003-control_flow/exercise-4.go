package main

import (
	"fmt"
)

func main() {
	// for without condition
	bd := 1985
	for {
		if bd > 2017 {
			break
		}
		fmt.Println(bd)
		bd++
	}
}
