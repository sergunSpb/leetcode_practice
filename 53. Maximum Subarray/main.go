package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) == 6)
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}) == 23)
}

func maxSubArray(nums []int) int {
	max := math.MinInt32

	t_sum := 0
	for _, v := range nums {
		t_sum += v
		if t_sum > max {
			max = t_sum
		}
		if t_sum < 0 {
			t_sum = 0
		}
	}

	return max
}
