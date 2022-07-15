package main

func main() {
	println(maxScoreSightseeingPair([]int{8, 1, 5, 2, 6}) == 11)
	println(maxScoreSightseeingPair([]int{1, 2}) == 2)
	println(maxScoreSightseeingPair([]int{1, 3, 8, 7, 6}) == 14)
	println(maxScoreSightseeingPair([]int{10, 4, 6, 4, 10}) == 16)
}

func maxScoreSightseeingPair(values []int) int {
	return dynPr(values)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dynPr(values []int) int {
	maxI, max := values[0], -1
	for i := 1; i < len(values); i++ {
		max = Max(max, maxI+values[i]-i)
		maxI = Max(maxI, values[i]+i)
	}
	return max
}

func dynPr2(values []int) int {
	i, j, max, maxI := 0, 1, -1, -1

	for j < len(values) {
		if j < len(values)-1 && values[j+1]-(j+1) >= values[j]-j {
			j++
			continue
		}
		for i < j {
			maxI = Max(maxI, values[i]+i)
			i++
		}
		max = Max(max, maxI+values[j]-j)
		j++
	}

	return max
}

func bruteForce(values []int) int {
	max := -1
	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			c := values[i] + values[j] + i - j
			if c > max {
				max = c
			}
		}
	}

	return max
}
