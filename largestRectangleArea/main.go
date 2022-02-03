package main

import "fmt"

func main() {
	a := largestRectangleArea([]int{2, 1, 5, 6, 2, 3})
	fmt.Println(a)
}

func largestRectangleArea(heights []int) int {
	len := len(heights)
	max := 0
	for i := 0; i < len; i++ {
		cur_min_height := heights[i]
		cur_max_square := heights[i]
		for y := i; y < len; y++ {
			if heights[y] < cur_min_height {
				cur_min_height = heights[y]
			}
			t_c := (y - i + 1) * cur_min_height
			if t_c > cur_max_square {
				cur_max_square = t_c
			}
		}
		if cur_max_square > max {
			max = cur_max_square
		}
	}
	return max
}
