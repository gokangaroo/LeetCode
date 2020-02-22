package main

import "fmt"

func main() {
	root := &TreeNode{5, &TreeNode{1, nil, nil}, &TreeNode{9, &TreeNode{6, nil, nil}, &TreeNode{10, nil, nil}}}
	fmt.Println(isValidBST(root))
}

func isValidBST(root *TreeNode) bool {
	return dfs(root, nil, nil)
}

func dfs(node *TreeNode, low, high *int) bool {
	if node == nil {
		return true
	}
	val := node.Val
	if low != nil && val <= *low {
		return false
	}
	if high != nil && val >= *high {
		return false
	}
	if !dfs(node.Right, &val, high) {
		return false
	}
	if !dfs(node.Left, low, &val) {
		return false
	}
	return true
}
