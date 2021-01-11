package main

import (
	"fmt"
	"github.com/sunimalherath/go-funcs/callback"
)

func main() {
	x := []int{1, 3, 2, 5, 4, 7, 6, 8, 45, 9, 11}

	fmt.Println("Sum of x :", callback.Sum(x...))
	fmt.Println("Evans :", callback.Evens(callback.Sum, x...))
}
