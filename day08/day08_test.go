package day08_test

import (
	"testing"

	"github.com/jonavdm/aoc-2024/day08"
	_ "github.com/jonavdm/aoc-2024/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]interface{}{413, 1417}, day08.Run("day08"))
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day08.Run("day08")
	}
}

