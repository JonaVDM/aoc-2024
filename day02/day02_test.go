package day02_test

import (
	"testing"

	"github.com/jonavdm/aoc-2024/day02"
	_ "github.com/jonavdm/aoc-2024/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]interface{}{686, 717}, day02.Run("day02"))
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day02.Run("day02")
	}
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
