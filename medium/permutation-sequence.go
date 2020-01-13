package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, k = 4, 9
	res := getPermutation(n, k)
	fmt.Println(res)
}

func getPermutation(n int, k int) string {
	//初始化每个位置的是否使用标记
	var visited = make([]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = false
	}
	return recursive(n, factorial(n-1), k, visited)
}

func recursive(n, f, k int, visited []bool) string {
	//每组数据得到第几组的第几个排列
	offset := k % f
	groupIndex := 0
	if offset > 0 {
		groupIndex = k/f + 1
	} else {
		groupIndex = k / f
	}
	//在位置切片上 标记此轮数字
	i := 0
	for ; i < len(visited) && groupIndex > 0; i++ {
		if !visited[i] {
			groupIndex--
		}
	}
	visited[i-1] = true //标记该轮使用的数字
	if n-1 > 0 {
		//如果正好整除  则保持f
		if offset == 0 {
			offset = f
		}
		return strconv.Itoa(i) + recursive(n-1, f/(n-1), offset, visited)
	} else {
		return strconv.Itoa(i)
	}
}

//阶乘结果
func factorial(n int) int {
	res := 1
	for i := n; i > 0; i-- {
		res *= i
	}
	return res
}
