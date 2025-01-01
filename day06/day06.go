package day06

import (
	"slices"

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

	path := solver.FindPath()

	return [2]interface{}{
		len(path),
		solver.FindBlocks(path),
	}
}

type Solver struct {
	Pos    utils.Point
	Height int
	Width  int
	Field  map[utils.Point]bool
}

func (s *Solver) FindPath() map[utils.Point]bool {
	guard := Guard{
		Field:  s.Field,
		Pos:    s.Pos,
		Height: s.Height,
		Width:  s.Width,
		Dir:    utils.Point{X: 0, Y: -1},
	}

	been := make(map[utils.Point]bool)
	been[s.Pos] = true

	for guard.Move() {
		been[guard.Pos] = true
	}

	return been
}

func (s *Solver) FindBlocks(path map[utils.Point]bool) int {
	keys := utils.GetKeys(path)
	been := make(map[utils.Point]bool)
	count := 0

	for _, key := range keys {
		for _, n := range utils.GetDirectAdjacend(key.X, key.Y, s.Height, s.Width) {
			if _, ok := been[n]; ok || s.Field[n] {
				continue
			}

			field := s.Field
			field[n] = true
			if !s.TryField(field) {
				been[n] = false
			} else {
				been[n] = true
				count++
			}
			field[n] = false
		}
	}

	return count
}

// TryField returns true if the guard runs in circles
func (s *Solver) TryField(field map[utils.Point]bool) bool {
	guard := Guard{
		Field:  s.Field,
		Pos:    s.Pos,
		Height: s.Height,
		Width:  s.Width,
		Dir:    utils.Point{X: 0, Y: -1},
	}

	been := make(map[utils.Point][]utils.Point)

	for guard.Move() {
		if _, ok := been[guard.Pos]; !ok {
			been[guard.Pos] = make([]utils.Point, 0)
		}

		if slices.Contains(been[guard.Pos], guard.Dir) {
			return true
		}

		been[guard.Pos] = append(been[guard.Pos], guard.Dir)
	}

	return false
}
