package main

import "fmt"

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
}

func pivotIndex(nums []int) int {
	if len(nums) <= 1 {
		return -1
	}
	sumL, sumR := 0, 0
	for i := 0; i < len(nums); i++ {
		sumR += nums[i]
	}
	for i := 0; i < len(nums); i++ {
		if i-1 >= 0 {
			sumL += nums[i-1]
		}
		sumR -= nums[i]
		if sumL == sumR {
			return i
		}
	}
	return -1
}
