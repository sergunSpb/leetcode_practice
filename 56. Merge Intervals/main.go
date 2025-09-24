package main

import "sort"

func main() {

}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}

	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})

	res := make([][]int, 0, 1)
	t := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if t[1] >= intervals[i][0] {
			t[1] = max(intervals[i][1], t[1])
			continue
		}
		res = append(res, t)
		t = intervals[i]
	}

	res = append(res, t)

	return res
}
