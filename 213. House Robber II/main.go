package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(rob([]int{2, 3, 2}) == 3)
	fmt.Println(rob([]int{1, 2, 3, 1}) == 4)

}

func rob(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}
	p1 := robBottomUp(nums[0 : len(nums)-1])
	p2 := robBottomUp(nums[1:])

	if p1 > p2 {
		return p1
	} else {
		return p2
	}
}

func robBottomUp(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}

	p1, p2 := nums[0], 0

	for i := 1; i < len(nums); i++ {
		if p1 > p2+nums[i] {
			p2 = p1
		} else {
			p1, p2 = p2+nums[i], p1
		}
	}

	return p1
}

func robTopDown(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}
	cache := make([]interface{}, len(nums))
	cache[0] = nums[0]
	cache[1] = int(math.Max(float64(nums[0]), float64(nums[1])))
	var helper func([]int) int
	helper = func(nums []int) int {
		if len(nums) < 3 {
			return cache[len(nums)-1].(int)
		}
		if cache[len(nums)-1] != nil {
			return cache[len(nums)-1].(int)
		}
		if helper(nums[0:len(nums)-2])+nums[len(nums)-1] > helper(nums[0:len(nums)-1]) {
			cache[len(nums)-1] = helper(nums[0:len(nums)-2]) + nums[len(nums)-1]
		} else {
			cache[len(nums)-1] = helper(nums[0 : len(nums)-1])
		}
		return cache[len(nums)-1].(int)
	}

	return helper(nums)
}
