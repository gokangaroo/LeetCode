package main

import "fmt"

func main() {
	fmt.Println(flipAndInvertImage([][]int{
		{1, 1, 0},
		{1, 0, 1},
		{0, 0, 0},
	}))
}

func flipAndInvertImage(A [][]int) [][]int {
	/*
	   先逆序每一行, 再替换每一行

	   也可以两个func合并到一起
	*/
	for i := 0; i < len(A); i++ {
		reverse(A[i])
		exchange(A[i])
	}
	return A
}

func reverse(A []int) {
	for i := 0; i < len(A)/2; i++ {
		A[i], A[len(A)-1-i] = A[len(A)-1-i], A[i]
	}
}

func exchange(A []int) {
	for i := 0; i < len(A); i++ {
		A[i] = A[i] ^ 1
	}
}
