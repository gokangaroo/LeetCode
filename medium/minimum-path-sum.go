package main

import "fmt"

func main() {
	var grid = [][]int{
		{1, 2, 5},
		{3, 2, 1},
	}
	fmt.Println(minPathSum(grid))
}

func minPathSum1(grid [][]int) int {
	nr := len(grid)
	nc := len(grid[0])
	for i := 0; i < nr; i++ {
		for j := 0; j < nc; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] = grid[i][j-1] + grid[i][j]
			} else if j == 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else {
				if grid[i-1][j] < grid[i][j-1] {
					//上面最小
					grid[i][j] = grid[i-1][j] + grid[i][j]
				} else {
					//左边最小
					grid[i][j] = grid[i][j-1] + grid[i][j]
				}
			}
		}
	}
	return grid[nr-1][nc-1]
}

func minPathSumDfs(grid [][]int) int {
	nr := len(grid)
	nc := len(grid[0])

	var path = make([][]int, nr)
	for i := 0; i < nr; i++ {
		for j := 0; j < nc; j++ {
			path[i] = append(path[i], 0)
		}
	}
	path[0][0] = grid[0][0]

	if nr > 1 {
		dfsMinimumPathSum(grid, path, 1, 0)
	}
	if nc > 1 {
		dfsMinimumPathSum(grid, path, 0, 1)
	}
	return path[nr-1][nc-1]
}

//坐标深度优先(有超时风险 并且不需要path记录路径值 由于只能向右或者向下)
func dfsMinimumPathSum(grid, path [][]int, i, j int) {
	nr := len(grid)
	nc := len(grid[0])
	if i == 0 {
		path[i][j] = path[i][j-1] + grid[i][j]
	}
	if j == 0 {
		path[i][j] = path[i-1][j] + grid[i][j]
	}
	if i != 0 && j != 0 {
		if path[i-1][j] < path[i][j-1] {
			//上面最小
			path[i][j] = path[i-1][j] + grid[i][j]
		} else {
			//左边最小
			path[i][j] = path[i][j-1] + grid[i][j]
		}
	}
	if i == nr-1 && j == nc-1 {
		return
	}
	//继续迭代
	if i+1 < nr {
		dfsMinimumPathSum(grid, path, i+1, j)
	}
	if j+1 < nc {
		dfsMinimumPathSum(grid, path, i, j+1)
	}
}

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	if len(grid[0]) == 0 {
		return 0
	}
	res := make([][]int, len(grid))
	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid); j++ { //res[j][i]
			if i == 0 {
				res[j] = make([]int, len(grid[0]))
				if j == 0 {
					res[j][0] = grid[j][0]
				} else {
					res[j][0] = grid[j-1][0] + grid[j][0]
				}
				continue
			}
			if j == 0 {
				res[j][i] = grid[j][i-1] + grid[j][i]
			} else {
				res[j][i] = min(res[j][i-1], res[j-1][i]) + grid[j][i]
			}
		}
	}
	return res[len(grid)-1][len(grid[0])-1]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
