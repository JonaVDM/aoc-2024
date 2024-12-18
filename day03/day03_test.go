package day03_test

import (
	"testing"

	"github.com/jonavdm/aoc-2024/day03"
	_ "github.com/jonavdm/aoc-2024/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	data := []string{
		"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
	}
	assert.Equal(t, [2]interface{}{161, 48}, day03.Run(data))
}
