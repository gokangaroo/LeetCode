package main

import "fmt"

func main() {
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 4, 7}))
}

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var (
		l    = 1 //1.记录当前扫描段的len
		maxL = 1 //2.记录maxLen
	)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			l++
			if l > maxL {
				maxL = l
			}
		} else {
			l = 1
		}
	}
	return maxL
}
