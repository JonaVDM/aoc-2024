package utils

type Point struct {
	X, Y int
}

// GetAdjacend returns the coordinates of all the connected squares.
// This includes the squares diagonal from it.
func GetAdjacend(x, y, height, width int) []Point {
	next := make([]Point, 0)

	for j := y - 1; j <= y+1; j++ {
		for i := x - 1; i <= x+1; i++ {
			if (i == x && j == y) || i < 0 || j < 0 || i >= width || j >= height {
				continue
			}

			next = append(next, Point{i, j})
		}
	}

	return next
}

// GetAdjacend returns the coordinates of all the connected squares.
// This does not include the squares diagonal from it.
func GetDirectAdjacend(x, y, height, width int) []Point {
	next := make([]Point, 0)

	pos := make([]Point, 4)
	pos[0] = Point{x - 1, y}
	pos[1] = Point{x, y - 1}
	pos[2] = Point{x + 1, y}
	pos[3] = Point{x, y + 1}

	for _, p := range pos {
		if p.X < 0 || p.Y < 0 || p.X >= width || p.Y >= height {
			continue
		}

		next = append(next, p)
	}

	return next
}

// GetRelativeAdjecend returns a list of all the adjecend sides from -1 to 1.
// Yes this is a hardcoded list.
func GetRelativeAdjecend() []Point {
	return []Point{
		{X: -1, Y: -1},
		{X: 0, Y: -1},
		{X: 1, Y: -1},
		{X: -1, Y: 0},
		{X: 1, Y: 0},
		{X: -1, Y: 1},
		{X: 0, Y: 1},
		{X: 1, Y: 1},
	}
}

func CopyMap(m map[string]interface{}) map[string]interface{} {
	cp := make(map[string]interface{})
	for k, v := range m {
		vm, ok := v.(map[string]interface{})
		if ok {
			cp[k] = CopyMap(vm)
		} else {
			cp[k] = v
		}
	}

	return cp
}

func InRange(x, y, w, h int) bool {
	return x >= 0 && x < w && y >= 0 && y < h
}

// Maze is a representation of a 2d map
type Maze[T any] map[int]map[int]T

func (m Maze[T]) Set(x, y int, t T) {
	if _, ok := m[x]; !ok {
		m[x] = make(map[int]T)
	}
	m[x][y] = t
}

func (m Maze[T]) Get(x, y int) T {
	return m[x][y]
}

// InitSimpleMaze takes in the input and turns it into a Maze object using a boolean for the wall
func InitSimpleMaze(data []string, wall rune) Maze[bool] {
	m := make(Maze[bool])

	for y, row := range data {
		for x, c := range row {
			if c == wall {
				m.Set(x, y, true)
			}
		}
	}

	return m
}
