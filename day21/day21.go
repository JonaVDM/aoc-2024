package day21

import (
	"strconv"

	"github.com/jonavdm/aoc-2024/utils"
)

func Run(data []string) [2]interface{} {
	var total int

	for _, row := range data {
		val, err := strconv.Atoi(row[0:3])
		if err != nil {
			panic("Invalid input for day 21")
		}

		seq, count := LastBot(row)
		am := MiddleBots(seq, 2)

		total += val * (am + count)
	}

	return [2]interface{}{
		total,
		0,
	}
}

func PathTo(start rune, end rune) (string, int) {
	locations := map[rune]utils.Point{
		'7': {X: 0, Y: 0},
		'8': {X: 1, Y: 0},
		'9': {X: 2, Y: 0},

		'4': {X: 0, Y: 1},
		'5': {X: 1, Y: 1},
		'6': {X: 2, Y: 1},

		'1': {X: 0, Y: 2},
		'2': {X: 1, Y: 2},
		'3': {X: 2, Y: 2},

		'0': {X: 1, Y: 3},
		'A': {X: 2, Y: 3},
	}

	s := locations[start]
	e := locations[end]

	x := s.X - e.X
	y := s.Y - e.Y

	var xout, yout string

	if x > 0 {
		xout += "<"
	} else if x < 0 {
		xout += ">"
	}
	if y > 0 {
		yout += "^"
	} else if y < 0 {
		yout += "v"
	}

	var out string

	corner := (s.Y == 3 && e.X == 0) || (e.Y == 3 && s.X == 0)

	// We go left, and the path doesn't go into the corner
	if x > 0 && !corner {
		out = xout + yout
	} else {
		out = yout + xout
	}

	return out, utils.AbsInt(x) + utils.AbsInt(y) - len(out)
}

func LastBot(seq string) ([]string, int) {
	start := 'A'
	out := []string{}
	var total int
	for _, c := range seq {
		path, count := PathTo(start, c)
		out = append(out, path+"A")
		start = c
		total += count
	}

	return out, total
}

func MiddleBots(seq []string, amount int) int {
	if amount == 0 {
		total := 0

		for _, s := range seq {
			total += len(s)
		}

		return total
	}
	extra := map[string]int{
		"<^A": 1,
		"^<A": 1,
		"<A":  2,
		"<vA": 1,
		"v<A": 1,
	}

	pattern := map[string][]string{
		"<^A": {"v<A", ">^A", ">A"},
		"^<A": {"<A", "v<A", ">^A"},
		"<A":  {"v<A", ">^A"},
		"<vA": {"v<A", ">A", "^>A"},
		"v<A": {"<vA", "<A", ">^A"},
		"^A":  {"<A", ">A"},
		"vA":  {"<vA", "^>A"},
		"^>A": {"<A", "v>A", "^A"},
		">^A": {"vA", "<^A", ">A"},
		">A":  {"vA", "^A"},
		"v>A": {"<vA", ">A", "^A"},
		">vA": {"vA", "<A", "^>A"},
	}

	total := 0
	newSeq := make([]string, 0)

	for _, s := range seq {
		total += extra[s]
		if pat, ok := pattern[s]; !ok {
			panic("unkown pattern " + s)
		} else {
			newSeq = append(newSeq, pat...)
		}
	}

	return total + MiddleBots(newSeq, amount-1)
}
