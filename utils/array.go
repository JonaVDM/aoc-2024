package utils

import "slices"

func IntRange(start, stop int) []int {
	out := make([]int, 0)
	for i := start; i < stop; i++ {
		out = append(out, i)
	}
	return out
}

func Overlap(a, b []int) bool {
	for _, x := range a {
		for _, y := range b {
			if x == y {
				return true
			}
		}
	}

	return false
}

func ListContains[T comparable](list []T, items []T) bool {
	for _, item := range items {
		if slices.Contains(list, item) {
			return true
		}
	}

	return false
}
