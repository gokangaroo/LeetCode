package main

import "fmt"

// https://leetcode-cn.com/problems/next-greater-element-ii/solution/jian-dan-bian-li-ordan-diao-zhan-by-zhua-d866/
func main() {
	fmt.Println(nextGreaterElements([]int{1, 2, 3, 2, 1}))
	fmt.Println(nextGreaterElements1([]int{1, 2, 3, 2, 1}))
}

func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				res[i] = nums[j]
				goto end
			}
		}
		for j := 0; j < i; j++ {
			if nums[j] > nums[i] {
				res[i] = nums[j]
				goto end
			}
		}
		res[i] = -1
	end:
	}
	return res
}

func nextGreaterElements1(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	for index := range res {
		res[index] = -1
	}
	var stack = make([]int, 0)                  // 单调非递增栈, 存下标
	for index := 0; index < length*2; index++ { // 假装拼接了前n-1个元素
		i := index % length // 取模得到真实idx
		top := len(stack) - 1
		for ; top >= 0 && nums[i] > nums[stack[top]]; top-- { // 大于栈顶元素, 则开始逐个弹出下标, 设置res[top]
			res[stack[top]] = nums[i]
		}
		stack = stack[:top+1]
		if index < length { // 入栈
			stack = append(stack, index)
		}
	}
	return res
}
