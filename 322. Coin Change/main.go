package main

import "math"

func main() {
	s := coinChange([]int{1}, 1)
	println(s)
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	dp := make([]int, amount)
	for i := range dp {
		dp[i] = math.MaxInt
	}
	for y := 1; y <= amount; y++ {
		for _, coin := range coins {
			if coin > y {
				continue
			}
			if coin == y {
				dp[y-1] = 1
				continue
			}

			dp[y-1] = min(dp[y-(coin+1)]+1, dp[y-1])
		}
	}
	if dp[len(dp)-1] == math.MaxInt {
		return -1
	}

	return dp[len(dp)-1]
}
