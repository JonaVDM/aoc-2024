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
		if ListContains(seen, s.Orders[item]) {
			return 0
		}
		seen[i] = item
	}

	num, _ := strconv.Atoi(spl[len(spl)/2])
	return num
}

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)

	solver := Solver{
		Orders: make(map[string][]string),
	}

	s := false
	var total int

	for _, row := range data {
		if row == "" {
			s = true
			continue
		}

		if !s {
			solver.ParseOrder(row)
			continue
		}

		total += solver.CheckPrint(row)
	}

	return [2]interface{}{
		total,
		0,
	}
}

func ListContains(list []string, items []string) bool {
	for _, item := range items {
		if slices.Contains(list, item) {
			return true
		}
	}

	return false
}
