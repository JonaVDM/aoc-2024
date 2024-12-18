package day12_test

import (
	"strings"
	"testing"

	"github.com/jonavdm/aoc-2024/day12"
	_ "github.com/jonavdm/aoc-2024/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	file := strings.Split(`RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`, "\n")
	assert.Equal(t, [2]interface{}{1930, 0}, day12.Run(file))
}
