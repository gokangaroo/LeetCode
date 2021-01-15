package main

import "fmt"

func main() {
	fmt.Println(removeStones([][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}))
}

func removeStones(stones [][]int) int {
	n := len(stones)
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	var union func(x, y int)
	union = func(x, y int) {
		i := find(x)
		j := find(y)
		parent[i] = parent[j] // x与y轴相连通(同一个连通分量)
	}

	// 双指针遍历, 如果两个石头在同行, 或者同列, 就直接union为同一个连通分量
	// 这里只要union stone在原stones的下标即可
	for i := 0; i < n-1; i++ {
		s1 := stones[i]
		for j := i + 1; j < n; j++ {
			s2 := stones[j]
			if s1[0] == s2[0] || s1[1] == s2[1] {
				union(i, j)
			}
		}
	}

	// 实际移除时, 同一个连通分量的石头会被移除到只剩一个,也就是最后没有再union的石头
	// 向一个方向遍历, 并持续进行union(将x轴与y轴联通), 最终留下的石头是i=find(i)的石头
	// 最后只需要计数即可(题目只关心结果, 不关心我们怎么移除)
	count := 0
	for i, _ := range parent {
		if i == find(i) {
			count++
		}
	}
	return n - count
}
