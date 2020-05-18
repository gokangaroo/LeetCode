package main

import "fmt"

func main() {
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
}

func findPeakElement(nums []int) int {
	// 0. 时间复杂度要求,所以需要是二分法
	// 此外只要找到一个峰值, 所以可以进行部分判断.
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[mid+1] { // 1. 右边在下坡, 左边肯定有一个峰值
			right = mid
		} else { //2. 左边在上坡, 右边肯定有一个峰值
			left = mid + 1
		}
	}
	return left
}
