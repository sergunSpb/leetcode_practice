package main

import "fmt"

func main() {
	fmt.Println(jump([]int{4, 1, 1, 3, 1, 1, 1}) == 2)
	fmt.Println(jump([]int{2, 3, 0, 1, 4}) == 2)
	fmt.Println(jump([]int{2, 3, 1, 1, 4}) == 2)

}

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	c, steps, curStepLength, curMax, curMaxInd := 0, 1, nums[0], 0, 0
	for in, val := range nums {
		if curStepLength < c {
			curStepLength = curMax
			c = in - curMaxInd
			steps++
		}
		if curMax < val-curMaxInd+in {
			curMax = val
			curMaxInd = in
		}
		c++
	}
	return steps
}
