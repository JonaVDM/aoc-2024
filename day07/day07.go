package day07

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jonavdm/aoc-2024/utils"
)

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)

	a := 0
	b := 0

	for _, row := range data {
		search, nums := Parse(row)
		a += Sovleble(search, nums, false)
		b += Sovleble(search, nums, true)
	}

	return [2]interface{}{
		a,
		b,
	}
}

func Parse(row string) (int, []int) {
	row = strings.ReplaceAll(row, ":", "")
	spl := strings.Split(row, " ")
	nums := make([]int, len(spl))
	for i, num := range spl {
		con, _ := strconv.Atoi(num)
		nums[i] = con
	}

	return nums[0], nums[1:]
}

type Row struct {
	N   int
	Rem []int
}

func Sovleble(search int, nums []int, two bool) int {
	q := make([]Row, 1)
	q[0] = Row{N: nums[0], Rem: nums[1:]}

	for len(q) > 0 {
		i := q[0]
		q = q[1:]

		n1 := i.N + i.Rem[0]
		n2 := i.N * i.Rem[0]
		n3, _ := strconv.Atoi(fmt.Sprintf("%d%d", i.N, i.Rem[0]))

		if len(i.Rem) > 1 {
			q = append(
				q,
				Row{N: n1, Rem: i.Rem[1:]},
				Row{N: n2, Rem: i.Rem[1:]},
			)

			if two {
				q = append(q, Row{N: n3, Rem: i.Rem[1:]})
			}
			continue
		}

		if n1 == search || n2 == search || (two && n3 == search) {
			return search
		}
	}

	return 0
}
