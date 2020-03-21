package main

import "fmt"

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Print(lengthOfLIS3(nums))
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

func lengthOfLIS2(nums []int) int {
	lenght := len(nums)
	if lenght == 0 || lenght == 1 {
		return lenght
	}
	max := 1
	res := make([]int, lenght)
	for i := 0; i < lenght; i++ { // 以num[i]结尾(必定包含)的最长升序
		res[i] = 1 //先都默认给1
		tmp := 0
		for j := 0; j < i; j++ { //j不能等于i
			if nums[j] < nums[i] && res[j] > tmp { //必须包含我i
				tmp = res[j]
			}
		}
		res[i] = tmp + 1
		// 比较
		if res[i] > max {
			max = res[i]
		}
	}
	return max
}

func lengthOfLIS3(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	tails := make([]int, length)
	tcap := 0 //已经装填的
	tails[tcap] = nums[0]
	tcap = 1
	for i := 1; i < length; i++ {
		// 二分主要找到第一个大于nums[i]的数字替换掉
		// 如果是最后一个, 刚好可以降低门槛
		// 如果是中间, 也不影响结果
		if nums[i] > tails[tcap-1] {
			tails[tcap] = nums[i] //放到最后面
			tcap++
			continue
		}
		// 二分查找的目的, 是替换tails第一个比nums[i]大的数
		// 因为即使比tails的最后一个要小, 并不能直接替换
		// 只有当前面都没有比nums[i]大的时候, 才可以替换tails的结尾
		// 所以取通用的方法, 也就是二分找这个位置
		// 需要利用left
		left := 0
		right := tcap
		for left <= right {
			mid := left + (right-left)/2
			if tails[mid] < nums[i] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		// 不是二分查找, 只能找最后的left作为插入值
		tails[left] = nums[i]
	}
	return tcap
}
