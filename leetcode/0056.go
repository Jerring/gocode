package leetcode

import (
	"slices"
)

func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a []int, b []int) int {
		return a[0] - b[0]
	})
	res := make([][]int, 0)
	for _, interval := range intervals {
		if len(res) == 0 || interval[0] > res[len(res)-1][1] {
			res = append(res, []int{interval[0], interval[1]})
		} else {
			tail := res[len(res)-1]
			tail[1] = max(tail[1], interval[1])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
