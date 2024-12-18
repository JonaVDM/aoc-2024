package day11_test

import (
	"testing"

	"github.com/jonavdm/aoc-2024/day11"
	_ "github.com/jonavdm/aoc-2024/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	file := []string{"125 17"}
	assert.Equal(t, [2]interface{}{55312, 65601038650482}, day11.Run(file))
}
