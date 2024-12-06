package day06_test

import (
	"testing"

	"github.com/jonavdm/aoc-2024/day06"
	_ "github.com/jonavdm/aoc-2024/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]interface{}{5162, 0}, day06.Run("day06"))
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day06.Run("day06")
	}
}
