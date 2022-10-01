package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	if x == 1 {
		return 1
	}

	z := 1.0
	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
