package main

import "fmt"

func main() {
	fmt.Println(findDisappearedNumbers([]int{3, 3, 2, 2, 1, 1, 3, 4}))
}

func findDisappearedNumbers(nums []int) []int {
	var t int
	for _, el := range nums {
		if el == 0 {
			continue
		}
		t = nums[el-1]
		nums[el-1] = 0
		for t != 0 {
			t, nums[t-1] = nums[t-1], 0
		}
	}

	retval := []int{}
	for in, el := range nums {
		if el != 0 {
			retval = append(retval, in+1)
		}
	}
	return retval
}
