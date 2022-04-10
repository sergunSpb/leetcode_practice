package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(deleteAndEarn([]int{3, 4, 2}) == 6)
	fmt.Println(deleteAndEarn([]int{2, 2, 3, 3, 3, 4}) == 9)
}
func deleteAndEarn(nums []int) int {
	// DP approach , two DP solutions : topBottom and bottomUp
	return declareAndEarnTopBottom(nums)
}

func deleteAndEarnBottomUp(nums []int) int {
	sums := make(map[int]int, 10001)

	max := 0.0
	for _, el := range nums {
		if _, ok := sums[el]; ok {
			sums[el] += el
		} else {
			sums[el] = el
		}

		max = math.Max(float64(el), max)
	}
	maxInt := int(max)
	counted := make([]int, maxInt+1)
	counted[1] = sums[1]

	for el := 2; el <= maxInt; el++ {
		counted[el] = int(math.Max(float64(counted[el-2]+sums[el]), float64(counted[el-1])))
	}

	return counted[maxInt]
}

func declareAndEarnTopBottom(nums []int) int {
	sums := make(map[int]int, 10001)

	max := 0.0
	for _, el := range nums {
		if _, ok := sums[el]; ok {
			sums[el] += el
		} else {
			sums[el] = el
		}

		max = math.Max(float64(el), max)
	}
	maxInt := int(max)

	cache := make([]interface{}, maxInt+1)
	var h func(int) float64
	h = func(point int) float64 {
		if point == 0 {
			return 0
		}
		if point == 1 {
			return float64(sums[1])
		}
		if v := cache[point]; v != nil {
			return v.(float64)
		} else {
			cache[point] = math.Max(h(point-2)+float64(sums[point]), h(point-1))
			return cache[point].(float64)
		}
	}

	return int(h(maxInt))
}
