package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, -2, -1}
	//arr := []int{-1, 0, 1, 2, -1, 4}
	//arr := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	//arr := []int{0, 0, 0, 0}
	sort.Ints(arr)
	res := threeSum(arr)
	fmt.Println(res)
}

/**
huija
*/
func threeSum(nums []int) [][]int {
	// 1.排序
	sort.Ints(nums)
	res := make([][]int, 0)

	if len(nums) < 3 {
		return res
	}

	// left
	for left := 0; left < len(nums)-2; left++ {
		mid := left + 1
		right := len(nums) - 1
		if nums[left] > 0 {
			return res
		}
		for mid < right { //多种情况
			sum := nums[left] + nums[mid] + nums[right]
			if sum > 0 {
				right--
			} else if sum < 0 {
				mid++
			} else {
				res = append(res, []int{nums[left], nums[mid], nums[right]})
				// 找到其他情况
				for mid < right && nums[mid+1] == nums[mid] {
					mid++
				}
				for mid < right && nums[right-1] == nums[right] {
					right--
				}
				mid++
				right--
			}
		}
		// 开始移动到重复的最后-->后会++就是非重复的第一个
		for left < len(nums)-1 && nums[left+1] == nums[left] {
			left++
		}
	}
	return res
}

/**
GeekGhc
*/
func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)

	if len(nums) < 3 {
		return res
	}
	for i := 0; i < len(nums); i++ {
		left := i + 1
		right := len(nums) - 1
		//去除重复值   以i为第一个元素
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		//如果最小的元素大于0  跳出
		if nums[i] > 0 {
			break
		}
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				arr := make([]int, 0)
				arr = append(arr, nums[i], nums[left], nums[right])
				res = append(res, arr)
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}
