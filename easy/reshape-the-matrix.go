package main

import "fmt"

func main() {
	fmt.Println(matrixReshape([][]int{{1, 2}, {3, 4}}, 1, 4))
}

func matrixReshape(nums [][]int, r int, c int) [][]int {
	/*
	   reshape, 重塑二维数组为r行, c列, 行遍历顺序填充

	   如果可行, 返回new, 否则返回old

	   填充r行, 每行填充c个数字
	*/
	ro := len(nums)
	if ro == 0 {
		return nums
	}
	co := len(nums[0])
	if ro*co != r*c {
		return nums
	}
	num := 0
	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
		for j := 0; j < c; j++ {
			x := num / co
			y := num % co
			num++ // i*c+j
			res[i][j] = nums[x][y]
		}
	}
	return res
}
