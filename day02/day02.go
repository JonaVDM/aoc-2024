package day02

import (
	"log/slog"

	"github.com/jonavdm/aoc-2024/utils"
)

func Run(file string) [2]interface{} {
	data := utils.ReadGridFileNum(file)

	return [2]interface{}{
		PartOne(data),
		PartTwo(data),
	}
}

func PartOne(data [][]int) int {
	var safe int
	for _, report := range data {
		if ValidLine(report) {
			safe++
		}
	}

	return safe
}

func PartTwo(data [][]int) int {
	var safe int
	for _, report := range data {
		if ValidLine(report) {
			safe++
			continue
		}

		slog.Debug("", "", report)
		for i := range report {
			copySlice := make([]int, len(report))
			copy(copySlice, report)
			newSlice := append(copySlice[:i], copySlice[i+1:]...)

			if ValidLine(newSlice) {
				safe++
				break
			}
		}
	}

	return safe
}

func ValidLine(line []int) bool {
	up := line[0] < line[len(line)-1]
	for i := range line {
		if i == 0 {
			continue
		}

		diff := line[i-1] - line[i]

		if utils.AbsInt(diff) > 3 || diff == 0 || (up && diff > 0) || (!up && diff < 0) {
			return false
		}
	}

	return true
}
