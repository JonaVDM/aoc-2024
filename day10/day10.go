package day10

import (
	"log/slog"
	"slices"

	"github.com/jonavdm/aoc-2024/utils"
)

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)
	var total int
	var scores int

	s := Solver{
		grid: utils.ConverToGridInts(data),
	}
	s.height = len(s.grid)
	s.width = len(s.grid[0])

	for y, row := range s.grid {
		for x, c := range row {
			if c == 0 {
				amount := s.FindPaths(x, y)
				slog.Default().Debug("Found a 0", "Pos", utils.Point{X: x, Y: y}, "Amount", amount)
				total += amount
				scores += s.FindScores(x, y)
			}
		}
	}

	return [2]interface{}{
		total,
		scores,
	}
}

type Solver struct {
	grid          [][]int
	width, height int
}

type entry struct {
	point  utils.Point
	search int
}

func (s *Solver) On(p utils.Point) int {
	return s.grid[p.Y][p.X]
}

func (s *Solver) FindPaths(x, y int) int {
	var total int
	q := make([]entry, 1)
	q[0] = entry{utils.Point{X: x, Y: y}, 1}
	been := make([]utils.Point, 1)
	been[0] = utils.Point{X: x, Y: y}

	for len(q) > 0 {
		e := q[0]
		q = q[1:]
		for _, p := range utils.GetDirectAdjacend(e.point.X, e.point.Y, s.height, s.width) {
			if slices.Contains(been, p) {
				continue
			}

			if s.On(p) == e.search && e.search == 9 {
				total++
				been = append(been, p)
				continue
			}

			if s.On(p) == e.search {
				been = append(been, p)
				q = append(q, entry{p, e.search + 1})
			}
		}
	}

	return total
}

func (s *Solver) FindScores(x, y int) int {
	var total int
	q := make([]entry, 1)
	q[0] = entry{utils.Point{X: x, Y: y}, 1}

	for len(q) > 0 {
		e := q[0]
		q = q[1:]
		for _, p := range utils.GetDirectAdjacend(e.point.X, e.point.Y, s.height, s.width) {
			if s.On(p) == e.search && e.search == 9 {
				total++
				continue
			}

			if s.On(p) == e.search {
				q = append(q, entry{p, e.search + 1})
			}
		}
	}

	return total
}
