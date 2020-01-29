package main

import "fmt"

func main() {
	arr := []int{2,3,4,5,6,7,0,1}
	target := 0
	res := search(arr, target)
	fmt.Println(res)
}

//根据O(log n) 明显为二分查找
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	if len(nums) == 0 {
		return -1
	}
	for {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		//前半部分有序
		if nums[left] <= nums[mid] {
			if target < nums[mid] && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { //后半部分有序
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		if left == mid || right == mid {
			return -1
		}
	}
}
