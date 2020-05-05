package main

import "fmt"

func main() {
	nums := []int{2,1}
	k := 1
	t := 1
	fmt.Println(containsNearbyAlmostDuplicate(nums, k, t))
}

// 先使用n^2解法看看
// 存在i,j满足条件
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for i := 0; i < len(nums); i++ { //需要考虑k>len的场景
		for j := max(i-k, 0); j < i; j++ {
			if !(nums[j]-nums[i] < - 1*t || nums[j]-nums[i] > t) {
				return true
			}
		}
	}
	return false
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
