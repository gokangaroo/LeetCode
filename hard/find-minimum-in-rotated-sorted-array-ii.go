package main

import "fmt"

func main() {
	fmt.Println(findMin([]int{2, 2, 2, 0, 1}))
}

func findMin(nums []int) int {
	// 还是二分法,非左就是右边
	// 1. 折叠期: 如果中>=左, 那就在右边, 否则 在左
	// 2. 直线期: 如果left<right||left==right(left不是开头): 直接返回left
	left, right := 0, len(nums)-1
	var mid int
	for left < right {
		// 2.直线期
		if nums[left] < nums[right] {
			break
		}
		if nums[left] == nums[right] && left != 0 {
			break
		}
		// 保证left>right且mid!=right等来避免各种特殊场景
		for left < right-1 {
			if nums[left] == nums[left+1] {
				left++
			} else if nums[right] == nums[right-1] {
				right--
			} else {
				break
			}
		}
		// 1.折叠期
		mid = (left + right) / 2
		if nums[mid] >= nums[left] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}
