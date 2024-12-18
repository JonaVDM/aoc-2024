package day12

import (
	"log/slog"

	"github.com/jonavdm/aoc-2024/utils"
)

func Run(data []string) [2]interface{} {
	return [2]interface{}{
		PartOne(data),
		0,
	}
}

type solver struct {
	been map[utils.Point]bool
	data []string
}

func PartOne(data []string) int {
	s := solver{
		been: make(map[utils.Point]bool),
		data: data,
	}

	var total int

	for y, row := range data {
		for x, let := range row {
			total += s.FindSimple(utils.Point{X: x, Y: y}, let)
		}
	}

	return total
}

func (s *solver) FindSimple(start utils.Point, letter rune) int {
	q := make([]utils.Point, 1)
	q[0] = start

	var area int
	var peri int

	for len(q) > 0 {
		e := q[0]
		q = q[1:]

		if _, ok := s.been[e]; ok {
			continue
		}
		s.been[e] = true

		area++
		peri += 4
		for _, l := range utils.GetDirectAdjacend(e.X, e.Y, len(s.data), len(s.data[0])) {
			if s.data[l.Y][l.X] == byte(letter) {
				q = append(q, utils.Point{X: l.X, Y: l.Y})
				peri--
			}
		}
	}

	if area == 0 {
		return 0
	}

	slog.Default().Debug("Part One",
		"letter", string(letter),
		"area", area,
		"perimeter", peri,
		"total", area*peri,
	)

	return area * peri
}
