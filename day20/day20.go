package day20

import (
	"log/slog"
	"strings"
	"time"

	"github.com/jonavdm/aoc-2024/utils"
)

const REQUIRED_BEAT int = 100

type Solver struct {
	Room   utils.Maze[bool]
	Height int
	Width  int
	Start  utils.Point
	End    utils.Point

	Required int
}

func Run(data []string) [2]interface{} {
	s := Parse(data)

	now := time.Now()
	dist := utils.BFS(s.Start, s.End, s)

	slog.Debug(
		"Initial distance found",
		"distance", len(dist)-1,
		"duration", time.Since(now),
		"ETA", time.Since(now)*time.Duration(len(dist)),
	)

	var count int
	differences := s.FindCheats(dist)
	for k, v := range differences {
		if k >= s.Required {
			count += v
		}
	}

	return [2]interface{}{
		count,
		0,
	}
}

func Parse(data []string) Solver {
	s := Solver{
		Room:     utils.InitSimpleMaze(data, '#'),
		Height:   len(data),
		Width:    len(data[0]),
		Required: REQUIRED_BEAT,
	}

	for y, row := range data {
		if i := strings.IndexRune(row, 'S'); i != -1 {
			s.Start = utils.Point{X: i, Y: y}
		} else if i := strings.IndexRune(row, 'E'); i != -1 {
			s.End = utils.Point{X: i, Y: y}
		}
	}

	return s
}

func (s Solver) Next(p utils.Point) []utils.Point {
	next := make([]utils.Point, 0)

	for _, n := range utils.GetDirectAdjacend(p.X, p.Y, s.Height, s.Width) {
		if s.Room.Get(n.X, n.Y) {
			continue
		}

		next = append(next, n)
	}

	return next
}

func (s *Solver) FindCheats(init []utils.Point) map[int]int {
	differences := make(map[int]int)
	done := make(map[utils.Point]bool)

	for _, i := range init {
		for _, search := range utils.GetDirectAdjacend(i.X, i.Y, s.Height, s.Width) {
			if _, ok := done[search]; ok {
				continue
			}
			if !s.Room.Get(search.X, search.Y) {
				continue
			}

			s.Room.Set(search.X, search.Y, false)
			done[search] = true

			dis := utils.BFS(s.Start, s.End, s)
			diff := len(init) - len(dis)
			if diff > 0 {
				differences[diff]++
			}

			s.Room.Set(search.X, search.Y, true)
		}
	}

	return differences
}
