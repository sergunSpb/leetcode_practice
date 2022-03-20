package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(countSquaresBruteForce([][]int{
		{0, 1, 1, 1},
		{1, 1, 1, 1},
		{0, 1, 1, 1},
	}))

	fmt.Println(countSquaresBruteForce([][]int{
		{1, 0, 1},
		{1, 1, 0},
		{1, 1, 0},
	}))

	fmt.Println(countSquaresBruteForce([][]int{
		{1, 0, 1},
	}))

	fmt.Println(countSquaresDP([][]int{
		{0, 1, 1, 1},
		{1, 1, 1, 1},
		{0, 1, 1, 1},
	}))

	fmt.Println(countSquaresDP([][]int{
		{1, 0, 1},
		{1, 1, 0},
		{1, 1, 0},
	}))

	fmt.Println(countSquaresDP([][]int{
		{1, 0, 1},
	}))

}

func countSquaresBruteForce(matrix [][]int) int {
	cnt := 0

	for y := range matrix {
		for x := range matrix[y] {
			if matrix[y][x] == 0 {
				continue
			}
			maxLen := int(math.Min(float64(len(matrix)-y), float64(len(matrix[y])-x)))
		inner:
			for l := 1; l <= maxLen; l++ {
				for y1 := y; y1 < y+l; y1++ {
					for x1 := x; x1 < x+l; x1++ {
						if matrix[y1][x1] != 1 {
							continue inner
						}
					}
				}
				cnt++
			}
		}
	}
	return cnt
}

func countSquaresDP(matrix [][]int) int {
	var cache = make([][]int, len(matrix))
	for y := range matrix {
		cache[y] = make([]int, len(matrix[y]))
	}
	cnt := 0
	for y := range matrix {
		for x := range matrix[y] {
			if matrix[y][x] == 0 {
				continue
			}
			if x == 0 || y == 0 {
				cache[y][x] = 1
			} else {
				cache[y][x] = 1 + int(math.Min(float64(cache[y-1][x]), math.Min(float64(cache[y][x-1]), float64(cache[y-1][x-1]))))
			}
			cnt += cache[y][x]
		}
	}

	return cnt
}
