package main

import "fmt"

func main() {
	ans := permute([]int{1, 2, 3})
	fmt.Println(ans)
	fmt.Println(len(ans))
}

func permute(nums []int) [][]int {
	retval := [][]int{}
	if len(nums) == 1 {
		return [][]int{nums}
	}
	for i, num := range nums {
		r := make([]int, 0, len(nums)-1)
		r = append(r, nums[:i]...)
		r = append(r, nums[i+1:]...)
		for _, sl := range permute(r) {
			retval = append(retval, append([]int{num}, sl...))
		}
	}
	return retval
}
