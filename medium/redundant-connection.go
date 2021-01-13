package main

import (
	"fmt"
)

func main() {
	edges := [][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}, {1, 5}}
	fmt.Println(findRedundantConnection(edges))
}

func findRedundantConnection(edges [][]int) []int {
	// 并查集
	parent := make([]int, len(edges)+1)
	// 1.初始parent都是自己
	for i := range parent {
		parent[i] = i
	}
	// 2. find func, 如果parent不是自己,就逐层往上找
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	// 3. union func, 将c的parent设置与parent的parent一样
	// 可以理解成, 链表合并, 图合并
	var union func(c, p int) bool
	union = func(c, p int) bool {
		x, y := find(c), find(p)
		if x == y {
			return false
		}

		parent[x] = y
		return true
	}
	// 4. union失败, 表明成环了
	for _, e := range edges {
		if !union(e[0], e[1]) {
			return e
		}
	}
	return nil
}
