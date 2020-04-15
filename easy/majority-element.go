package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
}

func majorityElement(nums []int) int {
	count := make(map[int]int)
	for _, v := range nums {
		count[v]++
	}
	for k, v := range count {
		if v > len(nums)/2 {
			return k
		}
	}
	return 0
}
