package main

import "fmt"

func main() {
	t := [][]int{
		{1, 2, 3,},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(findDiagonalOrder(t))
}

func findDiagonalOrder(matrix [][]int) []int {
	res := make([]int, 0)
	if len(matrix) == 0 {
		return res
	}

	right := true //默认第一个箭头向右边
	for x, y := 0, 0; ; {
		res = append(res, matrix[y][x])
		if x == len(matrix[0])-1 && y == len(matrix)-1 {
			break
		}
		if right { //右边箭头
			if y > 0 && x != len(matrix[0])-1 {
				y--
				x++
				continue
			}
			if x == len(matrix[0])-1 {
				right = false
				y++
			}
			if y == 0 {
				right = false
				x++
			}
		} else {
			if x > 0 && y != len(matrix)-1 {
				x--
				y++
				continue
			}
			if y == len(matrix)-1 {
				right = true
				x++
			}
			if x == 0 {
				right = true
				y++
			}
		}
	}
	return res
}
