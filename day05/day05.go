package day05

import (
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

// GetValue returns the center value of array `order`
func GetValue(order []string) int {
	num, err := strconv.Atoi(order[len(order)/2])
	if err != nil {
		panic(err)
	}

	return num
}
