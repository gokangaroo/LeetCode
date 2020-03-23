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

func fourSum2(nums []int, target int) [][]int {
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
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] { //到重复的最后一个
						left++
					}
					for left < right && nums[right] == nums[right-1] { //到重复的第一个
						right--
					}
					left++ //去重
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

// 暴力就是美, 双指针主要是三数之和的升级版本,外部套了一个循环.
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)

	length := len(nums)
	if length < 4 {
		return res
	}

	for i := 0; i < length-3; i++ {
		for m := i + 1; m < length-2; m++ {
			for n := m + 1; n < length-1; n++ {
				for x := n + 1; x < length; x++ {
					if nums[i]+nums[m]+nums[n]+nums[x] == target {
						res = append(res, []int{nums[i], nums[m], nums[n], nums[x]})
					}
					for x != length-1 && nums[x] == nums[x+1] { //重复的最后一个
						x++
					}
				}
				for n != length-1 && nums[n] == nums[n+1] { //重复的最后一个
					n++
				}
			}
			for m != length-1 && nums[m] == nums[m+1] { //重复的最后一个
				m++
			}
		}
		for i != length-1 && nums[i] == nums[i+1] { //重复的最后一个
			i++
		}
	}
	return res
}
