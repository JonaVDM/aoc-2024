package day13

import (
	"strconv"
	"strings"

	"github.com/jonavdm/aoc-2024/utils"
)

func Run(data []string) [2]interface{} {
	machines := Parse(data)

	var a int
	var b int
	for _, m := range machines {
		a += GetTickets(m)
		m.Prize.X += 10000000000000
		m.Prize.Y += 10000000000000
		b += GetTickets(m)
	}

	return [2]interface{}{
		a,
		b,
	}
}

type Machine struct {
	A, B, Prize utils.Point
}

func Parse(data []string) []Machine {
	machines := make([]Machine, 0)
	m := Machine{}

	for _, l := range data {
		if l == "" {
			machines = append(machines, m)
			m = Machine{}
		}

		if strings.Contains(l, "Button A:") {
			m.A = GetValues(l)
		}

		if strings.Contains(l, "Button B:") {
			m.B = GetValues(l)
		}

		if strings.Contains(l, "Prize:") {
			m.Prize = GetValues(l)
		}
	}

	machines = append(machines, m)

	return machines
}

func GetValues(s string) utils.Point {
	parts := strings.Split(s, ": ")
	parts = strings.Split(parts[1], ", ")

	xs := strings.Split(parts[0], "+")
	if len(xs) == 1 {
		xs = strings.Split(parts[0], "=")
	}
	x, _ := strconv.Atoi(xs[1])

	ys := strings.Split(parts[1], "+")
	if len(ys) == 1 {
		ys = strings.Split(parts[1], "=")
	}
	y, _ := strconv.Atoi(ys[1])

	return utils.Point{X: x, Y: y}
}

func GetTickets(m Machine) int {
	a := (m.Prize.Y*m.B.X - m.B.Y*m.Prize.X) / (m.A.Y*m.B.X - m.B.Y*m.A.X)
	b := (m.Prize.X - m.A.X*a) / m.B.X

	if m.A.X*a+m.B.X*b == m.Prize.X && m.A.Y*a+m.B.Y*b == m.Prize.Y {
		return a*3 + b
	} else {
		return 0
	}
}
