package day06_test

import (
	"testing"

	"github.com/jonavdm/aoc-2024/day06"
	_ "github.com/jonavdm/aoc-2024/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	test := []string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"........#.",
		"#.........",
		"......#...",
	}
	assert.Equal(t, [2]interface{}{41, 0}, day06.Run(test))
}
