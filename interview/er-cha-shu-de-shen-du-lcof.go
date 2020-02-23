package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}

	fmt.Println(maxDepth(root))
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0

	return dfs(root, res)
}

func dfs(node *TreeNode, depth int) int {
	depth++
	left, right := 0, 0
	if node.Left != nil {
		left = dfs(node.Left, depth)
	}
	if node.Right != nil {
		right = dfs(node.Right, depth)
	}
	if node.Left == nil && node.Right == nil {
		return depth
	}
	if left < right {
		return right
	} else {
		return left
	}
}
