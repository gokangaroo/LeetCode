package main

import "fmt"

func main() {
	fmt.Println(transpose([][]int{
		{1, 2, 3},
		{4, 5, 6},
	}))
}

func transpose(matrix [][]int) [][]int {
	/*
	   斜对角翻转 => res[j][i]=matrix[i][j]
	*/
	m := len(matrix)
	n := len(matrix[0])
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[j][i] = matrix[i][j]
		}
	}

	return res
}
