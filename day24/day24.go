package day24

import (
	"fmt"
	"slices"
	"strings"

	"github.com/jonavdm/aoc-2024/utils"
)

type Solver struct {
	Instructions map[string]struct {
		A, B     string
		Operator string
	}
	Register map[string]bool
}

func Run(data []string) [2]interface{} {
	solver := Parse(data)
	var size int

	for k := range solver.Instructions {
		if k[0] == 'z' {
			size++
		}
	}

	var total int
	for i := range size {
		val := solver.GetReg(fmt.Sprintf("z%02d", i))
		if val {
			total += utils.PowInt(2, i)
		}
	}

	return [2]interface{}{
		total,
		0,
	}
}

func Parse(data []string) Solver {
	s := Solver{
		Instructions: make(map[string]struct {
			A        string
			B        string
			Operator string
		}),
		Register: make(map[string]bool),
	}

	split := slices.Index(data, "")

	for _, d := range data[:split] {
		spl := strings.Split(d, ": ")
		s.Register[spl[0]] = spl[1] == "1"
	}

	for _, d := range data[split+1:] {
		spl := strings.Split(d, " ")
		s.Instructions[spl[4]] = struct {
			A        string
			B        string
			Operator string
		}{
			spl[0], spl[2], spl[1],
		}
	}

	return s
}

func RunInstruction(a, b bool, operator string) bool {
	switch operator {
	case "AND":
		return a && b
	case "OR":
		return a || b
	case "XOR":
		return a != b
	default:
		panic("what")
	}
}

func (s *Solver) GetReg(reg string) bool {
	v, ok := s.Register[reg]
	if ok {
		return v
	}

	inst, ok := s.Instructions[reg]
	if !ok {
		panic("Instruction not found " + reg)
	}
	a := s.GetReg(inst.A)
	b := s.GetReg(inst.B)
	v = RunInstruction(a, b, inst.Operator)
	s.Register[reg] = v
	return v
}
