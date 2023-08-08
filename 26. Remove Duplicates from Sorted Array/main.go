package main

import "fmt"

func main() {
	s := []int{1, 1, 2}
	fmt.Printf("%v \n", s[0:removeDuplicates(s)])

	s = []int{1, 2, 3}
	fmt.Printf("%v \n", s[0:removeDuplicates(s)])

	s = []int{1, 1}
	fmt.Printf("%v \n", s[0:removeDuplicates(s)])

	s = []int{1, 1, 2, 2, 2, 2, 4, 4, 4, 5, 6}
	fmt.Printf("%v \n", s[0:removeDuplicates(s)])
	fmt.Printf("%v \n", s[0:removeDuplicates(s)])
}

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	i := 0
	for i < len(nums)-2 {
		if nums[i+1] == nums[i] {
			j := i + 2
			for j < len(nums) {
				if nums[j] != nums[i] {
					nums[i+1] = nums[j]
					i++
				}
				j++
			}
			return i + 1
		}

		i++
	}

	if nums[i+1] != nums[i] {
		i++
	}

	return i + 1
}
