package main

import (
	"fmt"
)

func main() {
	fmt.Println(getMaxLen([]int{5, -20, -20, -39, -5, 0, 0, 0, 36, -32, 0, -7, -10, -7, 21, 20, -12, -34, 26, 2}) == 8)
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func getMaxLen(nums []int) int {
	var max, cNegEl, rNeg, length, cP, lNeg int
	lNeg = -1

	getMaxHelper := func(nLeftPos, nRightPos, size, amNeg int) int {
		if amNeg%2 == 0 {
			return size
		}
		return Max(size-nLeftPos-1, nRightPos)
	}
	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		if cur == 0 {
			cMax := getMaxHelper(lNeg, rNeg, length, cNegEl)
			if cMax > max {
				max = cMax
			}
			cNegEl, length, lNeg, rNeg, cP = 0, 0, -1, 0, 0
			continue
		}
		length++
		if cur < 0 {
			cNegEl++
			if lNeg < 0 {
				lNeg = cP
			}
			rNeg = cP
		}
		cP++
	}
	if nums[len(nums)-1] != 0 {
		cMax := getMaxHelper(lNeg, rNeg, length, cNegEl)
		if cMax > max {
			max = cMax
		}
	}

	return max
}
