package day07_test

import (
	"strings"
	"testing"

	"github.com/jonavdm/aoc-2024/day07"
	_ "github.com/jonavdm/aoc-2024/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	test := strings.Split(`190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`, "\n")
	assert.Equal(t, [2]interface{}{3749, 11387}, day07.Run(test))
}
