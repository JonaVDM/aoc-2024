package day01

import (
	"sort"
	"strconv"
	"strings"

	"github.com/jonavdm/aoc-2024/utils"
)

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)

	left, right := make([]int, 0), make([]int, 0)
	rightCount := make(map[int]int, 0)

	for _, row := range data {
		parts := strings.Split(row, "   ")

		l, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		r, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		left = append(left, l)
		right = append(right, r)

		rightCount[r]++
	}

	sort.Ints(left)
	sort.Ints(right)

	var diff int
	var sim int

	for i := range left {
		d := left[i] - right[i]
		diff += utils.AbsInt(d)

		sim += left[i] * rightCount[left[i]]
	}

	return [2]interface{}{
		diff,
		sim,
	}
}

