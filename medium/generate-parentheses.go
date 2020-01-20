package main

import (
	"fmt"
	"math"
)

func main() {
	n := 5
	res := generateParenthesis(n)
	fmt.Println(res)

}

func generateParenthesis(n int) []string {
	var res []string
	if n == 0 {
		return res
	}
	dfs("", 0, 0, n, &res)
	return res
}

// 深度优先搜索,结果都在叶子节点
func dfs(str string, left int, right int, n int, res *[]string) {
	// 遍历到叶子节点,结算
	if left == n && right == n {
		*res = append(*res, str)
		return
	}

	// 剪枝
	if left < right {
		return
	}

	// 进入左子树
	if left < n {
		dfs(str+"(", left+1, right, n, res)
	}

	// 进入右子树
	if right < n {
		dfs(str+")", left, right+1, n, res)
	}
}

// 也可以使用广度优先
// nodes这个大小一定要设置的恰到好处,不然效率会比深度优先差很多
type node struct {
	res   string
	left  int
	right int
}

func generateParenthesis2(n int) []string {
	var res []string
	if n == 0 {
		return res
	}
	var cursor node
	nodes := make(chan node, int(math.Pow(2, float64(2*n-2))))
	nodes <- node{"", n, n}
	for len(nodes) != 0 {
		cursor = <-nodes
		if cursor.left == 0 && cursor.right == 0 {
			res = append(res, cursor.res)
		}
		if cursor.left > 0 {
			nodes <- node{cursor.res + "(", cursor.left - 1, cursor.right}
		}
		if cursor.right > 0 && cursor.right > cursor.left {
			nodes <- node{cursor.res + ")", cursor.left, cursor.right - 1}
		}
	}
	return res
}

// 也可以使用动态规划
