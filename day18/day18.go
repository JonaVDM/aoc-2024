package day18

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/jonavdm/aoc-2024/utils"
)

const SIZE int = 70

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

func (r *Room) BFS() int {
	start := utils.Point{X: 0, Y: 0}
	goal := utils.Point{X: r.Range, Y: r.Range}

	q := make([]utils.Point, 1)
	q[0] = start
	prev := make(map[utils.Point]utils.Point)

	r.Prev = make([]utils.Point, 0)

	for len(q) > 0 {
		i := q[0]
		q = q[1:]

		if i == goal {
			var count int
			n := i
			for {
				n = prev[n]
				r.Prev = append(r.Prev, n)
				count++
				if n == start {
					break
				}
			}
			return count
		}

		for _, n := range utils.GetDirectAdjacend(i.X, i.Y, r.Range+1, r.Range+1) {
			if _, ok := prev[n]; ok || r.Room[n.X][n.Y] {
				continue
			}

			prev[n] = i
			q = append(q, n)
		}
	}

	return -1
}

func (r *Room) FindLast(minNum int) string {
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
