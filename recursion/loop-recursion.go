package recursion

func LoopFactorial(x int) int {
	sum := 1
	for ; x > 0; x-- {
		sum *= x
	}
	return sum
}
