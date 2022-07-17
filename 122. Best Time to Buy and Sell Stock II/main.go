package main

func main() {
	println(maxProfit([]int{1, 2, 3, 4, 5}) == 4)
	println(maxProfit([]int{7, 6, 4, 3, 1}) == 0)
	println(maxProfit([]int{7, 1, 5, 3, 6, 4}) == 7)
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	l_st, res := 0, 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] < prices[i] {
			if prices[i]-prices[l_st] > 0 {
				res += prices[i] - prices[l_st]
			}
			l_st = i + 1
		}
	}

	if l_st < len(prices) {
		res += prices[len(prices)-1] - prices[l_st]
	}

	return res
}
