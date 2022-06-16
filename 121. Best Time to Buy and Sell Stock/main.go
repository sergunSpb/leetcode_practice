package main

import "fmt"

func main() {
	fmt.Print(maxProfit([]int{7, 1, 5, 3, 6, 4}) == 5)

}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	max, tMin := 0, prices[0]
	for _, v := range prices {
		if v < tMin {
			tMin = v
		} else if v > tMin && v-tMin > max {
			max = v - tMin
		}
	}

	return max
}
