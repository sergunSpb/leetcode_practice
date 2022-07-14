package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}) == 6)
	fmt.Println(maxProduct([]int{-2}) == -2)
	fmt.Println(maxProduct([]int{-2, 4, -3}) == 24)
	fmt.Println(maxProduct([]int{-2, -4, -3}) == 12)
}

func maxProduct(nums []int) int {
	max := nums[0]
	cMax, cMin := max, max

	for i := 1; i < len(nums); i++ {
		cur := nums[i]
		if cur < 0 {
			cMin, cMax = cMax, cMin
		}
		cMin = int(math.Min(float64(cMin*cur), float64(cur)))
		cMax = int(math.Max(float64(cMax*cur), float64(cur)))

		if cMax > max {
			max = cMax
		}

	}

	return max
}
