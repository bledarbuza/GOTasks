package main

import(

)
func fibonacci( nr int) int {
	if nr == 0 || nr == 1 {
		return nr

	} else {
		return (fibonacci(nr-1) + fibonacci(nr-2))
	}
}
