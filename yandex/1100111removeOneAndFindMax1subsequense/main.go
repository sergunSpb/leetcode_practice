package main

func main() {
	println(getMax([]int{0, 1}) == 1)
	println(getMax([]int{1, 0}) == 1)
	println(getMax([]int{1, 1}) == 1)
	println(getMax([]int{0, 0}) == 0)
	println(getMax([]int{1, 1, 1, 1, 0, 1, 1, 0, 0, 1, 1, 1, 1, 1}) == 6)
}

func getMax(s []int) int {
	var max, cur, prev int
	for i := 0; i < len(s); i++ {
		if s[i] == 1 {
			cur++
			if cur+prev > max {
				max = cur + prev
			}
		} else {
			if i == len(s)-1 {
				break
			}
			if s[i+1] == 0 {
				prev = 0
			} else {
				prev = cur
			}

			cur = 0
		}
	}

	if max == len(s) {
		max--
	}

	return max
}
