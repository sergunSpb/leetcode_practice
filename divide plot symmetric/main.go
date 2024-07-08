package main

import (
	"errors"
	"fmt"
)

type Point struct {
	X int
	Y int
}

func (p Point) Key() string {
	return fmt.Sprintf("%d:%d", p.X, p.Y)
}

func main() {
	t()
	// println(checkLine([]Point{{1, 1}, {3, 1}, {3, 1}, {1, 1}, {2, 1}}))
}

func t() (v int, err error) {
	defer func() {
		println(v)
		println(err)
	}()

	return 1, errors.New("s")
}

func checkLine(ps []Point) bool {
	maxX, minX := ps[0].X, ps[0].X
	seen := make(map[Point]int, len(ps))
	for _, p := range ps {
		maxX = max(maxX, p.X)
		minX = min(minX, p.X)
		seen[p]++
	}

	mean := (maxX + minX) / 2
	for _, p := range ps {
		if p.X == mean {
			continue
		}

		if _, ok := seen[Point{Y: p.Y, X: 2*mean - p.X}]; ok {
			seen[Point{Y: p.Y, X: 2*mean - p.X}]--
		} else if !ok {
			return false
		}
	}

	for v, p := range seen {
		if v.X == mean {
			continue
		}
		if p != 0 {
			return false
		}
	}

	return true
}

func max(i, y int) int {
	if y > i {
		return y
	}
	return i
}

func min(i, y int) int {
	if y < i {
		return y
	}
	return i
}
