package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%+v\n", productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Printf("%+v\n", productExceptSelf([]int{1, 2, 3, 4, 0}))
	fmt.Printf("%+v\n", productExceptSelf([]int{1, 2, 3, 4, 0, 0}))
	fmt.Printf("%+v\n", productExceptSelf([]int{-1, -1, 3, 4}))

}

func productExceptSelf(nums []int) []int {
	var ar [61]int
	zeroPos := -1

	for i := range nums {
		if nums[i] == 0 {
			if zeroPos >= 0 {
				return make([]int, len(nums))
			}
			zeroPos = i
		}

		if nums[i] == 1 {
			continue
		}

		ar[nums[i]+30]++
	}

	for i := range nums {
		if zeroPos >= 0 && zeroPos != i {
			nums[i] = 0
			continue
		}

		pr := 1
		for y := range ar {
			count := ar[y]
			if y-30 == nums[i] {
				count--
			}

			if count > 0 {
				pr *= int(math.Pow(float64(y-30), float64(count)))
			}
		}
		nums[i] = pr
	}

	return nums
}
