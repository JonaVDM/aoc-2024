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
