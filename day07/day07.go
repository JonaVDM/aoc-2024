package day07

import (
	"slices"
	"strconv"
	"strings"

	"github.com/jonavdm/aoc-2024/utils"
)

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)

	valid := 0

	for _, row := range data {
		search, nums := Parse(row)
		slices.Reverse(nums)
		valid += Sovleble(search, nums)
	}

	return [2]interface{}{
		valid,
		0,
	}
}

func Parse(row string) (int, []int) {
	row = strings.ReplaceAll(row, ":", "")
	spl := strings.Split(row, " ")
	nums := make([]int, len(spl))
	for i, num := range spl {
		con, _ := strconv.Atoi(num)
		nums[i] = con
	}

	return nums[0], nums[1:]
}

func Sovleble(search int, nums []int) int {
	solves := Solve(nums)
	for _, solve := range solves {
		if solve == search {
			return search
		}
	}
	return 0
}

func Solve(nums []int) []int {
	if len(nums) == 1 {
		return []int{nums[0]}
	}

	out := make([]int, 0)
	for _, num := range Solve(nums[1:]) {
		out = append(out, nums[0]+num, nums[0]*num)
	}

	return out
}
