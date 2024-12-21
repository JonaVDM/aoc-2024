package day21_test

import (
	"fmt"
	"testing"

	"github.com/jonavdm/aoc-2024/day21"
	_ "github.com/jonavdm/aoc-2024/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	file := []string{
		"029A",
		"980A",
		"179A",
		"456A",
		"379A",
	}
	assert.Equal(t, [2]interface{}{126384, 0}, day21.Run(file))
}

func TestPathTo(t *testing.T) {
	tests := []struct {
		start rune
		end   rune
		out   string
		count int
	}{
		{'A', '7', "^<", 3},
		{'A', '8', "<^", 2},
		{'7', '2', "v>", 1},
		{'5', '6', ">", 0},
		{'4', '6', ">", 1},
		{'9', '7', "<", 1},
		{'2', '8', "^", 1},
		{'4', '1', "v", 0},
		{'2', '0', "v", 0},
		{'9', '6', "v", 0},
		{'A', '3', "^", 0},
		{'0', '5', "^", 1},
		{'7', '0', ">v", 2},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprint(string(tc.start), string(tc.end)), func(t *testing.T) {
			o, c := day21.PathTo(tc.start, tc.end)
			assert.Equal(t, tc.count, c)
			assert.Equal(t, tc.out, o)
		})
	}
}

// func TestLastBot(t *testing.T) {
// 	tests := []struct {
// 		in  string
// 		out int
// 	}{
// 		{"029A", len("<A^A>^^AvvvA")},
// 	}

// 	for _, tc := range tests {
// 		t.Run(tc.in, func(t *testing.T) {
// 			assert.Equal(t, tc.out, len(day21.LastBot(tc.in)))
// 		})
// 	}
// }

func TestMiddleBots(t *testing.T) {
	test := []struct {
		seq    []string
		amount int
		out    int
	}{
		{[]string{"<A", "^A", ">^A", "vA"}, 2, 65},

		// https://www.reddit.com/r/adventofcode/comments/1hj7f89/2024_day_21_part_1_found_a_rule_to_make_it_work/
		{[]string{"<^A"}, 1, 9},
		{[]string{"<^A"}, 2, 21},
		{[]string{"^<A"}, 1, 9},
		{[]string{"^<A"}, 2, 25},
	}

	for _, tc := range test {
		t.Run(fmt.Sprint(tc.seq), func(t *testing.T) {
			assert.Equal(t, tc.out, day21.MiddleBots(tc.seq, tc.amount))
		})
	}
}
