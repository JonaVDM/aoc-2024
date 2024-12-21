package day21

import (
	"strconv"

	"github.com/jonavdm/aoc-2024/utils"
)

func Run(data []string) [2]interface{} {
	var totalA int
	var totalB int

	cache := make(map[string]map[int]int)

	for _, row := range data {
		val, err := strconv.Atoi(row[0:3])
		if err != nil {
			panic("Invalid input for day 21")
		}

		seq, count := LastBot(row)
		var total int
		var b int

		for _, s := range seq {
			total += MiddleBots(s, 2, cache)
			b += MiddleBots(s, 25, cache)
		}

		totalA += val * (total + count)
		totalB += val * (b + count)
	}

	return [2]interface{}{
		totalA,
		totalB,
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

	// The order in which to move first
	// 1. move left
	// 2. move up / down
	// 3. move right
	// If one makes it so that you come to the corner, skip it for now
	var out string

	// Going left + vertical
	if x > 0 && !(s.Y == 3 && e.X == 0) {
		out = "<"
		if y > 0 {
			out += "^"
		} else if y < 0 {
			out += "v"
		}
	}

	// Going up + horizontal
	if out == "" && y > 0 {
		out = "^"
		if x > 0 {
			out += "<"
		} else if x < 0 {
			out += ">"
		}
	}

	// Going down + horizontal
	if out == "" && y < 0 && !(s.X == 0 && e.Y == 3) {
		out = "v"
		if x > 0 {
			out += "<"
		} else if x < 0 {
			out += ">"
		}
	}

	// Going right + vertical
	if x < 0 && out == "" {
		out = ">"
		if y > 0 {
			out += "^"
		} else if y < 0 {
			out += "v"
		}
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

func MiddleBots(seq string, amount int, cache map[string]map[int]int) int {
	if amount == 0 {
		return len(seq)
	}

	if _, ok := cache[seq]; ok {
		if val, ok := cache[seq][amount]; ok {
			return val
		}
	} else if cache != nil {
		cache[seq] = make(map[int]int)
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

	var total int

	newSeq, ok := pattern[seq]
	if !ok {
		panic("unkown pattern " + seq)
	}

	for _, n := range newSeq {
		total += MiddleBots(n, amount-1, cache)
	}

	if cache != nil {
		cache[seq][amount] = extra[seq] + total
	}

	return extra[seq] + total
}
