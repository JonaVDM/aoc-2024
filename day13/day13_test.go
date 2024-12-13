package day13_test

import (
	"testing"

	"github.com/jonavdm/aoc-2024/day13"
	_ "github.com/jonavdm/aoc-2024/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]interface{}{36571, 85527711500010}, day13.Run("day13"))
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day13.Run("day13")
	}
}
