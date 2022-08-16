package main

func main() {
	println(numberOfArithmeticSlices([]int{1, 2, 3, 4}) == 3)
	println(numberOfArithmeticSlices([]int{1}) == 0)
}

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	var ans, lSeq int
	for i := 2; i < len(nums); i++ {
		if nums[i-1]-nums[i-2] == nums[i]-nums[i-1] {
			ans += lSeq + 1
			lSeq++
		} else {
			lSeq = 0
		}
	}

	return ans
}
