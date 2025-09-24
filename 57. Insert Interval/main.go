package main

func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	i, n := 0, len(intervals)

	for i < n && newInterval[0] > intervals[i][1] {
		res = append(res, intervals[i])
		i++
	}

	for i < n && newInterval[1] >= intervals[i][0] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	res = append(res, newInterval)

	if n > i {
		res = append(res, intervals[i:]...)
	}

	return res
}
