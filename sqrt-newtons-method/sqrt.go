package main

import (
	"fmt"
	"math"
)

const esp float64 = 0.000001

func Sqrt(px *float64) float64 {
	i := 0
	z0, z1 := 1.0, 0.1
	for math.Abs(z0*z0-*px) > esp {
		z1 -= (z0*z0 - *px) / (2 * z0)
		z0 = z1
		i++
	}
	return z1
}

func main() {
	x := 444.333222111
	fmt.Println(Sqrt(&x))
}
