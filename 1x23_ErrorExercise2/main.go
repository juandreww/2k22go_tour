package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Unable to square root %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x == 1 {
		return 1, nil
	} else if x < 0 {
		return x, ErrNegativeSqrt(x)
	}

	z := 1.0
	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
