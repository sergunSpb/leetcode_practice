package main

import (
	"fmt"
)

func main() {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(cost)
	fmt.Println(minCostClimbingStairs(cost))
}

func minCostClimbingStairs(cost []int) int {
	c_h := []int{}
	for in := range cost {
		if in < 2 {
			c_h = append(c_h, 0)
		} else {
			if c_h[in-2]+cost[in-2] < c_h[in-1]+cost[in-1] {
				c_h = append(c_h, c_h[in-2]+cost[in-2])
			} else {
				c_h = append(c_h, c_h[in-1]+cost[in-1])
			}
		}
	}
	var r int
	if c_h[len(c_h)-2]+cost[len(c_h)-2] < c_h[len(c_h)-1]+cost[len(c_h)-1] {
		r = c_h[len(c_h)-2] + cost[len(c_h)-2]
	} else {
		r = c_h[len(c_h)-1] + cost[len(c_h)-1]
	}
	return r
}
