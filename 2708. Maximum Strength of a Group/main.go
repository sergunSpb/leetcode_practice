package main

func main() {
	println(maxStrength([]int{-1, 0, 134}))
}

func maxStrength(nums []int) int64 {
	var res int64 = 1

	neg := 0
	maxNeg := -10
	largest := -10

	for _, i := range nums {
		if i != 0 {
			res *= int64(i)
		}

		if i < 0 {
			neg++
			maxNeg = max(maxNeg, i)
		}
		largest = max(largest, i)
	}

	if neg < 2 && largest == 0 {
		return 0
	}
	if neg == 1 && largest < 0 {
		return int64(maxNeg)
	}

	if neg%2 == 1 {
		res = res / int64(maxNeg)
	}

	return res
}
