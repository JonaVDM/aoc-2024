package day10_test

import (
	"strings"
	"testing"

	"github.com/jonavdm/aoc-2024/day10"
	_ "github.com/jonavdm/aoc-2024/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	file := strings.Split(`89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`, "\n")
	assert.Equal(t, [2]interface{}{36, 81}, day10.Run(file))
}
