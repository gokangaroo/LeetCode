package main

import "fmt"

func main() {
	var matrix = [][]int{
		//{2},
		//{3},
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	res := spiralOrder1(matrix)
	fmt.Println(res)
}

func spiralOrder1(matrix [][]int) (res []int) {
	/*
	   1 <= m, n <= 10
	   // j++,i++,j--,i--
	*/
	var (
		i, j  int
		s, m  = 0, len(matrix[0]) - 1
		t, n  = 1, len(matrix) - 1
		count = (m + 1) * (n + 1)
	)
	res = append(res, matrix[0][0])
	count--
	for count > 0 {
		for j < m && count > 0 {
			j++
			res = append(res, matrix[i][j])
			count--
		}
		m--
		for i < n && count > 0 {
			i++
			res = append(res, matrix[i][j])
			count--
		}
		n--
		for j > s && count > 0 {
			j--
			res = append(res, matrix[i][j])
			count--
		}
		s++
		for i > t && count > 0 {
			i--
			res = append(res, matrix[i][j])
			count--
		}
		t++
	}
	return
}

func spiralOrder(matrix [][]int) []int {
	var res []int

	//初始化坐标和临界点
	var x, y, xStart, yStart = 0, 0, 0, 0
	xEnd := len(matrix[0]) - 1
	yEnd := len(matrix) - 1

	//顺时针规则
	for {
		//向右
		if x == yStart && y <= xEnd {
			for {
				res = append(res, matrix[x][y])
				if y == xEnd {
					x++
					break
				}
				y++
			}
			//转向
			yStart++
			if xStart > xEnd || yStart > yEnd {
				break
			}
		}
		//向下
		if y == xEnd && x <= yEnd {
			for {
				res = append(res, matrix[x][y])
				if x == yEnd {
					y--
					break
				}
				x++
			}
			xEnd--
			if xStart > xEnd || yStart > yEnd {
				break
			}
		}
		//向左
		if x == yEnd && y >= xStart {
			for {
				res = append(res, matrix[x][y])
				if y == xStart {
					x--
					break
				}
				y--
			}
			yEnd--
			if xStart > xEnd || yStart > yEnd {
				break
			}
		}
		//向上
		if y == xStart && x >= yStart {
			for {
				res = append(res, matrix[x][y])
				if x == yStart {
					y++
					break
				}
				x--
			}
			xStart++
			if xStart > xEnd || yStart > yEnd {
				break
			}
		}
	}
	return res
}
