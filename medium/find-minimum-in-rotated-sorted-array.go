package main

import "fmt"

func main() {
	fmt.Println(findMin([]int{2, 1}))
}

func findMin(nums []int) int {
	// 升序排列, 翻转后, 最小值, 与最大值都在中间, 非左即右
	// 1. 折叠期: 如果中>=左, 那就在右边, 否则 在左
	// 2. 直线期: 如果left<right: 直接返回left
	left, right := 0, len(nums)-1
	var mid int
	for left < right {
		if nums[left] < nums[right] {
			break
		}
		mid = (left + right) / 2
		if nums[mid] >= nums[left] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}
