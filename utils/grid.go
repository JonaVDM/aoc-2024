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
