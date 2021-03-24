package main

import "fmt"

func main() {
	fmt.Println(find132pattern([]int{1, 4, 0, -1, -2, -3, -1, -2}))
}

func find132pattern(nums []int) bool {
	/*
	   假设前两个数字是tail, head, 维护这两个数字

	   题解是单调栈..还没理解 φ(*￣0￣)
	*/
	if len(nums) < 3 {
		return false
	}
	tail := nums[0]
	tmp := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] <= tail {
			tail = nums[i]
			tmp = i
		} else {
			tmp = i
			break
		}
	}
	head := nums[tmp]
	for i := tmp; i < len(nums); i++ {
		if nums[i] >= head {
			head = nums[i]
			tmp = i
		} else {
			tmp = i
			break
		}
	}
	for i := tmp; i < len(nums); i++ {
		if nums[i] > tail && nums[i] < head {
			return true
		}
	}
	// 到这说明不行, 从后面再进行判断
	return find132pattern(nums[tmp:])
}
