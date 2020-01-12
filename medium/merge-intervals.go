package main

import (
	"fmt"
	"sort"
)

func main() {
	var intervals = [][]int{
		{1, 4},
		{2, 3},
		//{2, 6},
		//{15, 18},
	}

	res := merge(intervals)
	fmt.Println(res)
}

func merge(intervals [][]int) [][]int {
	var res [][]int
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for k, _ := range intervals {
		if k == len(intervals)-1 {
			res = append(res, intervals[k])
			break
		}
		if intervals[k][1] >= intervals[k+1][0] {
			//res = append(res, []int{intervals[k][0], intervals[k+1][1]})
			//下个位置替换合并后结果
			if intervals[k][1] >= intervals[k+1][1] {
				intervals[k+1] = []int{intervals[k][0], intervals[k][1]}
			} else {
				intervals[k+1] = []int{intervals[k][0], intervals[k+1][1]}
			}
		} else {
			res = append(res, intervals[k])
		}
	}

	return res
}
