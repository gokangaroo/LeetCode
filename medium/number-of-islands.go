package main

import "fmt"

func main() {
	// golang 这个byte 1和0有点二, 这道题有点二
	var grid = [][]byte{
		{49, 49, 49, 49, 49, 49, 49},
		{48, 48, 48, 48, 48, 48, 49},
		{49, 49, 49, 49, 49, 48, 49},
		{49, 48, 48, 48, 49, 48, 49},
		{49, 48, 49, 48, 49, 48, 49},
		{49, 48, 49, 49, 49, 48, 49},
		{49, 49, 49, 49, 49, 49, 49},
	}
	fmt.Println(numIslands(grid))
}

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
	over := make([][]byte, height)
	for index := 0; index < height; index++ {
		over[index] = make([]byte, width)
	}
	head := 0
	nodes := make(chan node2, height*width)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == 49 && over[y][x] == 0 { //没有被遍历过的可以做head
				head++
				over[y][x] = 1
				nodes <- node2{y, x}
				//fmt.Printf("head:%d %d\n", y, x)
				for len(nodes) != 0 {
					tmp := <-nodes
					if tmp.y < height-1 {
						if grid[tmp.y+1][tmp.x] == 49 && over[tmp.y+1][tmp.x] == 0 {
							over[tmp.y+1][tmp.x] = 1
							nodes <- node2{tmp.y + 1, tmp.x}
							//fmt.Printf("body:%d %d\n", tmp.y+1, tmp.x)
						}
					}
					if tmp.x < width-1 {
						if grid[tmp.y][tmp.x+1] == 49 && over[tmp.y][tmp.x+1] == 0 {
							over[tmp.y][tmp.x+1] = 1
							nodes <- node2{tmp.y, tmp.x + 1}
							//fmt.Printf("body:%d %d\n", tmp.y, tmp.x+1)
						}
					}
					if tmp.x > 0 { // 额外需要向左遍历,向上遍历
						if grid[tmp.y][tmp.x-1] == 49 && over[tmp.y][tmp.x-1] == 0 {
							over[tmp.y][tmp.x-1] = 1
							nodes <- node2{tmp.y, tmp.x - 1}
							//fmt.Printf("body:%d %d\n", tmp.y, tmp.x-1)
						}
					}
					if tmp.y > 0 {
						if grid[tmp.y-1][tmp.x] == 49 && over[tmp.y-1][tmp.x] == 0 {
							over[tmp.y-1][tmp.x] = 1
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
