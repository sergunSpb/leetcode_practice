package main

import (
	"fmt"
	"sort"
)

func main() {

}

func threeSumBs(nums []int) [][]int {
	sort.Ints(nums)

	seen := make(map[string]struct{})

	ans := make([][]int, 0)
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		y := i + 1
		for y < len(nums) {
			if ok := bs(nums[y+1:], -nums[i]-nums[y]); ok {
				x := fmt.Sprintf("%d.%d.%d", nums[i], nums[y], -nums[i]-nums[y])
				if _, ok1 := seen[x]; ok1 {
					y++
					continue
				}
				seen[x] = struct{}{}
				ans = append(ans, []int{nums[i], nums[y], -nums[i] - nums[y]})
			}
			y++
		}

	}

	return ans
}

func bs(nums []int, t int) bool {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == t {
			return true
		}
		if nums[mid] < t {
			low = mid + 1
			continue
		}
		if nums[mid] > t {
			high = mid - 1
			continue
		}
	}

	return false
}

// n^2 VVVVVV
func threeSumHashMap(nums []int) [][]int {
	ans := make([][]int, 0)
	sh := make(map[string]struct{})
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		t := twosum(nums, 0-nums[i], i)
		for _, r := range t {
			a := appendSort(nums[i], r)
			x := fmt.Sprintf("%d.%d.%d", a[0], a[1], a[2])
			if _, ok := sh[x]; !ok {
				sh[x] = struct{}{}
				ans = append(ans, a)
			}
		}
	}

	return ans
}

func twosum(nums []int, target int, skip int) [][]int {
	ms := make(map[int]int, len(nums))
	for i := range nums {
		if i == skip {
			continue
		}
		ms[target-nums[i]] = i
	}

	ans := make([][]int, 0)
	for i := range nums {
		if i == skip {
			continue
		}
		if p, ok := ms[nums[i]]; ok && p != i {
			ans = append(ans, sortTwo(nums[p], nums[i]))
		}
	}

	return ans
}

func sortTwo(i, p int) []int {
	if i < p {
		return []int{i, p}
	}

	return []int{p, i}
}

func appendSort(i int, to []int) []int {
	if i < to[0] {
		return []int{i, to[0], to[1]}
	}

	if i > to[1] {
		return append(to, i)
	}

	return []int{to[0], i, to[1]}
}
