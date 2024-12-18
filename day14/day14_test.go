package day14_test

import (
	"strings"
	"testing"

	"github.com/jonavdm/aoc-2024/day14"
	_ "github.com/jonavdm/aoc-2024/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	file := strings.Split(`p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`, "\n")

	// TODO: add some fancy injection
	assert.Equal(t, [2]interface{}{21, -1}, day14.Run(file))
}
