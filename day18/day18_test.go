package day18_test

import (
	"testing"

	"github.com/jonavdm/aoc-2024/day18"
	_ "github.com/jonavdm/aoc-2024/testing"
	"github.com/stretchr/testify/assert"
)

var Example []string = []string{
	"5,4",
	"4,2",
	"4,5",
	"3,0",
	"2,1",
	"6,3",
	"2,4",
	"1,5",
	"0,6",
	"3,3",
	"2,6",
	"5,1",
	"1,2",
	"5,5",
	"2,5",
	"6,5",
	"1,4",
	"0,4",
	"6,4",
	"1,1",
	"6,1",
	"1,0",
	"0,5",
	"1,6",
	"2,0",
}

func TestRoom_BFS(t *testing.T) {
	r := day18.Parse(Example, 12)
	r.Range = 6
	assert.Equal(t, 22, r.BFS())
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day18.Run("day18")
	}
}

