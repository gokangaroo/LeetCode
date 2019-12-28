package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-3, -2, -5, 3, -4}
	target := -1
	sum := threeSumClosest(nums, target)
	fmt.Println("sum = ", sum)
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	minSum := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		//首尾双向指针
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			//不断更新minSum
			if absInt(sum-target) < absInt(minSum-target) {
				minSum = sum
			}
			if sum > target {
				right--
			} else if sum < target {
				left++
			} else {
				//都相等了 肯定是target
				return sum
			}
		}
	}
	return minSum
}

func absInt(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
