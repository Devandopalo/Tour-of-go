package main

import (
	"fmt"
	"math"
)

const esp float64 = 0.00001

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return -1, ErrNegativeSqrt(x)
	}

	i := 0
	z0, z1 := 1.0, 0.1
	for math.Abs(z0*z0-x) > esp {
		z1 -= (z0*z0 - x) / (2 * z0)
		z0 = z1
		i++
	}
	return z0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
