package main

import "fmt"

func main() {
	fmt.Println(generateMatrix(3))
}

func generateMatrix(n int) [][]int {
	/*
	   将1-n^2, 螺旋式放入二维数组
	*/
	r := make([][]int, n)
	for i := 0; i < n; i++ {
		r[i] = make([]int, n)
	}
	var (
		i, j       = 0, 0
		s, x       = 0, n - 1
		t, y       = 1, n - 1
		num, count = 1, n * n
	)
	r[i][j] = num
	num++
	for num <= count {
		for j < x && num <= count {
			j++
			r[i][j] = num
			num++
		}
		x--
		for i < y && num <= count {
			i++
			r[i][j] = num
			num++
		}
		y--
		for j > s && num <= count {
			j--
			r[i][j] = num
			num++
		}
		s++
		for i > t && num <= count {
			i--
			r[i][j] = num
			num++
		}
		t++
	}
	return r
}
