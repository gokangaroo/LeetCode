package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{-1, 0, 1, 2, -1, 4}
	//arr := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	//arr := []int{0, 0, 0, 0}
	sort.Ints(arr)
	res := threeSum(arr)
	fmt.Println(res)
}

/**
huija
干, 这道题好多重复的数, 无语
 */
//func threeSum(nums []int) [][]int {
//	sort.Ints(nums)
//	res := make([][]int, 0)
//	mymap := make(map[int]int)
//	arrmap := make(map[int]int)
//	for i := 0; i < len(nums); i++ {
//		for j := i + 1; j < len(nums); j++ {
//			_, ok := mymap[0-nums[i]-nums[j]]
//			if ok {
//				arr := make([]int, 0)
//				arr = append(arr, 0-nums[i]-nums[j], nums[i], nums[j])
//				v, ok := arrmap[arr[0]]
//				if !ok || v != arr[1] {
//					res = append(res, arr)
//					arrmap[arr[0]] = arr[1]
//				}
//			}
//		}
//		mymap[nums[i]] = i
//	}
//	return res
//}

/**
GeekGhc
 */
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int,0)

	if len(nums)<3{
		return res
	}
	for i:=0;i<len(nums);i++{
		left := i+1
		right := len(nums)-1
		//去除重复值   以i为第一个元素
		if i>0 && nums[i]==nums[i-1]{
			continue
		}
		//如果最小的元素大于0  跳出
		if nums[i]>0 {
			break
		}
		for left < right{
			sum := nums[i]+nums[left]+nums[right]
			if sum == 0{
				arr := make([]int,0)
				arr = append(arr,nums[i],nums[left],nums[right])
				res = append(res,arr)
				for left<right && nums[left]==nums[left+1]{
					left++
				}
				for left<right && nums[right]==nums[right-1]{
					right--
				}
				left++
				right--
			} else if sum<0{
				left++
			} else {
				right--
			}
		}
	}
	return res
}
