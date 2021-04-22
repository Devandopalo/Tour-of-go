package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	arrY := make([][]uint8, dy, dy)

	for yi, _ := range arrY {
		arrX := make([]uint8, dx, dx)
		for xi, _ := range arrX {
			arrX[xi] = uint8((xi + yi) / 2)
		}
		arrY[yi] = arrX
	}

	return arrY
}

func main() {
	pic.Show(Pic)
}
