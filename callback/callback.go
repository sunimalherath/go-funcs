package callback

// Sum - accepts integers and return the sum of it.
func Sum(xi ...int) int {
	var sum int
	for _, v := range xi {
		sum += v
	}
	return sum
}

// Evens - accepts func like Sum and integers and return the total
func Evens(f func(xi ...int) int, ii ...int) int {
	var ei []int
	for _, v := range ii {
		if v%2 == 0 {
			ei = append(ei, v)
		}
	}
	return f(ei...)
}
