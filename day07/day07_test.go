package day07_test

import (
	"testing"

	"github.com/jonavdm/aoc-2024/day07"
	_ "github.com/jonavdm/aoc-2024/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]interface{}{1298103531759, 140575048428831}, day07.Run("day07"))
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day07.Run("day07")
	}
}
