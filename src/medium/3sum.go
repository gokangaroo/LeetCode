package main

import (
	"sort"
	"fmt"
)

func main()  {
	arr := []int{-1,0,1,2,-1,4}
	sort.Ints(arr)
	res := threeSum(arr)
	fmt.Println(res)
}

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