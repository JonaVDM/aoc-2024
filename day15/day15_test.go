package day15_test

import (
	"testing"

	"github.com/jonavdm/aoc-2024/day15"
	_ "github.com/jonavdm/aoc-2024/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]interface{}{1485257, 0}, day15.Run("day15"))
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day15.Run("day15")
	}
}

