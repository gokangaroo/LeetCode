package main

import "fmt"

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
}

func findMaxAverage(nums []int, k int) float64 {
	/*
	   滑动窗口的总和的最大值
	*/
	windows := 0
	for i := 0; i < k; i++ {
		windows = windows + nums[i]
	}
	max := windows
	for i := k; i < len(nums); i++ {
		windows = windows + nums[i] - nums[i-k]
		if windows > max {
			max = windows
		}
	}
	return float64(max) / float64(k)
}
