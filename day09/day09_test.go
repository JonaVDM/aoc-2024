package day09_test

import (
	"testing"

	"github.com/jonavdm/aoc-2024/day09"
	_ "github.com/jonavdm/aoc-2024/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	test := []string{"2333133121414131402"}
	assert.Equal(t, [2]interface{}{1928, 0}, day09.Run(test))
}
