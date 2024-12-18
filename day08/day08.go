package day08

import (
	"github.com/jonavdm/aoc-2024/utils"
)

type Solver struct {
	Points map[rune][]utils.Point
	Anti   map[int]map[int]bool

	Height, Width int
}

func Run(data []string) [2]interface{} {
	s := Solver{
		Points: make(map[rune][]utils.Point),
		Height: len(data),
		Width:  len(data[0]),
	}

	for y, row := range data {
		for x, let := range row {
			if let == '.' {
				continue
			}

			if _, ok := s.Points[let]; !ok {
				s.Points[let] = make([]utils.Point, 0)
			}

			s.Points[let] = append(s.Points[let], utils.Point{X: x, Y: y})
		}
	}

	return [2]interface{}{
		s.Run(false),
		s.Run(true),
	}
}

func (s *Solver) Run(two bool) int {
	var count int
	anti := make(map[int]map[int]bool)

	for _, pp := range s.Points {
		for i := 0; i < len(pp); i++ {
			if two {
				if _, ok := anti[pp[i].X]; !ok {
					anti[pp[i].X] = make(map[int]bool)
				}

				if !anti[pp[i].X][pp[i].Y] {
					count++
					anti[pp[i].X][pp[i].Y] = true
				}
			}

			for j := i + 1; j < len(pp); j++ {
				l, r := pp[i], pp[j]

				xd := r.X - l.X
				yd := r.Y - l.Y

				x1 := l.X
				y1 := l.Y
				x2 := r.X
				y2 := r.Y

				for {
					x1 -= xd
					y1 -= yd
					x2 += xd
					y2 += yd

					if _, ok := anti[x1]; !ok {
						anti[x1] = make(map[int]bool)
					}
					if _, ok := anti[x2]; !ok {
						anti[x2] = make(map[int]bool)
					}

					if s.InRange(x1, y1) && !anti[x1][y1] {
						count++
						anti[x1][y1] = true
					}
					if s.InRange(x2, y2) && !anti[x2][y2] {
						count++
						anti[x2][y2] = true
					}

					if !two || (!s.InRange(x1, y1) && !s.InRange(x2, y2)) {
						break
					}
				}
			}
		}
	}
	return count
}

func (s *Solver) InRange(x, y int) bool {
	return utils.InRange(x, y, s.Width, s.Height)
}
