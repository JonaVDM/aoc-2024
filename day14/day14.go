package day14

import (
	"strconv"
	"strings"

	"github.com/jonavdm/aoc-2024/utils"
)

const HEIGHT int = 103
const WIDTH int = 101

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)

	corner := make([]int, 4)

	for _, row := range data {
		pos, vel := Parse(row)
		end := Simulate(pos, vel, 100)
		c := GetCorner(end)
		if c == -1 {
			continue
		}
		corner[c]++
	}

	return [2]interface{}{
		corner[0] * corner[1] * corner[2] * corner[3],
		0,
	}
}

func Parse(row string) (utils.Point, utils.Point) {
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

	return position, velocity
}

func Simulate(pos, vel utils.Point, amount int) utils.Point {
	x := (pos.X + vel.X*amount) % WIDTH
	if x < 0 {
		x += WIDTH
	}

	y := (pos.Y + vel.Y*amount) % HEIGHT
	if y < 0 {
		y += HEIGHT
	}
	return utils.Point{X: x, Y: y}
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
