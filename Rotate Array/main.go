package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	rotate(s, 7)
	fmt.Println(s)
}

func rotate_brute_force(nums []int, k int) {
	l := len(nums)
	k = k % l
	for i := 0; i < k; i++ {
		t := nums[0]
		for y := 0; y < l; y++ {
			i := (y + 1) % l
			nums[i], t = t, nums[i]
		}
	}
}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start, end = start+1, end-1
	}
}

func rotate(nums []int, k int) {
	l := len(nums)
	k = k % l
	reverse(nums, 0, l-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, l-1)
}
