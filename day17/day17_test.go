package day17_test

import (
	"testing"

	"github.com/jonavdm/aoc-2024/day17"
	_ "github.com/jonavdm/aoc-2024/testing"
	"github.com/stretchr/testify/assert"
)

// func TestRun(t *testing.T) {
// 	assert.Equal(t, [2]interface{}{0, 0}, day17.Run("day17"))
// }

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day17.Run("day17")
	}
}

func TestMachine_Simulate(t *testing.T) {
	testCases := []struct {
		name   string
		in     day17.Machine
		out    day17.Machine
		outStr string
	}{
		{
			name: "example case",
			in: day17.Machine{
				A:    729,
				B:    0,
				C:    0,
				Inst: []int{0, 1, 5, 4, 3, 0},
			},
			out: day17.Machine{
				A:    0,
				B:    0,
				C:    0,
				Inst: []int{0, 1, 5, 4, 3, 0},
			},
			outStr: "4,6,3,5,6,3,5,2,1,0",
		},
		{
			name: "Small example 1",
			in: day17.Machine{
				A:    0,
				B:    0,
				C:    9,
				Inst: []int{2, 6},
			},
			out: day17.Machine{
				A:    0,
				B:    1,
				C:    9,
				Inst: []int{2, 6},
			},
			outStr: "",
		},
		{
			name: "Small example 2",
			in: day17.Machine{
				A:    10,
				B:    0,
				C:    0,
				Inst: []int{5, 0, 5, 1, 5, 4},
			},
			out: day17.Machine{
				A:    10,
				B:    0,
				C:    0,
				Inst: []int{5, 0, 5, 1, 5, 4},
			},
			outStr: "0,1,2",
		},
		{
			name: "Small example 3",
			in: day17.Machine{
				A:    2024,
				B:    0,
				C:    0,
				Inst: []int{0, 1, 5, 4, 3, 0},
			},
			out: day17.Machine{
				A:    0,
				B:    0,
				C:    0,
				Inst: []int{0, 1, 5, 4, 3, 0},
			},
			outStr: "4,2,5,6,7,7,7,7,3,1,0",
		},
		{
			name: "Small example 4",
			in: day17.Machine{
				A:    0,
				B:    29,
				C:    0,
				Inst: []int{1, 7},
			},
			out: day17.Machine{
				A:    0,
				B:    26,
				C:    0,
				Inst: []int{1, 7},
			},
			outStr: "",
		},
		{
			name: "Small example 5",
			in: day17.Machine{
				A:    0,
				B:    2024,
				C:    43690,
				Inst: []int{4, 0},
			},
			out: day17.Machine{
				A:    0,
				B:    44354,
				C:    43690,
				Inst: []int{4, 0},
			},
			outStr: "",
		},
		{
			name: "Opcode 0",
			in: day17.Machine{
				A:    12,
				B:    0,
				C:    0,
				Inst: []int{0, 2},
			},
			out: day17.Machine{
				A:    3,
				B:    0,
				C:    0,
				Inst: []int{0, 2},
			},
			outStr: "",
		},
		{
			name: "Opcode 1",
			in: day17.Machine{
				A:    0,
				B:    12,
				C:    0,
				Inst: []int{1, 5},
			},
			out: day17.Machine{
				A:    0,
				B:    9,
				C:    0,
				Inst: []int{1, 5},
			},
			outStr: "",
		},
		{
			name: "Opcode 2",
			in: day17.Machine{
				A:    12,
				B:    0,
				C:    0,
				Inst: []int{2, 4},
			},
			out: day17.Machine{
				A:    12,
				B:    4,
				C:    0,
				Inst: []int{2, 4},
			},
			outStr: "",
		},
		{
			// Note that this instruction assumes that opcode 5 works
			name: "Opcode 3 (zero value)",
			in: day17.Machine{
				A:    0,
				B:    0,
				C:    0,
				Inst: []int{3, 5, 5, 1},
			},
			out: day17.Machine{
				A:    0,
				B:    0,
				C:    0,
				Inst: []int{3, 5, 5, 1},
			},
			outStr: "1",
		},
		{
			// Note that this instruction assumes that opcode 5 works
			name: "Opcode 3 (non-zero value)",
			in: day17.Machine{
				A:    1,
				B:    0,
				C:    0,
				Inst: []int{3, 5, 5, 1},
			},
			out: day17.Machine{
				A:    1,
				B:    0,
				C:    0,
				Inst: []int{3, 5, 5, 1},
			},
			outStr: "",
		},
		{
			name: "Opcode 4",
			in: day17.Machine{
				A:    0,
				B:    12,
				C:    5,
				Inst: []int{4, 123},
			},
			out: day17.Machine{
				A:    0,
				B:    9,
				C:    5,
				Inst: []int{4, 123},
			},
			outStr: "",
		},
		{
			name: "Opcode 5",
			in: day17.Machine{
				A:    12,
				B:    0,
				C:    0,
				Inst: []int{5, 1, 5, 4},
			},
			out: day17.Machine{
				A:    12,
				B:    0,
				C:    0,
				Inst: []int{5, 1, 5, 4},
			},
			outStr: "1,4",
		},
		{
			name: "Opcode 6",
			in: day17.Machine{
				A:    12,
				B:    0,
				C:    0,
				Inst: []int{6, 2},
			},
			out: day17.Machine{
				A:    12,
				B:    3,
				C:    0,
				Inst: []int{6, 2},
			},
			outStr: "",
		},
		{
			name: "Opcode 7",
			in: day17.Machine{
				A:    12,
				B:    0,
				C:    0,
				Inst: []int{7, 2},
			},
			out: day17.Machine{
				A:    12,
				B:    0,
				C:    3,
				Inst: []int{7, 2},
			},
			outStr: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out := tc.in.Simulate()
			assert.Equal(t, tc.out, tc.in)
			assert.Equal(t, tc.outStr, out)
		})
	}
}
