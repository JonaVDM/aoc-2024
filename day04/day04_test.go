package day04_test

import (
	"testing"

	"github.com/jonavdm/aoc-2024/day04"
	_ "github.com/jonavdm/aoc-2024/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]interface{}{2504, 1923}, day04.Run("day04"))
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day04.Run("day04")
	}
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

func TestFindXmas(t *testing.T) {
	testCases := []struct {
		name  string
		input []string
		x     int
		y     int
		out   int
	}{
		{
			name:  "single row horizontal",
			input: []string{"AAAAXMASAA"},
			x:     4,
			y:     0,
			out:   2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			solver := day04.Solver{
				Data: tc.input,
			}
			out := solver.FindM(tc.x, tc.y)

			assert.Equal(t, 1, out)
		})
	}
}
