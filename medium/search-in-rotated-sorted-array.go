package main

import "fmt"

func main() {
	arr := []int{2, 3, 4, 5, 6, 7, 0, 1}
	target := 0
	res := search(arr, target)
	fmt.Println(res)
}

//根据O(log n) 明显为二分查找
func search2(nums []int, target int) int {
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

func search(nums []int, target int) int {
	index := -1
	if len(nums) == 0 {
		return index
	}
	// 二分法==>特殊在, 一边符合升序, 一边不符合
	i, j := 0, len(nums)-1
	mid := j + (i-j)/2
	for i <= j {
		if nums[mid] == target {
			index = mid
			break
		}
		// 如果左边升序, 右边不升, 那么右边有个低谷, 而且最大值小于左边最小值
		// 可以根据mid值直接判断在哪边(画图即可)
		// 这道题需要注意到注意到y轴上的数值差距
		if nums[mid] > nums[i] { //左边升序
			if nums[mid] > target && target >= nums[i] { //&&后面就是y轴数值差距的条件判断
				j = mid - 1
			} else {
				i = mid + 1
			}
		} else { //右边升序  nums[mid]<=nums[j]<=nums[i]
			if nums[mid] < target && target <= nums[j] {
				i = mid + 1
			} else {
				j = mid - 1
			}
		}
		mid = j + (i-j)/2
	}
	return index
}
