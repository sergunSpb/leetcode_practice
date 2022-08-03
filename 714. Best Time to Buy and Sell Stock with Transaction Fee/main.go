package main

func main() {

	println(maxProfit([]int{1, 3, 2, 8, 4, 9}, 2) == 8)
	println(maxProfit([]int{1, 3, 7, 5, 10, 3}, 3) == 6)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int, fee int) int {

	pr_baught := -prices[0]
	var sold, baught, pr_sold int

	for i := 1; i < len(prices); i++ {
		sold = max(pr_baught+prices[i]-fee, pr_sold)
		baught = max(pr_baught, pr_sold-prices[i])
		pr_baught, pr_sold = baught, sold
	}

	return sold
}
