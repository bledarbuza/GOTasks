package main
import (


)
func fibonacci1( n int) int{
	if n == 0 || n == 1{
		return n

	}else{
		return(fibonacci(n-1)+fibonacci(n-2))
	}

}