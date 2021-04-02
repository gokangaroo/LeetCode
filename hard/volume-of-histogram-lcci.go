package main

import "fmt"

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trap(height []int) int {
	/*
	   1. left leftMax
	   2. right rightMax

	   移动max较小的那一个， 当left=right， 移动结束
	*/
	if len(height) == 0 {
		return 0
	}
	left := 0
	right := len(height) - 1
	leftMax := height[left]
	rightMax := height[right]
	r := 0
	for left < right {
		if leftMax <= rightMax {
			left++
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				r += leftMax - height[left]
			}
		} else {
			right--
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				r += rightMax - height[right]
			}
		}
	}
	return r
}
