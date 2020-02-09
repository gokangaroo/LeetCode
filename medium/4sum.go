package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{0, 0, 0, 0}
	target := 1
	fmt.Println(fourSum(arr, target))
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)

	length := len(nums)

	if length < 4 {
		return res
	}

	for i := 0; i <= length-4; i++ {
		//三个数筛选(光标当前为最小)
		sum := target - nums[i]

		for j := i + 1; j <= length-3; j++ {
			left := j + 1
			right := length - 1

			for left < right {
				tmp := nums[j] + nums[left] + nums[right]
				if tmp == sum {
					arr := make([]int, 0)
					arr = append(arr, nums[i], nums[j], nums[left], nums[right])
					res = append(res, arr)
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if tmp < sum {
					left++
				} else {
					right--
				}
			}
			//second去重
			for nums[j] == nums[j+1] && j+1 < right {
				j++
			}
		}
		//first去重
		for nums[i] == nums[i+1] && i+1 < length-1 {
			i++
		}
	}
	return res
}
