package main

import "fmt"

func main() {
	fmt.Println(checkPossibility([]int{5, 7, 1, 8}))
}

func checkPossibility(nums []int) bool {
	count := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if count == 0 {
				return false
			}
			if i-1 >= 0 && nums[i-1] > nums[i+1] {
				// 这种情况下只能变i+1而不是i
				nums[i+1] = nums[i]
			} else {
				nums[i] = nums[i+1]
			}
			count--
		}
	}
	return true
}
