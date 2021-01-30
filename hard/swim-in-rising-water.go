package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(swimInWater([][]int{
		{0, 2},
		{1, 3},
	}))
}

// 2. 并查集写法(0,0和n-1,n-1连通即成功), 循环时间戳
var direction = [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}} //上下左右

func swimInWater(grid [][]int) int {
	n := len(grid) // 正方形的边长
	if n == 0 {
		return 0
	}
	// n*n个sink拿出来排序
	sinks := make([]*sink, 0, n*n)
	sinkTrue := make(map[int]bool, n*n/8)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sinks = append(sinks, &sink{
				i:      i,
				j:      j,
				height: grid[i][j],
			})
		}
	}
	// 按照高度排序=>变成了爬楼梯, 但是也要考虑原上下左右是否有更短的(union过去)
	sort.Slice(sinks, func(i, j int) bool {
		return sinks[i].height < sinks[j].height
	})

	// 并查集, union将Head合并
	// 抽象编号sink, n*x+y
	uf := make([]int, n*n)
	for i := 0; i < len(uf); i++ {
		uf[i] = i
	}
	var find func(int) int
	find = func(i int) int {
		for uf[i] != i {
			return find(uf[i])
		}
		return i
	}
	var union func(i, j int)
	union = func(i, j int) {
		rootI := find(i)
		rootJ := find(j)
		if rootI < rootJ {
			uf[rootJ] = rootI
		} else {
			uf[rootI] = rootJ
		}
	}

	// 时间推移, 直到右尾与首部连接
	t := 0                             // t是时间
	var inArea = func(x, y int) bool { // 是否越界
		return x >= 0 && x < n && y >= 0 && y < n
	}
	for i := 0; find(n*n-1) != 0; t++ { // find(n*n-1) == 0表示首尾相连, 在同一个图内
		for i < n*n && sinks[i].height == t {
			cur := sinks[i]
			id := cur.i*n + cur.j // id是sink编号(编号是抽象) => n*x+y => n*(n-1)+n-1 => n*n-1
			sinkTrue[id] = true
			for _, d := range direction {
				nx, ny := cur.i+d[0], cur.j+d[1]         // near x,y
				if inArea(nx, ny) && sinkTrue[nx*n+ny] { // 如果sinkTrue, 就要union
					// 在这个时间点, 可以从上下左右的某个方向游进来
					union(id, nx*n+ny)
				}
			}
			i++
		}
	}
	return t - 1 //这个时间, 让n*n-1编号的sink与0编号的sink想连接了
}

type sink struct {
	i      int
	j      int
	height int
}

// FIXME 2. 二分查找+dfs写法, 这是copy的, 后面看
func swimInWater2(grid [][]int) int {
	return sort.Search(max(grid)+1, func(i int) bool { // 这里其实有点浪费，在[0,max]的区间里搜所的
		if i < grid[0][0] {
			return false
		}
		return canReach(i, grid)
	})
}

func max(grid [][]int) int {
	result := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid); c++ {
			if grid[r][c] > result {
				result = grid[r][c]
			}
		}
	}
	return result
}

func canReach(t int, grid [][]int) bool {
	const maxN = 50
	n := len(grid)
	visited := [maxN][maxN]bool{}
	var dfs func(r, c int) bool
	dfs = func(r, c int) bool {
		if r < 0 || c < 0 || r >= n || c >= n ||
			visited[r][c] || grid[r][c] > t {
			return false
		}
		if r == n-1 && c == n-1 {
			return true
		}
		visited[r][c] = true
		return dfs(r+1, c) || dfs(r-1, c) || dfs(r, c+1) || dfs(r, c-1)
	}
	return dfs(0, 0)
}
