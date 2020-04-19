package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	x := &TreeNode{7,
		&TreeNode{3, nil, nil}, &TreeNode{15,
			&TreeNode{9, nil, nil}, &TreeNode{20, nil, nil}}}
	fmt.Println(maxDepth(x))
}

func maxDepth(root *TreeNode) int {
	h := 0
	dfs(root, 1, &h)
	return h
}

func dfs(node *TreeNode, depth int, h *int) {
	if node == nil {
		return
	}
	if depth > *h {
		*h = depth
	}
	dfs(node.Left, depth+1, h)
	dfs(node.Right, depth+1, h)
}
