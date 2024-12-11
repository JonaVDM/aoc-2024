package day11_test

import (
	"testing"

	"github.com/jonavdm/aoc-2024/day11"
	_ "github.com/jonavdm/aoc-2024/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]interface{}{184927, 220357186726677}, day11.Run("day11"))
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day11.Run("day11")
	}
}
