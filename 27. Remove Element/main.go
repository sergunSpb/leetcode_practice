package main

func removeElement(nums []int, val int) int {
	i := 0

	for _, w := range nums {
		if w != val {
			nums[i] = w
			i++
		}
	}

	return i
}
