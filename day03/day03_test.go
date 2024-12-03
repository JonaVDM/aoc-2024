package day03_test

import (
	"testing"

	"github.com/jonavdm/aoc-2024/day03"
	_ "github.com/jonavdm/aoc-2024/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]interface{}{196826776, 106780429}, day03.Run("day03"))
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day03.Run("day03")
	}
}
