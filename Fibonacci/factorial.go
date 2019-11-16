package main

import(
	"fmt"
)

var fact int = 1
var i int = 1
var n int

func factorial(n int) int {
	if (n < 0) {
		fmt.Print("Nr > 0")

	} else {
		for i = 0; i < n; i++ {
			fact *= int(i)
		}
	}
	return fact
}

