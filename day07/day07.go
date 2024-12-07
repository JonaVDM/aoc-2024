package day07

import (
	"strconv"
	"strings"
	"sync"

	"github.com/jonavdm/aoc-2024/utils"
)

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)

	a := 0
	b := 0

	ac := make(chan int)
	bc := make(chan int)
	cc := make(chan int)

	var wg sync.WaitGroup

	for _, row := range data {
		wg.Add(1)
		go func() {
			defer wg.Done()
			search, nums := Parse(row)
			ac <- Sovleble(search, nums, false)
			bc <- Sovleble(search, nums, true)
		}()
	}
	go func() {
		wg.Wait()
		close(ac)
		close(bc)
		cc <- 1
	}()

OUTER:
	for {
		select {
		case an := <-ac:
			a += an
		case bn := <-bc:
			b += bn
		case <-cc:
			break OUTER
		}
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

		m := utils.PowInt(10, utils.NumberLength(i.Rem[0]))
		n3 := i.N*m + i.Rem[0]

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
