package utils_test

import (
	"testing"

	"github.com/jonavdm/aoc-2024/utils"
	"github.com/stretchr/testify/assert"
)

func TestNumberLength(t *testing.T) {
	testCases := []struct {
		in   int
		out  int
		desc string
	}{
		{
			123,
			3,
			"123 = 3",
		},
		{
			1234567,
			7,
			"1234567 = 7",
		},
		{
			3,
			1,
			"1 = 3",
		},
		{
			1234567890123456789,
			19,
			"1234567890123456789 = 19",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			assert.Equal(t, tc.out, utils.NumberLength(tc.in))
		})
	}
}
