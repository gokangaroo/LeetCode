package main

import "fmt"

func main() {
	var matrix = [][]int{
		//{2},
		//{3},
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	res := spiralOrder(matrix)
	fmt.Println(res)
}

func spiralOrder(matrix [][]int) []int {
	var res []int

	if len(matrix) == 0 {
		return res
	}

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
