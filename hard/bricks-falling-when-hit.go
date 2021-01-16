package main

import "fmt"

func main() {
	fmt.Println(hitBricks([][]int{{1, 0, 0, 0}, {1, 1, 1, 0}}, [][]int{{1, 0}}))
}

const (
	up   int = iota - 2 //-2
	left                //-1
	_
	right //1
	down  //2
)

func hitBricksTemp(grid [][]int, hits [][]int) []int {
	// 1. 第一行
	// 2. 四周有一块稳定的砖
	// hits表示拿掉这块地方的砖, res则表示该操作受到牵连掉落的砖块
	res := make([]int, len(hits))
	if len(grid) <= 1 {
		return res
	}
	if len(hits) == 0 {
		return res
	}
	// 使用并查集, 将板块union来表示是否能够掉落
	// 如果一个砖连到多个稳定的上面怎么办, 那只能使用数组存parents
	var parents = make([][][]int, len(grid))
	// 1. 初始化
	for i := 0; i < len(grid); i++ {
		parents[i] = make([][]int, len(grid[0])) //默认都是nil
	}
	// 2. union, direction是b相对于a的位置
	var union func(a []int, b []int, direction int)
	union = func(a, b []int, direction int) {
		pa := parents[a[0]][a[1]]
		pb := parents[b[0]][b[1]]
		if pa != nil && pb == nil { //b需要靠a稳定
			parents[b[0]][b[1]] = []int{direction}
		}
		if pa == nil && pb != nil { //a需要靠b稳定
			parents[a[0]][a[1]] = []int{-1 * direction}
		}
		//双方都可以让对方稳定,防止是已经有的, 所以需要去重
		//而且还要反向再去进行union告知
		if !arrayContains(pa, -1*direction) && a[0] > 0 { //a不是第一行才需要继续
			pa = append(pa, -1*direction)
			// TODO 需要向parents衍生, 表示自己也能给与支持(这个方案还没写find, 可能形成死锁,还得再想想怎么做)
		}
		if !arrayContains(pb, direction) && b[0] > 0 { //b不是第一行才需要继续
			pb = append(pb, direction)
			// TODO 需要向parents衍生, 表示自己也能给与支持
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				if i > 0 && grid[i-1][j] == 1 {
					union([]int{i, j}, []int{i - 1, j}, up)
				}
				if i < len(grid)-1 && grid[i+1][j] == 1 {
					union([]int{i, j}, []int{i + 1, j}, down)
				}
				if j > 0 && grid[i][j-1] == 1 {
					union([]int{i, j}, []int{i, j - 1}, left)
				}
				if j < len(grid)-1 && grid[i][j+1] == 1 {
					union([]int{i, j}, []int{i, j + 1}, right)
				}
			}
		}
	}
	// 用深度优先, 结合并查集递归敲掉
	for i, v := range hits {
		res[i] = fall(grid, v, parents) - 1 // -1是忽略自己
	}
	return res
}

func fall(grid [][]int, kick []int, parents [][][]int) (count int) {
	if grid[kick[0]][kick[1]] == 0 { //么有这块砖
		return 0
	}
	if kick[0] == 0 { //稳定的砖
		return 0
	}
	p := parents[kick[0]][kick[1]]
	var n = make([]int, 0) //最后剩余的parents的个数(还能依靠)
	for _, v := range p {
		switch v {
		case up:
			if grid[kick[0]-1][kick[1]] != 0 {
				n = append(n, up)
			}
		case down:
			if grid[kick[0]+1][kick[1]] != 0 {
				n = append(n, down)
			}
		case left:
			if grid[kick[0]][kick[1]-1] != 0 {
				n = append(n, left)
			}
		case right:
			if grid[kick[0]][kick[1]+1] != 0 {
				n = append(n, right)
			}
		}
	}
	if len(n) == 0 {
		count++
		grid[kick[0]][kick[1]] = 0 //掉落自己
	}
	parents[kick[0]][kick[1]] = n

	// 继续尝试掉落他的上下左右
	if kick[0] > 0 { //上: 先不考虑顶部, 看看效率
		count += fall(grid, []int{kick[0] - 1, kick[1]}, parents)
	}
	if kick[0] < len(grid)-1 { //下
		count += fall(grid, []int{kick[0] + 1, kick[1]}, parents)
	}
	if kick[1] > 0 { //左
		count += fall(grid, []int{kick[0], kick[1] - 1}, parents)
	}
	if kick[1] < len(grid[0])-1 { //右
		count += fall(grid, []int{kick[0], kick[1] + 1}, parents)
	}
	return count
}

func arrayContains(arr []int, p int) bool {
	for _, v := range arr {
		if v == p {
			return true
		}
	}
	return false
}

/*
	力扣找的方案
*/
var (
	rows, cols  int
	_directions = [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
)

func hitBricks(grid [][]int, hits [][]int) []int {
	rows, cols = len(grid), len(grid[0])

	// 第一步， 把 grid 中的砖头全部击碎，通常算法问题不能修改输入数据， 这一步非必需，可以认为是一种答题规范
	// 注意：因为 Go Slice 的机制，这里必须进行深拷贝
	copied := make([][]int, rows)
	for i, v := range grid {
		copied[i] = make([]int, len(v))
		for j := range v {
			copied[i][j] = grid[i][j]
		}
	}

	for _, v := range hits {
		copied[v[0]][v[1]] = 0
	}

	// 第二步，建图，把砖块和砖块的连接关系输入并查集， size 表示二维网格的大小， 也表示虚拟的“屋顶”在并查集中的编号
	size := rows * cols
	uf := NewUnionFind(size + 1)

	// 将下标为 0 的这一行的砖块与“屋顶”相连
	for i := 0; i < cols; i++ {
		if copied[0][i] == 1 {
			uf.union(i, size)
		}
	}
	// 其余网格，如果是砖块向上、向左看一下，如果也是砖块，在并查集中进行合并
	for i := 1; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if copied[i][j] == 1 {
				// 如果上方也是砖块
				if copied[i-1][j] == 1 {
					uf.union(getIndex(i-1, j), getIndex(i, j))
				}
				// 如果左边也是砖块
				if j > 0 && copied[i][j-1] == 1 {

					uf.union(getIndex(i, j-1), getIndex(i, j))
				}
			}
		}
	}

	// 第三步，按照 hits 的逆序， 在 copied 中补回砖块，把每一次因为补回砖块而与屋顶相连的砖块的增量记录到 res 数组中
	n := len(hits)
	res := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		x, y := hits[i][0], hits[i][1]

		// 注意： 这里不能 copy, 语义上表示，如果原来在 grid 中，这一块是空白的， 这一步补回产生任何砖块掉落
		// 逆向补回的时候，与屋顶相连的砖块数量也肯定不会增加
		if grid[x][y] == 0 {
			continue
		}

		// 补回之前与屋顶相连的砖块数
		origin := uf.getSize(size)

		// 注意： 如果补回的这个节点在第一行，要告诉并查集它与屋顶相连（逻辑同第2步）
		if x == 0 {
			uf.union(y, size)
		}

		// 在 4 个方向上看一下，如果相邻的4个方向有砖块,合并它们
		for _, i := range _directions {
			newX, newY := x+i[0], y+i[1]

			if inArea(newX, newY) && copied[newX][newY] == 1 {
				uf.union(getIndex(x, y), getIndex(newX, newY))
			}
		}

		// 补回之后与屋顶相连的砖块数
		current := uf.getSize(size)
		// 减去的 1 是逆向补回的砖块（正向移除的砖块），与 0 比较大小，是因为存在一种情况，添加当前砖块，不会使得与屋顶连接的砖块数更多
		res[i] = max(0, current-origin-1)

		// 真正补上这个砖块
		copied[x][y] = 1
	}

	return res

}

// 二维坐标转一维坐标
func getIndex(x, y int) int {
	return x*cols + y
}

// 输入坐标在二维网格中是否越界
func inArea(x, y int) bool {
	return x >= 0 && x < rows && y >= 0 && y < cols
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type unionFind struct {
	parent []int // 当前节点的父亲节点
	size   []int // 以当前节点为根子树的节点总数
}

func NewUnionFind(n int) *unionFind {
	p := make([]int, n)
	s := make([]int, n)

	for i := 0; i < n; i++ {
		p[i] = i
		s[i] = 1
	}

	return &unionFind{
		parent: p,
		size:   s,
	}

}

// 路径压缩，只要求每个不相交集合的“根节点”的子树包含的节点总数数值正确即可，因此在路径压缩的过程中不用维护数组 size
func (u *unionFind) find(x int) int {
	if x != u.parent[x] {
		u.parent[x] = u.find(u.parent[x])
	}
	return u.parent[x]

}

func (u *unionFind) union(x, y int) {
	rootX, rootY := u.find(x), u.find(y)
	if rootX == rootY {
		return
	}

	u.parent[rootX] = rootY
	u.size[rootY] += u.size[rootX]
}

// 在并查集的根节点的子树包含的节点总数
func (u *unionFind) getSize(x int) int {
	root := u.find(x)
	return u.size[root]
}
