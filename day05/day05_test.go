package day05_test

import (
	"testing"

	"github.com/jonavdm/aoc-2024/day05"
	_ "github.com/jonavdm/aoc-2024/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	test := []string{
		"47|53",
		"97|13",
		"97|61",
		"97|47",
		"75|29",
		"61|13",
		"75|53",
		"29|13",
		"97|29",
		"53|29",
		"61|53",
		"97|53",
		"61|29",
		"47|13",
		"75|47",
		"97|75",
		"47|61",
		"75|61",
		"47|29",
		"75|13",
		"53|13",
		"",
		"75,47,61,53,29",
		"97,61,53,29,13",
		"75,29,13",
		"75,97,47,61,53",
		"61,13,29",
		"97,13,75,29,47",
	}
	assert.Equal(t, [2]interface{}{143, 123}, day05.Run(test))
}

func TestSolver_FixPrint(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"75,97,47,61,53", "97,75,47,61,53"},
		{"61,13,29", "61,29,13"},
		{"97,13,75,29,47", "97,75,47,29,13"},
	}

	solver := day05.Solver{
		Orders: map[string][]string{
			"61": {"13", "53", "29"},
			"29": {"13"},
			"53": {"29", "13"},
			"47": {"53", "13", "61", "29"},
			"97": {"13", "61", "47", "29", "53", "75"},
			"75": {"29", "53", "47", "61", "13"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.in, func(t *testing.T) {
			assert.Equal(t, tc.out, solver.FixPrint(tc.in))
		})
	}
}
