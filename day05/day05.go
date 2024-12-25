package day05

import (
	"slices"
	"strconv"
	"strings"

	"github.com/jonavdm/aoc-2024/utils"
)

type Solver struct {
	Orders map[string][]string
}

func Run(data []string) [2]interface{} {
	solver := Solver{
		Orders: make(map[string][]string),
	}

	s := false
	var total int
	var fixed int

	for _, row := range data {
		if row == "" {
			s = true
			continue
		}

		if !s {
			solver.ParseOrder(row)
			continue
		}

		if score := solver.CheckPrint(row); score != 0 {
			total += score
		} else {
			s := solver.FixPrint(row)
			fixed += solver.CheckPrint(s)
		}
	}

	return [2]interface{}{
		total,
		fixed,
	}
}

func (s *Solver) ParseOrder(order string) {
	parts := strings.Split(order, "|")
	if _, ok := s.Orders[parts[0]]; !ok {
		s.Orders[parts[0]] = make([]string, 1)
		s.Orders[parts[0]][0] = parts[1]
	} else {
		s.Orders[parts[0]] = append(s.Orders[parts[0]], parts[1])
	}
}

func (s *Solver) CheckPrint(order string) int {
	spl := strings.Split(order, ",")
	seen := make([]string, len(spl))

	for i, item := range spl {
		if utils.ListContains(seen, s.Orders[item]) {
			return 0
		}
		seen[i] = item
	}

	num, _ := strconv.Atoi(spl[len(spl)/2])
	return num
}

func (s *Solver) FixPrint(order string) string {
	spl := strings.Split(order, ",")
	slices.SortFunc(spl, func(a, b string) int {
		if slices.Contains(s.Orders[a], b) {
			return -1
		}
		return 1
	})

	return strings.Join(spl, ",")
}
