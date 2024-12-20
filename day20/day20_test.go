package day20_test

import (
	"strings"
	"testing"

	"github.com/jonavdm/aoc-2024/day20"
	_ "github.com/jonavdm/aoc-2024/testing"
	"github.com/jonavdm/aoc-2024/utils"
	"github.com/stretchr/testify/assert"
)

func TestSolver_FindCheats(t *testing.T) {
	file := strings.Split(`###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`, "\n")

	expected := make(map[int]int)
	expected[2] = 14
	expected[4] = 14
	expected[6] = 2
	expected[8] = 4
	expected[10] = 2
	expected[12] = 3
	expected[20] = 1
	expected[36] = 1
	expected[38] = 1
	expected[40] = 1
	expected[64] = 1

	s := day20.Parse(file)
	init := utils.BFS(s.Start, s.End, s)
	assert.Equal(t, expected, s.FindCheats(init))
}
