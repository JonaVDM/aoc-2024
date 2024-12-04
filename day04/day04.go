package day04

import (
	"log/slog"

	"github.com/jonavdm/aoc-2024/utils"
)

type Solver struct {
	Data []string
}

func Run(file string) [2]interface{} {
	solver := Solver{
		Data: utils.ReadFile(file),
	}

	var count int
	for i, row := range solver.Data {
		for j, l := range row {
			if l == 'X' {
				count += solver.FindM(j, i)
			}
		}
	}

	return [2]interface{}{
		count,
		solver.FindMasses(),
	}
}

func (s *Solver) FindM(x, y int) int {
	count := 0
	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if i < 0 || i >= len(s.Data) || j < 0 || j >= len(s.Data[i]) {
				continue
			}

			if s.Data[i][j] == 'M' {
				slog.Debug("XM")
				if s.FindAS(j, i, -(x - j), -(y - i)) {
					count++
				}
			}
		}
	}
	return count
}

func (s *Solver) FindAS(x, y int, xd, yd int) bool {
	if y+yd*2 < 0 || y+yd*2 >= len(s.Data) || x+xd*2 < 0 || x+xd*2 >= len(s.Data[0]) {
		return false
	}

	if s.Data[y+yd][x+xd] != 'A' {
		return false
	}

	if s.Data[y+yd*2][x+xd*2] != 'S' {
		return false
	}

	return true
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

			if s.CharAt(j+1, i+1, 'A') && s.CharAt(j+2, i+2, 'S') {
				if aMap[i+1][j+1] {
					aCount++
				} else {
					aMap[i+1][j+1] = true
				}
			}

			if s.CharAt(j+1, i-1, 'A') && s.CharAt(j+2, i-2, 'S') {
				if aMap[i-1][j+1] {
					aCount++
				} else {
					aMap[i-1][j+1] = true
				}
			}

			if s.CharAt(j-1, i+1, 'A') && s.CharAt(j-2, i+2, 'S') {
				if aMap[i+1][j-1] {
					aCount++
				} else {
					aMap[i+1][j-1] = true
				}
			}

			if s.CharAt(j-1, i-1, 'A') && s.CharAt(j-2, i-2, 'S') {
				if aMap[i-1][j-1] {
					aCount++
				} else {
					aMap[i-1][j-1] = true
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
