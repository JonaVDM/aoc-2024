package day11

import (
	"math"
	"strconv"
	"strings"

	"github.com/jonavdm/aoc-2024/utils"
)

func Run(data []string) [2]interface{} {
	nums := make([]int, 0)
	for _, c := range strings.Split(data[0], " ") {
		n, _ := strconv.Atoi(c)
		nums = append(nums, n)
	}

	return [2]interface{}{
		Solve(nums, 25),
		Solve(nums, 75),
	}
}

func Solve(input []int, amount int) int {
	nums := make(map[int]int)
	for _, n := range input {
		nums[n]++
	}

	for i := 0; i < amount; i++ {
		newNums := make(map[int]int)
		for n, v := range nums {
			out := ParseNumber(n)
			for _, n := range out {
				newNums[n] += v
			}
		}
		nums = newNums
	}

	var total int
	for _, v := range nums {
		total += v
	}
	return total
}

func ParseNumber(num int) []int {
	if num == 0 {
		return []int{1}
	}

	if utils.NumberLength(num)%2 == 0 {
		return SplitNumber(num)
	}

	return []int{num * 2024}
}

func SplitNumber(num int) []int {
	digits := int(math.Log10(float64(num)) + 1)
	power := int(math.Pow(10, float64(digits/2)))

	first := num / power
	second := num % power

	return []int{first, second}
}
