package main

import "fmt"

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
}

func summaryRanges(nums []int) (res []string) {
	// 1. 无重复, 有序, nums
	// 2. 最小有序区间array
	/*
	   "a->b" ，如果 a != b
	   "a" ，如果 a == b
	*/
	if len(nums) == 0 {
		return
	}
	for a := 0; a < len(nums); a++ {
		var b int
		for b = a; b < len(nums); b++ {
			if b == a {
				if b == len(nums)-1 {
					break
				}
				continue
			}
			if nums[b]-nums[b-1] == 1 {
				if b == len(nums)-1 {
					break
				}
				continue
			}
			// nums[b]-nums[a]!=1, 两个区间, 由a来触发
			b = b - 1
			break
		}
		if a == b {
			res = append(res, fmt.Sprintf("%d", nums[a]))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", nums[a], nums[b]))
		}
		a = b
	}
	return
}
