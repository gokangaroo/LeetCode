package main

import "fmt"

func main() {
	var grid = [][]byte{
		{'1', '1', '1', '1', '1', '1', '1'},
		{'0', '0', '0', '0', '0', '0', '1'},
		{'1', '1', '1', '1', '1', '0', '1'},
		{'1', '0', '0', '0', '1', '0', '1'},
		{'1', '0', '1', '0', '1', '0', '1'},
		{'1', '0', '1', '1', '1', '0', '1'},
		{'1', '1', '1', '1', '1', '1', '1'},
	}
	fmt.Println(numIslands(grid))
}

// 使用一个struct来模型化问题.
type node2 struct {
	y int
	x int
}

func numIslands(grid [][]byte) int {
	// 等价于找1的集合=>竖直水平表示是一座岛屿上的
	// 岛屿定义一个head, 就是遍历头头, 最后head的数量就是岛屿数量
	// head遍历过的地方就标记为已经遍历过了
	height := len(grid)
	if height == 0 {
		return 0
	}
	width := len(grid[0])
	if width == 0 {
		return 0
	}
	head := 0
	// 从head来看, 似乎开始是深度遍历, 往这个head下面找
	// 但是从后面node来看, 又是广度遍历, 每个只找最近的一个.
	nodes := make(chan node2, height*width)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == '1' { //没有被遍历过的可以做head
				head++
				grid[y][x] = '0'
				nodes <- node2{y, x}
				//fmt.Printf("head:%d %d\n", y, x)
				for len(nodes) != 0 {
					tmp := <-nodes
					if tmp.y < height-1 {
						if grid[tmp.y+1][tmp.x] == '1' {
							grid[tmp.y+1][tmp.x] = '0'
							nodes <- node2{tmp.y + 1, tmp.x}
							//fmt.Printf("body:%d %d\n", tmp.y+1, tmp.x)
						}
					}
					if tmp.x < width-1 {
						if grid[tmp.y][tmp.x+1] == '1' {
							grid[tmp.y][tmp.x+1] = '0'
							nodes <- node2{tmp.y, tmp.x + 1}
							//fmt.Printf("body:%d %d\n", tmp.y, tmp.x+1)
						}
					}
					if tmp.x > 0 { // 额外需要向左遍历,向上遍历
						if grid[tmp.y][tmp.x-1] == '1' {
							grid[tmp.y][tmp.x-1] = '0'
							nodes <- node2{tmp.y, tmp.x - 1}
							//fmt.Printf("body:%d %d\n", tmp.y, tmp.x-1)
						}
					}
					if tmp.y > 0 {
						if grid[tmp.y-1][tmp.x] == '1' {
							grid[tmp.y-1][tmp.x] = '0'
							nodes <- node2{tmp.y - 1, tmp.x}
							//fmt.Printf("body:%d %d\n", tmp.y-1, tmp.x)
						}
					}
				}
			}
		}
	}
	return head
}
