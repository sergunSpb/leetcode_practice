package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(sortArray([]int{4, 1, 2, 6, 20, 0, -2, 4}))
}

func b(i int, sep string) string {
	b:= strings.Builder{}
	for i > 1000 {
		s = i % 1000
	}
}

func sortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	p := len(nums) - 1
	var l, r []int
	for i := 0; i < len(nums); i++ {
		if i == p {
			continue
		}

		if nums[i] < nums[p] {
			l = append(l, nums[i])
		} else {
			r = append(r, nums[i])
		}
	}

	return append(append(sortArray(l), nums[p]), sortArray(r)...)
}

// func part(arr []int, pivot int) (l []int, r []int) {
// 	for _, i := range arr {
// 		if i < pivot {
// 			l = append(l, i)
// 		} else {
// 			r = append(r, i)
// 		}
// 	}

// 	return r, l
// }

//MergeSort
// func sortArray(nums []int) []int {
// 	if len(nums) <= 1 {
// 		return nums
// 	}
// 	mid := len(nums) / 2
// 	return merge(sortArray(nums[:mid]), sortArray(nums[mid:]))
// }

// func merge(a1 []int, a2 []int) []int {
// 	p1, p2 := 0, 0
// 	l1, l2 := len(a1), len(a2)
// 	var res []int
// 	for p1 < l1 && p2 < l2 {
// 		if a1[p1] < a2[p2] {
// 			res = append(res, a1[p1])
// 			p1++
// 		} else {
// 			res = append(res, a2[p2])
// 			p2++
// 		}
// 	}
// 	res = append(res, a1[p1:]...)
// 	return append(res, a2[p2:]...)
// }


[
  "CategoryID: <nil pointer> != int64\internal/aggregator/cache/get.gonLocationID: <nil pointer> != int64\nUserID: <nil pointer> != int64\nPrice.value: <nil pointer> != int64\nPrice.filled: false != true\nTitle: <nil pointer> != string\nStatusID: <nil pointer> != int64\nFinishTime.value: <nil pointer> != float64\nFinishTime.filled: false != true\nMicrocatID.value: <nil pointer> != int64"
]