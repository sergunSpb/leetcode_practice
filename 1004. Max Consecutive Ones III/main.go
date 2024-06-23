package main

func longestOnes(nums []int, k int) int {
	var l, r, ans, curZeros int

	for _, c := range nums {
		if c == 0 {
			curZeros++
		}

		for curZeros > k {
			if nums[l] == 0 {
				curZeros--
			}
			l++
		}

		r++
		ans = max(ans, r-l)
	}

	return ans
}

func max(i, y int) int {
	if i > y {
		return i
	}

	return y
}
