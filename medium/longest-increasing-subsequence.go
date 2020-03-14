package main

import "fmt"

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18, 12, 13, 14, 15}
	fmt.Print(lengthOfLIS(nums))
}

//经典最长子序列  二分插入
func lengthOfLIS1(nums []int) int {
	length := len(nums)
	tails := make([]int, length)
	res := 0
	for i := 0; i < length; i++ {
		left := 0
		right := res
		for {
			if left >= right {
				break
			}
			mid := (left + right) / 2
			if tails[mid] < nums[i] {
				left = mid + 1
			} else {
				right = mid
			}
		}
		//替换中间值
		tails[left] = nums[i]
		//最右边插入
		if res == right {
			res++
		}
	}
	return res
}

//动态规划
func lengthOfLIS(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	dp := make([]int, length)
	dp[0] = 1
	maxVal := 1
	for i := 1; i < length; i++ {
		tmp := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				tmp = max(tmp, dp[j])
			}
		}
		dp[i] = tmp + 1
		maxVal = max(dp[i], maxVal)
	}
	return maxVal
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
