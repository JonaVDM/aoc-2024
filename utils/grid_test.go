package utils_test

import (
	"testing"

	"github.com/jonavdm/aoc-2024/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetAdjacend(t *testing.T) {
	type Test struct {
		x, y          int
		height, width int
		out           []utils.Point
	}

	tests := []Test{
		{
			x:      0,
			y:      0,
			height: 10,
			width:  10,
			out: []utils.Point{
				{1, 0},
				{0, 1},
				{1, 1},
			},
		},
		{
			x:      2,
			y:      2,
			height: 10,
			width:  10,
			out: []utils.Point{
				{1, 1},
				{2, 1},
				{3, 1},
				{1, 2},
				{3, 2},
				{1, 3},
				{2, 3},
				{3, 3},
			},
		},
		{
			x:      9,
			y:      9,
			height: 10,
			width:  10,
			out: []utils.Point{
				{8, 8},
				{9, 8},
				{8, 9},
			},
		},
		{
			x:      0,
			y:      0,
			height: 1,
			width:  1,
			out:    []utils.Point{},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.out, utils.GetAdjacend(test.x, test.y, test.height, test.width))
	}
}

func BenchmarkGetAdjacend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utils.GetAdjacend(5, 5, 10, 10)
	}
}
