package day16

import (
	"slices"

	"github.com/jonavdm/aoc-2024/utils"
)

const (
	NORTH = iota
	EAST
	SOUTH
	WEST
)

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)

	s := parse(data)

	return [2]interface{}{
		s.findShortest(),
		0,
	}
}

type Solver struct {
	maze  map[int]map[int]bool
	start utils.Point
	end   utils.Point
}

func parse(data []string) Solver {
	s := Solver{
		maze: make(map[int]map[int]bool),
	}

	for y, r := range data {
		for x, c := range r {
			if _, ok := s.maze[x]; !ok {
				s.maze[x] = make(map[int]bool)
			}

			if c == '#' {
				s.maze[x][y] = true
				continue
			}

			if c == 'S' {
				s.start = utils.Point{X: x, Y: y}
			}

			if c == 'E' {
				s.end = utils.Point{X: x, Y: y}
			}
		}
	}

	return s
}

type item struct {
	value int
	p     utils.Point
	dir   int
}

func getEdges(p utils.Point, dir int) []item {
	n := item{1, utils.Point{X: p.X, Y: p.Y - 1}, NORTH}
	e := item{1, utils.Point{X: p.X + 1, Y: p.Y}, EAST}
	s := item{1, utils.Point{X: p.X, Y: p.Y + 1}, SOUTH}
	w := item{1, utils.Point{X: p.X - 1, Y: p.Y}, WEST}

	if dir == NORTH {
		e.value += 1000
		w.value += 1000
		return []item{
			n, e, w,
		}
	}

	if dir == SOUTH {
		e.value += 1000
		w.value += 1000
		return []item{
			s, e, w,
		}
	}

	if dir == EAST {
		n.value += 1000
		s.value += 1000
		return []item{
			e, n, s,
		}
	}

	if dir == WEST {
		n.value += 1000
		s.value += 1000
		return []item{
			w, n, s,
		}
	}

	return []item{}
}

func (s *Solver) findShortest() int {
	been := make(map[utils.Point]int)
	q := make([]item, 1)
	q[0] = item{0, utils.Point{X: s.start.X, Y: s.start.Y}, EAST}

	for len(q) > 0 {
		i := q[0]
		q = q[1:]

		if v, ok := been[i.p]; ok && v < i.value {
			continue
		}

		been[i.p] = i.value

		if i.p == s.end {
			return i.value
		}

		for _, e := range getEdges(i.p, i.dir) {
			if s.maze[e.p.X][e.p.Y] {
				continue
			}

			e.value += i.value
			q = append(q, e)
		}

		slices.SortFunc(q, func(a, b item) int {
			return a.value - b.value
		})
	}

	return -1
}
