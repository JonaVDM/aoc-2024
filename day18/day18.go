package day18

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/jonavdm/aoc-2024/utils"
)

const SIZE int = 70

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)

	r := Parse(data, 1024)

	return [2]interface{}{
		r.BFS(),
		r.FindLast(1024),
	}
}

func Parse(data []string, amount int) Room {
	room := Room{
		Room:   make(map[int]map[int]bool),
		Range:  SIZE,
		Coords: make([]utils.Point, len(data)),
	}

	for i, r := range data {
		spl := strings.Split(r, ",")
		x, _ := strconv.Atoi(spl[0])
		y, _ := strconv.Atoi(spl[1])

		room.Coords[i] = utils.Point{
			X: x, Y: y,
		}
	}

	for i, c := range room.Coords {
		if i >= amount {
			break
		}
		room.Set(c.X, c.Y)
	}

	return room
}

type Room struct {
	Room   map[int]map[int]bool
	Range  int
	Coords []utils.Point
	Prev   []utils.Point
}

func (r *Room) Set(x, y int) {
	if _, ok := r.Room[x]; !ok {
		r.Room[x] = make(map[int]bool)
	}
	r.Room[x][y] = true
}

func (r *Room) Next(p utils.Point) []utils.Point {
	next := make([]utils.Point, 0)

	for _, n := range utils.GetDirectAdjacend(p.X, p.Y, r.Range+1, r.Range+1) {
		if !r.Room[n.X][n.Y] {
			next = append(next, n)
		}
	}

	return next
}

func (r *Room) BFS() int {
	start := utils.Point{X: 0, Y: 0}
	goal := utils.Point{X: r.Range, Y: r.Range}

	path := utils.BFS(start, goal, r)
	r.Prev = path

	return len(path) - 1
}

func (r *Room) FindLast(minNum int) string {
	r.BFS()
	for {
		minNum++
		r.Set(r.Coords[minNum].X, r.Coords[minNum].Y)
		if !slices.Contains(r.Prev, r.Coords[minNum]) {
			continue
		}
		if r.BFS() == -1 {
			break
		}
	}

	return fmt.Sprintf(
		"%d,%d",
		r.Coords[minNum].X,
		r.Coords[minNum].Y,
	)
}
