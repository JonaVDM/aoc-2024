package day04_test

import (
	"testing"

	"github.com/jonavdm/aoc-2024/day04"
	_ "github.com/jonavdm/aoc-2024/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	test := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}
	assert.Equal(t, [2]interface{}{18, 9}, day04.Run(test))
}

func TestFindMasses(t *testing.T) {
	testCases := []struct {
		desc   string
		input  []string
		output int
	}{
		{
			desc:   "simple x",
			output: 1,
			input:  []string{"M-M", "-A-", "S-S"},
		},
		{
			desc:   "double x",
			output: 2,
			input:  []string{"M-M-S-S", "-A---A-", "S-S-M-M"},
		},
		{
			desc:   "backwards x",
			output: 1,
			input:  []string{"S-S", "-A-", "M-M"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			solver := day04.Solver{
				Data: tc.input,
			}
			out := solver.FindMasses()
			assert.Equal(t, tc.output, out)
		})
	}
}
