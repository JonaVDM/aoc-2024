package day01

import (
	"sort"
	"strconv"
	"strings"

	"github.com/jonavdm/aoc-2024/utils"
)

func Run(data []string) [2]interface{} {
	left, right := make([]int, len(data)), make([]int, len(data))
	rightCount := make(map[int]int, 0)

	for i, row := range data {
		var err error
		parts := strings.Split(row, "   ")

		left[i], err = strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		right[i], err = strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		rightCount[right[i]]++
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
