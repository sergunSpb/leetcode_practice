package main

func main() {
	println(maxProfit([]int{1, 2, 3}) == 2)

	println(maxProfit([]int{1, 2, 3, 0, 2}) == 3)
	println(maxProfit([]int{1, 2, 3, 0, 10}) == 11)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	sell := make([]int, len(prices))
	buy := make([]int, len(prices))
	rest := make([]int, len(prices))

	sell[0] = 0
	buy[0] = -prices[0]
	rest[0] = 0
	for i := 1; i < len(prices); i++ {
		sell[i] = max(sell[i-1], buy[i-1]+prices[i])
		buy[i] = max(buy[i-1], rest[i-1]-prices[i])
		rest[i] = max(rest[i-1], max(buy[i-1], sell[i-1]))
	}

	return sell[len(prices)-1]
}
