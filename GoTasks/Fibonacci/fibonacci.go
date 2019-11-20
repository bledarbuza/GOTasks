package Fibonacci

func Fibonacci(x int) int {
	var v1, v2, res int
	var vargu []int
	v1 = 1
	v2 = 1

	vargu = append(vargu, v1)
	vargu = append(vargu, v2)

	for i := 0; i < x; i++ {
		res = v1 + v2
		v2 = v1
		v1 = res

		vargu = append(vargu, res)

	}
	return vargu[x-1]
}
