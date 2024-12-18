package day15

import (
	"strings"

	"github.com/jonavdm/aoc-2024/utils"
)

func Run(data []string) [2]interface{} {
	room := Parse(data)
	room.do()

	return [2]interface{}{
		room.count(),
		0,
	}
}

type Room struct {
	items        map[int]map[int]Item
	player       utils.Point
	instructions string
}

type Item struct {
	isWall bool
	isBox  bool
}

func Parse(data []string) Room {
	room := Room{
		items: make(map[int]map[int]Item),
	}
	for y, l := range data {
		if l == "" {
			room.instructions = strings.Join(data[y+1:], "")
			break
		}

		room.items[y] = make(map[int]Item)

		for x, c := range l {
			room.items[y][x] = Item{
				isWall: c == '#',
				isBox:  c == 'O',
			}

			if c == '@' {
				room.player = utils.Point{
					X: x,
					Y: y,
				}
			}
		}
	}

	return room
}

func (r *Room) do() {
	for _, c := range r.instructions {
		var v utils.Point
		switch c {
		case '^':
			v.Y = -1
		case '>':
			v.X = 1
		case 'v':
			v.Y = 1
		case '<':
			v.X = -1
		}

		next := utils.Point{X: r.player.X + v.X, Y: r.player.Y + v.Y}
		if !r.move(next, v) {
			continue
		}
		r.player = next
	}
}

func (r *Room) at(pos utils.Point) Item {
	return r.items[pos.Y][pos.X]
}

func (r *Room) move(pos, vel utils.Point) bool {
	next := utils.Point{X: pos.X + vel.X, Y: pos.Y + vel.Y}
	if r.at(pos).isWall {
		return false
	}

	if r.at(pos).isBox {
		if !r.move(next, vel) {
			return false
		}

		r.items[pos.Y][pos.X] = Item{}
		r.items[next.Y][next.X] = Item{isBox: true}
	}

	return true
}

func (r *Room) Visulize() string {
	out := make([]string, len(r.items))

	for i := range out {
		out[i] = strings.Repeat(".", len(r.items[0]))
	}

	tmp := []rune(out[r.player.Y])
	tmp[r.player.X] = '@'
	out[r.player.Y] = string(tmp)

	for i, m := range r.items {
		for j, c := range m {
			tmp := []rune(out[i])
			if c.isBox {
				tmp[j] = 'O'
			} else if c.isWall {
				tmp[j] = '#'
			}
			out[i] = string(tmp)
		}
	}

	return strings.Join(out, "\n")
}

func (r *Room) count() int {
	var count int

	for i, m := range r.items {
		for j, c := range m {
			if !c.isBox {
				continue
			}

			count += i*100 + j
		}
	}
	return count
}
