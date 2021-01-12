package recursion

func Factorial(n int) int {
	if n != 0 {
		return n * Factorial(n -1)
	}
	return 1
}
