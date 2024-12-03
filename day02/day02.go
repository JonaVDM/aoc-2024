package day02

import (
	"fmt"

	"github.com/jonavdm/aoc-2024/utils"
)

func Run(file string) [2]interface{} {
	data := utils.ReadGridFileNum(file)

	return [2]interface{}{
		PartOne(data),
		0,
	}
}

func PartOne(data [][]int) int {
	var bad int
	for _, report := range data {
		up := report[0] < report[len(report)-1]

		for i := range report {
			if i == 0 {
				continue
			}

			diff := report[i-1] - report[i]
			if utils.AbsInt(diff) > 3 || diff == 0 || (up && diff > 0) || (!up && diff < 0) {
				bad++
				break
			}
		}
	}

	return len(data) - bad
}

func PartTwo(data [][]int) int {
	var bad int
	for row, report := range data {
		first := true
		skip := false
		up := report[0] < report[len(report)-1]

		for i := range report {
			if i == 0 {
				continue
			}

			diff := report[i-1] - report[i]
			if skip {
				diff = report[i-2] - report[i]
				skip = false
			}

			if utils.AbsInt(diff) > 3 || diff == 0 || (up && diff > 0) || (!up && diff < 0) {
				if first {
					first = false
					skip = true
					continue
				}
				bad++
				fmt.Println(row, false)
				break
			}
		}
	}

	return len(data) - bad
}
