package main

import "fmt"

func main() {
	nums:=[]int{1,3,5,6}
	fmt.Println(searchInsert(nums,7))
}
func searchInsert(nums []int, target int) int {
	// 直接遍历
	length:= len(nums)
	if length==0||target<nums[0] {
		return 0
	}
	if target==nums[length-1] {
	 	return length-1
	}
	for j:=0;j< len(nums)-1;j++ {
		if target==nums[j] {
			return j
		}else if target>nums[j]  && target<nums[j+1] {
			return j+1
		}
	}
	return length
	/*二分法:
	left, right := 0, len(nums)-1
	mid := left + (right-left)/2
	for left+1 < right {
		if nums[mid] == target {
			right = mid
		} else if nums[mid] < target {
			left = mid
		} else {
			right = mid
		}
		mid = left + (right-left)/2
	}
	if nums[left] >= target {
		return left
	}
	if nums[right] >= target {
		return right
	}
	return right + 1
	*/
}