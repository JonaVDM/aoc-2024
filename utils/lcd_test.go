package utils_test

import (
	"testing"

	"github.com/jonavdm/aoc-2024/utils"
	"github.com/stretchr/testify/assert"
)

func TestLeastCommonDenominator(t *testing.T) {
	tests := []struct {
		in    []int
		out   int
		label string
	}{
		{[]int{23, 25}, 575, "23, 25"},
		{[]int{100, 200, 300}, 600, "100, 200, 300"},
		{[]int{129, 3129, 1230}, 55164270, "129, 3129, 123"},
	}

	for _, test := range tests {
		assert.Equal(t, test.out, utils.LeastCommonDenominator(test.in), test.label)
	}
}
