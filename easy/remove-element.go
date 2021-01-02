package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}[:]
	val := 2
	nLen := removeElement(nums, val)
	nums = nums[:nLen]
	fmt.Println(nums)
}

// cnblogs.com/endurance9/p/10347336.html
func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			continue
		}
		i++
	}
	return len(nums)
}
