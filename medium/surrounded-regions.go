package main

import "fmt"

func main() {
	var grid = [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
		//{'O', 'O', 'O'},
		//{'O', 'O', 'O'},
		//{'O', 'O', 'O'},
		//{'O', 'X', 'O', 'O', 'O', 'O', 'O', 'O', 'O'},
		//{'O', 'O', 'O', 'X', 'O', 'O', 'O', 'O', 'X'},
		//{'O', 'X', 'O', 'X', 'O', 'O', 'O', 'O', 'X'},
		//{'O', 'O', 'O', 'O', 'X', 'O', 'O', 'O', 'O'},
		//{'X', 'O', 'O', 'O', 'O', 'O', 'O', 'O', 'X'},
		//{'X', 'X', 'O', 'O', 'X', 'O', 'X', 'O', 'X'},
		//{'O', 'O', 'O', 'X', 'O', 'O', 'O', 'O', 'O'},
		//{'O', 'O', 'O', 'X', 'O', 'O', 'O', 'O', 'O'},
		//{'O', 'O', 'O', 'O', 'O', 'X', 'X', 'O', 'O'},
	}
	solveDfs(grid)
	fmt.Print(grid)
}

type BNode struct {
	x int
	y int
}

//BFS 边界出发  排除标记点
func solve(board [][]byte) {
	nr := len(board)
	if nr == 0 {
		return
	}
	nc := len(board[0])
	if nc == 0 {
		return
	}

	nodes := make(chan BNode, nr*nc)
	//边界出发
	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if board[r][c] == 'O' && ((r == 0 || r == nr-1) || (c == 0 || c == nc-1)) {
				nodes <- BNode{r, c}
				//标记当前节点 
				board[r][c] = '#'
				for len(nodes) != 0 {
					current := <-nodes
					board[current.x][current.y] = '#'
					if current.x-1 > 0 && board[current.x-1][current.y] == 'O' {
						board[current.x-1][current.y] = '#'
						nodes <- BNode{current.x - 1, current.y}
					}
					if current.y-1 > 0 && board[current.x][current.y-1] == 'O' {
						board[current.x][current.y-1] = '#'
						nodes <- BNode{current.x, current.y - 1}
					}
					if current.x+1 < nr && board[current.x+1][current.y] == 'O' {
						board[current.x+1][current.y] = '#'
						nodes <- BNode{current.x + 1, current.y}
					}
					if current.y+1 < nc && board[current.x][current.y+1] == 'O' {
						board[current.x ][current.y+1] = '#'
						nodes <- BNode{current.x, current.y + 1}
					}
				}
			}
		}
	}

	//重新渲染 借助中间标识
	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if board[r][c] == 'O' {
				board[r][c] = 'X'
			} else if board[r][c] == '#' {
				board[r][c] = 'O'
			}
		}
	}
}

//DFS 边界出发  排除标记点
func solveDfs(board [][]byte) {
	nr := len(board)
	if nr == 0 {
		return
	}
	nc := len(board[0])
	if nc == 0 {
		return
	}

	//边界出发
	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if board[r][c] == 'O' && ((r == 0 || r == nr-1) || (c == 0 || c == nc-1)) {
				dfs1(board, r, c)
			}
		}
	}
	//重新渲染 借助中间标识
	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if board[r][c] == 'O' {
				board[r][c] = 'X'
			} else if board[r][c] == '#' {
				board[r][c] = 'O'
			}
		}
	}
}

func dfs1(board [][]byte, i, j int) {
	//到底图的边缘
	if board[i][j] == 'X' || board[i][j] == '#' {
		return
	}
	board[i][j] = '#'
	nr := len(board)
	nc := len(board[0])
	if i+1 < nr {
		dfs1(board, i+1, j)
	}
	if i-1 > 0 {
		dfs1(board, i-1, j)
	}
	if j+1 < nc {
		dfs1(board, i, j+1)
	}
	if j-1 > 0 {
		dfs1(board, i, j-1)
	}
}
