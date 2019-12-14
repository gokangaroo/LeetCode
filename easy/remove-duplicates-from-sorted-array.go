package main

import "fmt"

func main() {
	nums := []int{1, 1, 2, 3}
	len := removeDuplicates(nums)
	for i := 1; i < len; i++ {
		fmt.Println(nums[i])
	}
}

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length == 0 {
		return length
	}
	var x, y int
	for x, y = 0, 1; y < length; y++ {
		if nums[x] != nums[y] {
			x++
			nums[x] = nums[y]
		}
	}
	return x + 1
}
