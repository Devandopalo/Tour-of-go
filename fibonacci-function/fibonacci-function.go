package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	x0 := 0
	x1 := 0
	return func() int {
		xn := x0 + x1
		if xn == 0 {
			x1++
			return xn
		} else {
			x0 = x1
			x1 = xn
			return xn
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
