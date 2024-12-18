package day02_test

import (
	"testing"

	"github.com/jonavdm/aoc-2024/day02"
	_ "github.com/jonavdm/aoc-2024/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	input := []string{
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	}
	assert.Equal(t, [2]interface{}{2, 4}, day02.Run(input))
}

func TestValidLine(t *testing.T) {
	testCases := []struct {
		in  []int
		out bool
	}{
		{
			[]int{7, 6, 4, 2, 1},
			true,
		},
		{
			[]int{1, 2, 7, 8, 9},
			false,
		},
		{
			[]int{9, 7, 6, 2, 1},
			false,
		},
		{
			[]int{1, 3, 2, 4, 5},
			false,
		},
		{
			[]int{8, 6, 4, 4, 1},
			false,
		},
		{
			[]int{1, 3, 6, 7, 9},
			true,
		},
	}
	for _, tC := range testCases {
		assert.Equal(t, tC.out, day02.ValidLine(tC.in), tC.in)
	}
}
