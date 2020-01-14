package main

import "fmt"

func main() {
	var m, n = 4, 3
	res := uniquePaths(m, n)
	fmt.Println(res)
}

func uniquePaths(m int, n int) int {
	//借助数组路径实现
	//var arr = make([]int, m)
	//for k := range arr {
	//	arr[k] = 1
	//}
	//
	//for i := 1; i < n; i++ {
	//	for j := 1; j < m; j++ {
	//		arr[j] += arr[j-1]
	//	}
	//}
	//return arr[m-1]

	//求解子问题
	arr := [100][100]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				arr[i][j] = 1
			} else {
				arr[i][j] = arr[i][j-1] + arr[i-1][j]
			}
		}
	}
	return arr[m-1][n-1]
}
