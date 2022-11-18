package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] <= intervals[j][0]
		} else {
			return intervals[i][1] > intervals[j][1]
		}
	})

	ans := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		a, b := ans[len(ans)-1], intervals[i]
		if a[1] >= b[0] {
			if a[1] < b[1] {
				ans[len(ans)-1] = []int{a[0], b[1]}
			}
		} else {
			ans = append(ans, b)
		}
	}

	return ans
}

func main() {
	res := merge([][]int{{0, 1}, {3, 4}, {2, 6}, {2, 18}, {5, 10}, {19, 19}})
	fmt.Println(res)
}
