package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
}

func twoSum(nums []int, target int) []int {
	mymap := make(map[int]int)
	for i, value := range nums { //value=nums[i]
		mid, ok := mymap[target-value]
		if ok {
			return []int{mid, i}
		}
		mymap[value] = i
	}
	panic("没有两个数相加等于目标数")
}
