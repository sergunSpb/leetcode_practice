package main

func main() {
	println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}) == 49)
}

func maxArea(height []int) int {
	i, j := 0, len(height)-1

	ms := -1

	for i < j {
		cs := min(height[i], height[j]) * (j - i)
		if cs > ms {
			ms = cs
		}

		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}

	return ms
}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}
