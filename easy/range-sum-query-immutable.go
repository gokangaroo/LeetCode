package main

import "fmt"

func main() {
	constructor := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println((&constructor).SumRange(2, 5))
}

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	// 在0位填充0, 后面本来是sums[j]-sums[i]+nums[i]可以改为sums[j+1]-sums[i]
	sums := make([]int, len(nums)+1)
	for i, v := range nums {
		sums[i+1] = sums[i] + v
	}
	return NumArray{sums}
}

func (na *NumArray) SumRange(i, j int) int {
	return na.sums[j+1] - na.sums[i]
}
