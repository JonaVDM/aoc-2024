package utils

import (
	"strconv"
	"strings"
)

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

func ConverToGridInts(input []string, seperator string) [][]int {
	out := make([][]int, len(input))
	for r, row := range input {
		val := make([]int, 0)
		for _, c := range strings.Split(row, seperator) {
			con, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			val = append(val, con)
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
