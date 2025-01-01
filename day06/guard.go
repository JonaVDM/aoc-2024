package day06

import "github.com/jonavdm/aoc-2024/utils"

type Guard struct {
	Field map[utils.Point]bool

	Height int
	Width  int

	Pos utils.Point
	Dir utils.Point
}

func (s *Guard) Rotate() {
	if s.Dir.Y == -1 {
		s.Dir.X = 1
		s.Dir.Y = 0
	} else if s.Dir.X == 1 {
		s.Dir.X = 0
		s.Dir.Y = 1
	} else if s.Dir.Y == 1 {
		s.Dir.X = -1
		s.Dir.Y = 0
	} else if s.Dir.X == -1 {
		s.Dir.X = 0
		s.Dir.Y = -1
	} else {
		panic("inpossible direction")
	}
}

func (s *Guard) Move() bool {
	if wall := s.Field[utils.Point{X: s.Pos.X + s.Dir.X, Y: s.Pos.Y + s.Dir.Y}]; wall {
		s.Rotate()
	}

	x := s.Pos.X + s.Dir.X
	y := s.Pos.Y + s.Dir.Y

	if x < 0 || x >= s.Width || y < 0 || y >= s.Height {
		return false
	}

	s.Pos.X = x
	s.Pos.Y = y

	return true
}
