package day25

type Solver struct {
	Keys  [][5]int
	Locks [][5]int
}

func Run(data []string) [2]interface{} {
	solver := Parse(data)

	return [2]interface{}{
		solver.CountMatches(),
		0,
	}
}

func Parse(data []string) Solver {
	solver := Solver{
		Keys:  make([][5]int, 0),
		Locks: make([][5]int, 0),
	}

	for i := 0; i < len(data); i += 8 {
		k := [5]int{}

		for x := range 5 {
			for y := range 5 {
				if data[y+i+1][x] == '#' {
					k[x]++
				}
			}
		}

		if data[i] == "#####" {
			solver.Locks = append(solver.Locks, k)
		} else {
			solver.Keys = append(solver.Keys, k)
		}
	}

	return solver
}

func (s *Solver) CountMatches() int {
	var count int

	for _, lock := range s.Locks {
		for _, key := range s.Keys {
			if KeyFits(key, lock) {
				count++
			}
		}
	}

	return count
}

func KeyFits(key, lock [5]int) bool {
	for i := range 5 {
		if key[i]+lock[i] > 5 {
			return false
		}
	}
	return true
}
