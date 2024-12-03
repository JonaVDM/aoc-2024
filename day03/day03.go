package day03

import (
	"log/slog"
	"regexp"
	"strconv"

	"github.com/jonavdm/aoc-2024/utils"
)

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)

	totalA := 0
	totalB := 0
	enabled := true

	for _, line := range data {
		re := regexp.MustCompile(`(?m)mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
		for _, match := range re.FindAllStringSubmatch(line, -1) {
			if match[0] == "don't()" {
				enabled = false
				continue
			}

			if match[0] == "do()" {
				enabled = true
				continue
			}

			l, _ := strconv.Atoi(match[1])
			r, _ := strconv.Atoi(match[2])
			slog.Default().Debug("Match found",
				"match", match[0],
				"left", l,
				"right", r,
				"mul", l*r,
			)

			totalA += l * r
			if enabled {
				totalB += l * r
			}
		}
	}

	return [2]interface{}{
		totalA,
		totalB,
	}
}
