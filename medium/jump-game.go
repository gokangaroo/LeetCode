package main

import "fmt"

func main() {
	nums := []int{2, 0, 6, 9, 8, 4, 5}
	res := canJump(nums)
	fmt.Println(res)
}

//动态规划  以子问题解决  会产生耗时问题
//func canJump(nums []int) bool {
//	length := len(nums)
//	if length == 1 {
//		return true
//	}
//	var boolRes = false
//	var i = length - 2
//	//倒数第二个开始递归
//	for ; i >= 0; i-- {
//		//从后开始如果不符合条件 直接退出
//		if nums[i] >= (length - i - 1) {
//			boolRes = boolRes || canJump(nums[:i+1])
//			if i == 0 {
//				return true
//			}
//		}
//	}
//	return boolRes
//}

//贪心最大值即可
func canJump(nums []int) bool {
	max := 0
	for i := 0; i < len(nums); i++ {
		if i > max {
			return false
		}
		if max < i+nums[i] {
			max = i + nums[i]
		}
	}
	return true
}
