package main

import "fmt"

func main() {
	p := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 1, 1},
	}
	fmt.Println(maxAreaOfIsland(p))
}

func maxAreaOfIsland(grid [][]int) int {
	res := 0
	if len(grid) == 0 {
		return res
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ { //0和1, 遍历完毕的1设置为-1
			if grid[i][j] == 1 {
				tmp := dfs3(i, j, 0, grid)
				if tmp > res {
					res = tmp
				}
			}
		}
	}
	return res
}

func dfs3(y, x, area int, grid [][]int) int { //从岛屿第一个点开始
	// 剪枝
	if grid[y][x] != 1 {
		return area
	}
	// 染色
	area++
	grid[y][x] = -1
	// 深度遍历
	if x < len(grid[0])-1 {
		area = dfs3(y, x+1, area, grid)
	}
	if y < len(grid)-1 {
		area = dfs3(y+1, x, area, grid)
	}
	if x > 0 {
		area = dfs3(y, x-1, area, grid)
	}
	if y > 0 {
		area = dfs3(y-1, x, area, grid)
	}
	return area
}
