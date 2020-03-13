package main

import "fmt"

func main() {
	var nums = []int{6, 5, 5}
	fmt.Print(majorityElement(nums))
}

func majorityElement(nums []int) int {
	res, count := 0, 0
	hash := make(map[int]int)
	for _, v := range nums {
		hash[v]++
		if hash[v] >= count {
			res = v
			count = hash[v]
		}
	}
	return res
}
