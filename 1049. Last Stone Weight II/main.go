package main

import "fmt"

func main() {
	fmt.Println(lastStoneWeightII([]int{31, 26, 33, 21, 40}))
}

func lastStoneWeightII(stones []int) int {
	mem := make(map[[2]int]int)
	return helper(mem, stones, 0, 0)
}

func helper(mem map[[2]int]int, stones []int, sum, i int) int {
	if i >= len(stones) {
		return sum
	}
	k := [2]int{sum, i}
	if v, ok := mem[k]; ok {
		return v
	}
	a := helper(mem, stones, sum+stones[i], i+1)
	b := helper(mem, stones, sum-stones[i], i+1)
	if a < 0 || (b >= 0 && b < a) {
		a = b
	}
	mem[k] = a
	return a
}
