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
	assert.Equal(t, [2]interface{}{6, 16}, day19.Run(file))
}

func TestFindNext(t *testing.T) {
	towels := []string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"}
	tests := []struct {
		towels  []string
		pattern string
		out     []string
	}{
		{towels, "brwrr", []string{"rwrr", "wrr"}},
		{towels, "wrr", []string{"r"}},
		{towels, "r", []string{""}},
		{towels, "ubwu", []string{}},
	}

	for _, tc := range tests {
		t.Run(tc.pattern, func(t *testing.T) {
			assert.Equal(t, tc.out, day19.FindNext(tc.towels, tc.pattern))
		})
	}
}

func TestIsMatch(t *testing.T) {
	towels := []string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"}
	tests := []struct {
		towels  []string
		pattern string
		out     int
	}{
		{towels, "brwrr", 2},
		{towels, "bggr", 1},
		{towels, "gbbr", 4},
		{towels, "rrbgbr", 6},
		{towels, "ubwu", 0},
		{towels, "bwurrg", 1},
		{towels, "brgr", 2},
		{towels, "bbrgwb", 0},
	}

	for _, tc := range tests {
		t.Run(tc.pattern, func(t *testing.T) {
			assert.Equal(t, tc.out, day19.MatchCount(tc.towels, tc.pattern))
		})
	}
}
