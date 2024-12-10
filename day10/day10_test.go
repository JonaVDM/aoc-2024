package day10_test

import (
	"testing"

	"github.com/jonavdm/aoc-2024/day10"
	_ "github.com/jonavdm/aoc-2024/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]interface{}{674, 1372}, day10.Run("day10"))
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day10.Run("day10")
	}
}
