package main

import (
	"fmt"
)

func main() {
	fmt.Println(fairCandySwap(
		[]int{1, 2, 5},
		[]int{2, 4}))
}

func fairCandySwap(A []int, B []int) []int {
	/*
	   入参: A的糖果们大小, B的糖果们大小, 数量与大小都不大
	   操作: 交换一个糖果, 使得总大小一致
	   返回: 某个交换大小序列
	*/
	var a, b int
	var bMap = make(map[int]struct{})
	for _, v := range A {
		a += v
	}
	for _, v := range B {
		b += v
		bMap[v] = struct{}{}
	}
	// 反向指标相等
	gap := (a - b) / 2
	for i := 0; i < len(A); i++ {
		if _, ok := bMap[A[i]-gap]; ok {
			return []int{A[i], A[i] - gap}
		}
	}
	return []int{}
}
