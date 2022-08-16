package main

import "fmt"

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}) == 6)
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}) == 9)
}

//dynamic approach
func trap(height []int) int {
	sum := 0
	l_max := make([]int, len(height))
	r_max := make([]int, len(height))
	l_max[0] = height[0]
	for i := 1; i < len(height); i++ {
		l_max[i] = max(l_max[i-1], height[i])

	}
	r_max[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		r_max[i] = max(r_max[i+1], height[i])
	}

	for i := 1; i < len(height)-1; i++ {
		h1 := height[i]
		sum += min(l_max[i], r_max[i]) - h1
	}

	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func trapBruteForceN2(height []int) int {
	sum := 0
	for i1 := 1; i1 < len(height)-1; i1++ {
		h1 := height[i1]
		l_max := -1
		for j := i1; j >= 0; j-- {
			l_max = max(l_max, height[j])

		}
		r_max := -1
		for j := i1; j < len(height); j++ {
			r_max = max(r_max, height[j])

		}

		sum += min(l_max, r_max) - h1
	}

	return sum
}

func trapN3(height []int) int {
	sum := 0
	for i1 := 0; i1 < len(height)-1; i1++ {
		h1 := height[i1]
		if h1 <= height[i1+1] {
			continue
		}
		for depth := 0; depth < h1-height[i1+1]; depth++ {
			for i2 := i1 + 1; i2 < len(height); i2++ {
				if height[i2] >= h1-depth {
					sum += i2 - i1 - 1
					break
				}
			}
		}
	}

	return sum
}
