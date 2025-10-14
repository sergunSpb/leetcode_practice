package main

func main() {
	println(longestSubarray([]int{1, 1, 1, 1, 0, 1, 1}))
}

func longestSubarray(nums []int) int {
	l, r, cz := 0, 0, 0
	ans := 0
	wasZero := false

	for _, v := range nums {
		if v == 0 {
			wasZero = true
			cz++
			if cz <= 1 {
				l = r
			} else {
				l = 0
			}
			r = 0

			continue
		}
		cz = 0
		r++
		ans = max(ans, l+r)
	}

	if !wasZero {
		ans--
	}

	return ans
}
