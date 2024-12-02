package day02_test

import (
	"testing"

	"github.com/jonavdm/aoc-2024/day02"
	_ "github.com/jonavdm/aoc-2024/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]interface{}{686, 716}, day02.Run("day02"))
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day02.Run("day02")
	}
}
