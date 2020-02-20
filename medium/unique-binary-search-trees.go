package main

import "fmt"

func main() {
	n := 3
	fmt.Println(numTrees(n))
}

/**
二叉搜索树 也就是排序好的数 进行中序遍历
*/
func numTrees(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	F := make([]int, n+1)
	F[0] = 1
	F[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			left := j - 1
			right := i - j
			F[i] += F[left] * F[right]
		}
	}
	return F[n]
}
