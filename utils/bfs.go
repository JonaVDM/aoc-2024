package utils

type BFSOBject interface {
	Next(p Point) []Point
}

// BFS returns the shortest path to a point
//
// The BFSObject should have a function called Next, that function returns the
// the next set of coords based on some given coord. The Next function should
// **not** keep track of which points have already been seen.
func BFS(start, end Point, obj BFSOBject) []Point {
	q := make([]Point, 1)
	q[0] = start

	prev := make(map[Point]Point)
	path := make([]Point, 0)

	for len(q) > 0 {
		i := q[0]
		q = q[1:]

		if i == end {
			n := i
			path = append(path, n)
			for {
				n = prev[n]
				path = append(path, n)
				if n == start {
					break
				}
			}
			return path
		}

		for _, n := range obj.Next(i) {
			if _, ok := prev[n]; ok {
				continue
			}

			prev[n] = i
			q = append(q, n)
		}
	}

	return path
}
