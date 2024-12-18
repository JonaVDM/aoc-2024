package day17_test

import (
	"testing"

	"github.com/jonavdm/aoc-2024/day17"
	_ "github.com/jonavdm/aoc-2024/testing"
	"github.com/stretchr/testify/assert"
)

func TestMachine_Simulate(t *testing.T) {
	testCases := []struct {
		name         string
		a1, b1, c1   int
		a2, b2, c2   int
		instructions []int
		out          string
	}{
		{
			"example case",
			729, 0, 0,
			0, 0, 0,
			[]int{0, 1, 5, 4, 3, 0},
			"4,6,3,5,6,3,5,2,1,0",
		},
		{
			"Small example 1",
			0, 0, 9,
			0, 1, 9,
			[]int{2, 6},
			"",
		},
		{
			"Small example 2",
			10, 0, 0,
			10, 0, 0,
			[]int{5, 0, 5, 1, 5, 4},
			"0,1,2",
		},
		{
			"Small example 3",
			2024, 0, 0,
			0, 0, 0,
			[]int{0, 1, 5, 4, 3, 0},
			"4,2,5,6,7,7,7,7,3,1,0",
		},
		{
			"Small example 4",
			0, 29, 0,
			0, 26, 0,
			[]int{1, 7},
			"",
		},
		{
			"Small example 5",
			0, 2024, 43690,
			0, 44354, 43690,
			[]int{4, 0},
			"",
		},
		{
			"Opcode 0",
			12, 0, 0,
			3, 0, 0,
			[]int{0, 2},
			"",
		},
		{
			"Opcode 1",
			0, 12, 0,
			0, 9, 0,
			[]int{1, 5},
			"",
		},
		{
			"Opcode 2",
			12, 0, 0,
			12, 4, 0,
			[]int{2, 4},
			"",
		},
		{
			// Note that this instruction assumes that opcode 5 works
			"Opcode 3 (zero value)",
			0, 0, 0,
			0, 0, 0,
			[]int{3, 5, 5, 1},
			"1",
		},
		{
			// Note that this instruction assumes that opcode 5 works
			"Opcode 3 (non-zero value)",
			1, 0, 0,
			1, 0, 0,
			[]int{3, 5, 5, 1},
			"",
		},
		{
			"Opcode 4",
			0, 12, 5,
			0, 9, 5,
			[]int{4, 123},
			"",
		},
		{
			"Opcode 5",
			12, 0, 0,
			12, 0, 0,
			[]int{5, 1, 5, 4},
			"1,4",
		},
		{
			"Opcode 6",
			12, 0, 0,
			12, 3, 0,
			[]int{6, 2},
			"",
		},
		{
			"Opcode 6",
			12, 0, 0,
			12, 0, 3,
			[]int{7, 2},
			"",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			in := day17.Machine{
				tc.a1, tc.b1, tc.c1,
				tc.instructions,
			}
			out := day17.Machine{
				tc.a2, tc.b2, tc.c2,
				tc.instructions,
			}

			output := in.Simulate()
			assert.Equal(t, out, in)
			assert.Equal(t, tc.out, output)
		})
	}
}
