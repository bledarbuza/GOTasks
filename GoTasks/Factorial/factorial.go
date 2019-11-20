package Factorial

func Factorial(x int) int {

	var vargu []int
	var res int = 1
	for i := x; i >= 1; i-- {
		res = res * i
		vargu = append(vargu, res)
	}
	return vargu[x-1]
}
