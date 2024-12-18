package day04

import (
	"github.com/jonavdm/aoc-2024/utils"
)

type Solver struct {
	Data []string
}

func Run(data []string) [2]interface{} {
	solver := Solver{
		Data: data,
	}

	return [2]interface{}{
		solver.FindXMas(),
		solver.FindMasses(),
	}
}

func (s *Solver) FindXMas() int {
	var count int
	for i, row := range s.Data {
		for j, char := range row {
			if char != 'X' {
				continue
			}

			for _, p := range utils.GetRelativeAdjecend() {
				if s.CharsInRow(j, i, p, "MAS") {
					count++
				}
			}
		}
	}

	return count
}

func (s *Solver) FindMasses() int {
	aMap := make(map[int]map[int]bool)
	for i := range s.Data {
		aMap[i] = make(map[int]bool)
	}

	var aCount int
	for i, row := range s.Data {
		for j, char := range row {
			if char != 'M' {
				continue
			}

			points := []utils.Point{
				{X: +1, Y: +1},
				{X: +1, Y: -1},
				{X: -1, Y: +1},
				{X: -1, Y: -1},
			}

			for _, p := range points {
				if s.CharsInRow(j, i, p, "AS") {
					if aMap[i+p.Y][j+p.X] {
						aCount++
					} else {
						aMap[i+p.Y][j+p.X] = true
					}
				}
			}
		}
	}

	return aCount
}

func (s *Solver) CharAt(x, y int, char rune) bool {
	if y < 0 || y >= len(s.Data) || x < 0 || x >= len(s.Data[y]) {
		return false
	}

	return s.Data[y][x] == byte(char)
}

func (s *Solver) CharsInRow(x, y int, direction utils.Point, letters string) bool {
	for i, l := range letters {
		if !s.CharAt(x+direction.X*(i+1), y+direction.Y*(i+1), l) {
			return false
		}
	}

	return true
}
