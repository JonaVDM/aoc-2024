package day19_test

import (
	"strings"
	"testing"

	"github.com/jonavdm/aoc-2024/day19"
	_ "github.com/jonavdm/aoc-2024/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	file := strings.Split(`r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`, "\n")
	assert.Equal(t, [2]interface{}{6, 0}, day19.Run(file))
}

func TestFindNext(t *testing.T) {
	tests := []struct {
		towels  []string
		pattern string
		out     []int
		label   string
	}{
		{[]string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"}, "brwrr", []int{2, 7}, ""},
		{[]string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"}, "wrr", []int{1}, ""},
		{[]string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"}, "r", []int{0}, ""},
		{[]string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"}, "ubwu", []int{}, ""},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.out, day19.FindNext(tc.towels, tc.pattern), tc.label)
	}
}

func TestFindMatch(t *testing.T) {
	towels := []string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"}
	tests := []struct {
		towels  []string
		pattern string
		out     [][]int
	}{
		{towels, "brwrr", [][]int{{2, 0, 1, 0}, {7, 1, 0}}},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.out, day19.FindMatch(tc.towels, tc.pattern))
	}
}
