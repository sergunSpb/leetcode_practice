package main

func maximumTotalCost(nums []int) int64 {
	dp := make([][2]int64, len(nums))
	dp[0][0] = int64(nums[0])
	dp[0][1] = dp[0][0]

	for i := 1; i < len(nums); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]) + int64(nums[i])
		dp[i][1] = dp[i-1][0] - int64(nums[i])
	}

	return max(dp[len(nums)-1][0], dp[len(nums)-1][1])
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}

	return b
}
