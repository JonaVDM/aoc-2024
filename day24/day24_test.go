package day24_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jonavdm/aoc-2024/day24"
	_ "github.com/jonavdm/aoc-2024/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	file := strings.Split(`x00: 1
x01: 0
x02: 1
x03: 1
x04: 0
y00: 1
y01: 1
y02: 1
y03: 1
y04: 1

ntg XOR fgs -> mjb
y02 OR x01 -> tnw
kwq OR kpj -> z05
x00 OR x03 -> fst
tgd XOR rvg -> z01
vdt OR tnw -> bfw
bfw AND frj -> z10
ffh OR nrd -> bqk
y00 AND y03 -> djm
y03 OR y00 -> psh
bqk OR frj -> z08
tnw OR fst -> frj
gnj AND tgd -> z11
bfw XOR mjb -> z00
x03 OR x00 -> vdt
gnj AND wpb -> z02
x04 AND y00 -> kjc
djm OR pbm -> qhw
nrd AND vdt -> hwm
kjc AND fst -> rvg
y04 OR y02 -> fgs
y01 AND x02 -> pbm
ntg OR kjc -> kwq
psh XOR fgs -> tgd
qhw XOR tgd -> z09
pbm OR djm -> kpj
x03 XOR y03 -> ffh
x00 XOR y04 -> ntg
bfw OR bqk -> z06
nrd XOR fgs -> wpb
frj XOR qhw -> z04
bqk OR frj -> z07
y03 OR x01 -> nrd
hwm AND bqk -> z03
tgd XOR rvg -> z12
tnw OR pbm -> gnj`, "\n")
	assert.Equal(t, [2]interface{}{2024, 0}, day24.Run(file))
}

func TestRunInstruction(t *testing.T) {
	tests := []struct {
		a, b     bool
		operator string
		out      bool
	}{
		{false, false, "AND", false},
		{false, true, "AND", false},
		{true, false, "AND", false},
		{true, true, "AND", true},

		{false, false, "OR", false},
		{false, true, "OR", true},
		{true, false, "OR", true},
		{true, true, "OR", true},

		{false, false, "XOR", false},
		{false, true, "XOR", true},
		{true, false, "XOR", true},
		{true, true, "XOR", false},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprint(tc.a, tc.operator, tc.b), func(t *testing.T) {
			assert.Equal(t, tc.out, day24.RunInstruction(tc.a, tc.b, tc.operator))
		})
	}
}
