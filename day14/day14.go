package day14

import (
	"slices"
	"strconv"
	"strings"

	"github.com/jonavdm/aoc-2024/utils"
)

const HEIGHT int = 103
const WIDTH int = 101

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)

	corner := make([]int, 4)
	f := make(Field, len(data))

	for i, row := range data {
		g := Parse(row)
		f[i] = g
		end := g.Simulate(100)
		c := GetCorner(end)
		if c == -1 {
			continue
		}
		corner[c]++
	}

	tree := -1
	for i := 1; i < 1_000_000; i++ {
		f.Move()
		if f.IsTree() {
			tree = i
			break
		}
	}

	return [2]interface{}{
		corner[0] * corner[1] * corner[2] * corner[3],
		tree,
	}
}

func Parse(row string) Guard {
	parts := strings.Split(row, " ")

	pos := strings.Replace(parts[0], "p=", "", 1)
	posP := strings.Split(pos, ",")
	position := utils.Point{}
	position.X, _ = strconv.Atoi(posP[0])
	position.Y, _ = strconv.Atoi(posP[1])

	vel := strings.Replace(parts[1], "v=", "", 1)
	velP := strings.Split(vel, ",")
	velocity := utils.Point{}
	velocity.X, _ = strconv.Atoi(velP[0])
	velocity.Y, _ = strconv.Atoi(velP[1])

	return Guard{
		P: position,
		V: velocity,
	}
}

type Field []Guard

func (f *Field) Visual() string {
	out := make([]string, HEIGHT)

	for i := 0; i < HEIGHT; i++ {
		out[i] = strings.Repeat(".", WIDTH)
	}

	for _, g := range *f {
		tmp := []rune(out[g.P.Y])
		tmp[g.P.X] = '#'
		out[g.P.Y] = string(tmp)
	}

	return strings.Join(out, "\n")
}

func (f Field) Move() {
	for i := range f {
		f[i].Move()
	}
}

func (f Field) IsTree() bool {
	m := make(map[int][]int)

	for _, g := range f {
		if _, ok := m[g.P.X]; !ok {
			m[g.P.X] = make([]int, 0)
		}
		m[g.P.X] = append(m[g.P.X], g.P.Y)
	}

	for _, c := range m {
		if len(c) < 10 {
			continue
		}

		slices.Sort(c)
		var inRow int
		for i := 0; i < len(c)-1; i++ {
			if c[i+1]-c[i] != 1 {
				inRow = 0
				continue
			}
			inRow++

			if inRow == 10 {
				return true
			}
		}
	}

	return false
}

type Guard struct {
	P utils.Point
	V utils.Point
}

func (g *Guard) Simulate(amount int) utils.Point {
	x := (g.P.X + g.V.X*amount) % WIDTH
	if x < 0 {
		x += WIDTH
	}

	y := (g.P.Y + g.V.Y*amount) % HEIGHT
	if y < 0 {
		y += HEIGHT
	}
	return utils.Point{X: x, Y: y}
}

func (g *Guard) Move() {
	x := (g.P.X + g.V.X) % WIDTH
	if x < 0 {
		x += WIDTH
	}

	y := (g.P.Y + g.V.Y) % HEIGHT
	if y < 0 {
		y += HEIGHT
	}
	g.P.X = x
	g.P.Y = y
}

func GetCorner(pos utils.Point) int {
	if pos.X < WIDTH/2 && pos.Y < HEIGHT/2 {
		return 0
	}

	if pos.X < WIDTH/2 && pos.Y > HEIGHT/2 {
		return 1
	}

	if pos.X > WIDTH/2 && pos.Y < HEIGHT/2 {
		return 2
	}

	if pos.X > WIDTH/2 && pos.Y > HEIGHT/2 {
		return 3
	}

	return -1
}
