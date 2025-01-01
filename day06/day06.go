package day06

import (
	"github.com/jonavdm/aoc-2024/utils"
)

func Run(data []string) [2]interface{} {
	solver := Solver{
		Field:  make(map[utils.Point]bool),
		Height: len(data),
		Width:  len(data[0]),
	}

	for y, row := range data {
		for x, l := range row {
			if l == '.' {
				continue
			}

			if l == '^' {
				solver.Pos = utils.Point{X: x, Y: y}
				continue
			}

			if l != '#' {
				panic("dunno")
			}

			solver.Field[utils.Point{X: x, Y: y}] = true
		}
	}

	return [2]interface{}{
		solver.FindCount(),
		0,
	}
}

type Solver struct {
	Pos    utils.Point
	Height int
	Width  int
	Field  map[utils.Point]bool
}

func (s *Solver) FindCount() int {
	guard := Guard{
		Field:  s.Field,
		Pos:    s.Pos,
		Height: s.Height,
		Width:  s.Width,
		Dir:    utils.Point{X: 0, Y: -1},
	}

	been := make(map[utils.Point]bool)
	been[s.Pos] = true

	for {
		if !guard.Move() {
			break
		}
		been[guard.Pos] = true
	}

	return len(been)
}
