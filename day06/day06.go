package day06

import (
	"github.com/jonavdm/aoc-2024/utils"
)

type Solver struct {
	InitialPosition utils.Point
	Pos             utils.Point
	Dir             utils.Point

	Height, Width int

	Field  map[int]map[int]bool
	Been   map[int]map[int]bool
	Count  int
	Enable bool
}

func (s *Solver) Move() bool {
	if row, ok := s.Field[s.Pos.X+s.Dir.X]; ok && row[s.Pos.Y+s.Dir.Y] {
		s.Rotate()
	}

	x := s.Pos.X + s.Dir.X
	y := s.Pos.Y + s.Dir.Y

	if x < 0 || x >= s.Width || y < 0 || y >= s.Height {
		return false
	}

	if _, ok := s.Been[x]; s.Enable && (!ok || !s.Been[x][y]) {
		s.Count++
		if !ok {
			s.Been[x] = make(map[int]bool)
		}
		s.Been[x][y] = true
	}

	s.Pos.X = x
	s.Pos.Y = y

	return true
}

func (s *Solver) Rotate() {
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

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)

	s1 := Solver{
		Dir:    utils.Point{X: 0, Y: -1},
		Field:  make(map[int]map[int]bool),
		Been:   make(map[int]map[int]bool),
		Height: len(data),
		Width:  len(data[0]),
		Enable: true,
	}

	s2 := Solver{
		Dir:    utils.Point{X: 0, Y: -1},
		Field:  make(map[int]map[int]bool),
		Height: len(data),
		Width:  len(data[0]),
	}

	for y, row := range data {
		for x, l := range row {
			if l == '.' {
				continue
			}

			if l == '^' {
				s1.Pos.X = x
				s1.Pos.Y = y + 1

				s2.InitialPosition.X = x
				s2.InitialPosition.Y = y
				continue
			}

			if l != '#' {
				panic("dunno")
			}

			if _, ok := s1.Field[x]; !ok {
				s1.Field[x] = make(map[int]bool)
				s2.Field[x] = make(map[int]bool)
			}
			s1.Field[x][y] = true
			s2.Field[x][y] = true
		}
	}

	for {
		if !s1.Move() {
			break
		}
	}

	// for x := 0; x < s2.Width; x++ {
	// 	for y := 0; y < s2.Height; y++ {
	// 		s := s2
	// 		s.Pos.X, s.Pos.Y = s.InitialPosition.X, s.InitialPosition.Y
	// 		s.Field = utils.CopyMap(s1.Field)
	// 		// if !s2.Move() {
	// 		// 	break
	// 		// }
	// 	}
	// }

	return [2]interface{}{
		s1.Count,
		0,
	}
}
