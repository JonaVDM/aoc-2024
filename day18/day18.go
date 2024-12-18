package day18

import (
	"strconv"
	"strings"

	"github.com/jonavdm/aoc-2024/utils"
)

const SIZE int = 70

type Room struct {
	Room  map[int]map[int]bool
	Range int
}

func (r *Room) Set(x, y int) {
	if _, ok := r.Room[x]; !ok {
		r.Room[x] = make(map[int]bool)
	}
	r.Room[x][y] = true
}

func (r Room) BFS() int {
	q := make([]utils.Point, 1)
	been := make(map[utils.Point]int)
	q[0] = utils.Point{X: 0, Y: 0}
	been[q[0]] = 0
	goal := utils.Point{X: r.Range, Y: r.Range}

	for len(q) > 0 {
		i := q[0]
		q = q[1:]

		if i == goal {
			return been[i]
		}

		for _, n := range utils.GetDirectAdjacend(i.X, i.Y, r.Range+1, r.Range+1) {
			if _, ok := been[n]; ok || r.Room[n.X][n.Y] {
				continue
			}

			been[n] = been[i] + 1
			q = append(q, n)
		}
	}

	return -1
}

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)

	r := Parse(data, 1024)

	return [2]interface{}{
		r.BFS(),
		0,
	}
}

func Parse(data []string, amount int) Room {
	room := Room{
		Room:  make(map[int]map[int]bool),
		Range: SIZE,
	}

	for i, r := range data {
		if i >= amount {
			break
		}

		spl := strings.Split(r, ",")
		x, _ := strconv.Atoi(spl[0])
		y, _ := strconv.Atoi(spl[1])
		room.Set(x, y)
	}

	return room
}
