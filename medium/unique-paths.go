package main

import "fmt"

func main() {
	var m, n = 4, 3
	res := uniquePaths(m, n)
	fmt.Println(res)
}

func uniquePaths(m int, n int) int {
	var arr = make([]int, m)
	for k := range arr {
		arr[k] = 1
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			arr[j] += arr[j-1]
		}
	}
	return arr[m-1]
}
