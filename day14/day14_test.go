package day14_test

import (
	"testing"

	"github.com/jonavdm/aoc-2024/day14"
	_ "github.com/jonavdm/aoc-2024/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]interface{}{211773366, 0}, day14.Run("day14"))
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day14.Run("day14")
	}
}
