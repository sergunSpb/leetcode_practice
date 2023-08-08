package main

import (
	"math"
)

func main() {
	println(maximumSum([]int{-1}) == -1)
	println(maximumSum([]int{-1, -1, -1, -1}) == -1)
	println(maximumSum([]int{1, 2, 3, 1}) == 7)
	println(maximumSum([]int{1, -2, 0, 3}) == 4)
	println(maximumSum([]int{1, -2, -2, 3}) == 3)
	println(maximumSum([]int{11, -10, -11, 8, 7, -6, 9, 4, 11, 6, 5, 0}) == 50)
	println(maximumSum([]int{11, -10, -11, 8, 7, -6, 9, 4, 11, 6, 5, 0}) == 50)
	println(maximumSum([]int{2, 1, -2, -5, -2}) == 3)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func kadana(arr []int) int {
	max := math.MinInt
	cur := 0
	for _, i := range arr {
		cur += i
		if cur > max {
			max = cur
		}
		if cur < 0 {
			cur = 0
		}
	}

	return max
}

func maximumSum(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	// for case when no deletions
	max := kadana(arr)

	//for cases when up to 1 deletion
	mToRight := make([]int, len(arr))
	mToRight[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		if mToRight[i-1] > 0 {
			mToRight[i] = mToRight[i-1] + arr[i]
		} else {
			mToRight[i] = arr[i]
		}
	}

	mToLeft := make([]int, len(arr))
	mToLeft[len(arr)-1] = arr[len(arr)-1]
	for i := len(arr) - 2; i >= 0; i-- {
		if mToLeft[i+1] > 0 {
			mToLeft[i] = mToLeft[i+1] + arr[i]
		} else {
			mToLeft[i] = arr[i]
		}
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}

		if i == 0 {
			if mToLeft[i+1] > max {
				max = mToLeft[i+1]
			}
			continue
		}
		if i == len(arr)-1 {
			if mToRight[i-1] > max {
				max = mToRight[i-1]
			}
			continue
		}
		if mToLeft[i+1]+mToRight[i-1] > max {
			max = mToLeft[i+1] + mToRight[i-1]
		}

		if mToLeft[i+1]+mToRight[i-1] > max {
			max = mToLeft[i+1] + mToRight[i-1]
		}
	}

	return max
}
