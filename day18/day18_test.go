package day18_test

import (
	"testing"

	"github.com/jonavdm/aoc-2024/day18"
	_ "github.com/jonavdm/aoc-2024/testing"
	"github.com/stretchr/testify/assert"
)

var Day18Example []string = []string{
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
	r := day18.Parse(Day18Example, 12)
	r.Range = 6
	assert.Equal(t, 22, r.BFS())
}

func TestRoom_FindLast(t *testing.T) {
	r := day18.Parse(Day18Example, 12)
	r.Range = 6
	assert.Equal(t, "6,1", r.FindLast(12))
}
