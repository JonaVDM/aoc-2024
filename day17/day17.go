package day17

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jonavdm/aoc-2024/utils"
)

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)

	a, _ := strconv.Atoi(strings.Split(data[0], " ")[2])
	b, _ := strconv.Atoi(strings.Split(data[1], " ")[2])
	c, _ := strconv.Atoi(strings.Split(data[2], " ")[2])

	instr := strings.Split(data[4], " ")[1]
	instNums := make([]int, 0)

	for _, n := range strings.Split(instr, ",") {
		num, _ := strconv.Atoi(n)
		instNums = append(instNums, num)
	}

	m := Machine{
		a, b, c, instNums,
	}

	return [2]interface{}{
		m.Simulate(),
		0,
	}
}

type Machine struct {
	A, B, C int
	Inst    []int
}

func (m *Machine) Val(index int) int {
	v := m.Inst[index]
	if v <= 3 {
		return v
	}

	if v == 4 {
		return m.A
	}

	if v == 5 {
		return m.B
	}

	if v == 6 {
		return m.C
	}

	panic("invalid instruction")
}

func (m *Machine) Simulate() string {
	i := 0
	out := make([]int, 0)

	for i < len(m.Inst) {
		switch m.Inst[i] {
		case 0:
			m.A = m.A / utils.PowInt(2, m.Val(i+1))
		case 1:
			m.B = m.B ^ m.Inst[i+1]
		case 2:
			m.B = m.Val(i+1) % 8
		case 3:
			if m.A == 0 {
				break
			}
			i = m.Inst[i+1] - 2
		case 4:
			m.B = m.B ^ m.C
		case 5:
			out = append(out, m.Val(i+1)%8)
		case 6:
			m.B = m.A / utils.PowInt(2, m.Val(i+1))
		case 7:
			m.C = m.A / utils.PowInt(2, m.Val(i+1))
		}

		i += 2
	}

	strout := make([]string, len(out))
	for i, n := range out {
		strout[i] = fmt.Sprint(n)
	}
	return strings.Join(strout, ",")
}
