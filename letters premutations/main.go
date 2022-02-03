package main

import "fmt"

func main() {
	ans := letterCombinations("")
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

var dLetersMaps map[string][]string = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	if len(digits) == 1 {
		return dLetersMaps[digits]
	}
	retval := []string{}

	for _, s := range dLetersMaps[string(digits[0])] {
		n_s := string(digits[1:])
		for _, st := range letterCombinations(n_s) {
			retval = append(retval, s+st)
		}
		retval = append(retval)
	}

	return retval
}
