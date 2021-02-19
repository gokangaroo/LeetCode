package main

import "fmt"

func main() {
	fmt.Println(longestOnes([]int{0, 0, 0, 1}, 4))
}

func longestOnes(A []int, K int) int {
	/*
	   可以将K个值从0变为1, 返回最长1的个数

	   双指针走法, 当K再次需要时, 需要走左边
	*/
	var (
		max  = 0
		i, j = 0, 0
	)
	for i < len(A) {
		if i > len(A)-K {
			break
		}
		for j < len(A) {
			if A[j] == 0 && K == 0 {
				break
			}
			if A[j] == 0 {
				K--
			}
			if j-i+1 > max {
				max = j - i + 1
			}
			j++
		}
		if A[i] == 0 {
			K++
		}
		i++
	}
	return max
}
