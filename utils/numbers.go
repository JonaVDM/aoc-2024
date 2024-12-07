package utils

import "math"

func NumberLength(i int) int {
	if i == 0 {
		return 1
	}
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}

func PowInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
