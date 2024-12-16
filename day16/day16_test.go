package day16_test

import (
	"testing"

	"github.com/jonavdm/aoc-2024/day16"
	_ "github.com/jonavdm/aoc-2024/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]interface{}{106512, 0}, day16.Run("day16"))
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day16.Run("day16")
	}
}

