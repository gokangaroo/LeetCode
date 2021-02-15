package main

import "fmt"

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
}

func findMaxConsecutiveOnes(nums []int) int {
	tmp := 0
	max := tmp
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			tmp++
			if tmp > max {
				max = tmp
			}
		} else {
			tmp = 0
		}
	}
	return max
}
