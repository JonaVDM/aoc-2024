package utils

import "strconv"

func ConvertToInts(input []string) []int {
	out := make([]int, len(input))

	for i, v := range input {
		con, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		out[i] = con
	}

	return out
}

func ConverToGridInts(input []string) [][]int {
	out := make([][]int, len(input))
	for r, row := range input {
		val := make([]int, len(row))
		for i, c := range row {
			con, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			val[i] = con
		}
		out[r] = val
	}
	return out
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
