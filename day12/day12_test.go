package day12_test

import (
	"testing"

	"github.com/jonavdm/aoc-2024/day12"
	_ "github.com/jonavdm/aoc-2024/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]interface{}{1363484, 0}, day12.Run("day12"))
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day12.Run("day12")
	}
}

