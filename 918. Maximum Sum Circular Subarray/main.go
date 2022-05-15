package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSubarraySumCircular([]int{1, -2, 3, -2}) == 3)
	fmt.Println(maxSubarraySumCircular([]int{-3, -2, -3}) == -2)
	fmt.Println(maxSubarraySumCircular([]int{5, -3, 5}) == 10)
	fmt.Println(maxSubarraySumCircular([]int{-1, 3, -3, 9, -6, 8, -5, -5, -6, 10}) == 20)
	fmt.Println(maxSubarraySumCircular([]int{3, -1, 2, -1}) == 4)
}

func maxSubarraySumCircular(nums []int) int {
	max, t_sum := math.MinInt32, 0
	for in := 0; in < len(nums); in++ {
		t_sum += nums[in]
		if t_sum > max {
			max = t_sum
		}
		if t_sum < 0 {
			t_sum = 0
		}
	}

	max1 := max
	max, t_sum, maxSLBeg := math.MinInt32, 0, []int{}
	for in := 0; in < len(nums); in++ {
		t_sum += nums[in]
		if t_sum > max {
			max = t_sum
		}
		maxSLBeg = append(maxSLBeg, max)
	}

	max, t_sum, maxSLEnd := math.MinInt32, 0, make([]int, len(nums)+1, len(nums)+1)
	for in := len(nums); in > 0; in-- {
		t_sum += nums[in-1]
		if t_sum > max {
			max = t_sum
		}
		maxSLEnd[in-1] = max
	}

	max = math.MinInt32
	for i := 0; i < len(nums); i++ {
		if maxSLEnd[i+1]+maxSLBeg[i] > max {
			max = maxSLEnd[i+1] + maxSLBeg[i]
		}
	}

	if max1 > max {
		max = max1
	}

	return max
}
