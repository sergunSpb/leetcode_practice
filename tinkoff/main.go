package main

import "fmt"

func solve(input []int) []int {
	neg := make([]int, 0)
	pos := make([]int, 0)
	res := make([]int, 0, len(input))

	for _, i := range input {
		if i < 0 {
			neg = append(neg, i*i)
		} else {
			pos = append(pos, i*i)
		}
	}

	i, y := 0, 0
	for i < len(neg) && y < len(pos) {
		if neg[len(neg)-i-1] < pos[y] {
			res = append(res, neg[len(neg)-i-1])
			i++
		} else {
			res = append(res, pos[y])
			y++
		}
	}

	if i < len(neg) {
		for x := len(neg) - i - 1; x >= 0; x-- {
			res = append(res, neg[x])
		}
	} else {
		res = append(res, pos[y:]...)
	}

	return res
}

func main() {
	var input = []int{-5, -3, 0, 1, 2, 4}
	fmt.Print(solve(input))
	input = []int{0, 1, 2, 4}
	fmt.Print(solve(input))
	input = []int{-5, -3}
	fmt.Print(solve(input))
}
