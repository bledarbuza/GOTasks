package main

import "fmt"

func main(){
	nr := int(9)
	fmt.Println(fibonacci(nr))

	fmt.Println("Shtyp numrin : ")
	fmt.Scanln(&n)
	fmt.Print(factorial(n))

	var n,i,j int
	j = 0

	fmt.Println("Shkruaj numrin: ")
	fmt.Scanln(&n)

	fmt.Println("Fibonacci : ")
	for  i = 1 ; i <= n-1; i++ {
		fmt.Print(fibonacci1(j), " , ")
		j++
	}
	fmt.Println()




	v := func(){


		func (a, b int){
			fmt.Print(a+b)
		}(5, 5)
	}
	v()
}
