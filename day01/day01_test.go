package day01_test

import (
	"testing"

	"github.com/jonavdm/aoc-2024/day01"
	_ "github.com/jonavdm/aoc-2024/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	testFile := []string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}

	assert.Equal(t, [2]interface{}{11, 31}, day01.Run(testFile))
}
