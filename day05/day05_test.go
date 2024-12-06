package day05_test

import (
	"testing"

	"github.com/jonavdm/aoc-2024/day05"
	_ "github.com/jonavdm/aoc-2024/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]interface{}{4959, 0}, day05.Run("day05"))
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day05.Run("day05")
	}
}