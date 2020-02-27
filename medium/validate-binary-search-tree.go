package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{10, &TreeNode{5, nil, nil}, &TreeNode{15, &TreeNode{6, nil, nil}, &TreeNode{20, nil, nil}}}
	fmt.Println(isValidBST(root))
}

//func isValidBST(root *TreeNode) bool {
//	return dfs(root, nil, nil)
//}
//
//func dfs(node *TreeNode, low, high *int) bool {
//	if node == nil {
//		return true
//	}
//	val := node.Val
//	if low != nil && val <= *low {
//		return false
//	}
//	if high != nil && val >= *high {
//		return false
//	}
//	if !dfs(node.Right, &val, high) {
//		return false
//	}
//	if !dfs(node.Left, low, &val) {
//		return false
//	}
//	return true
//}

func isValidBST(root *TreeNode) bool {
	return dfs(root, math.MinInt64, math.MaxInt64)
}

func dfs(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || max <= root.Val {
		return false
	}
	return dfs(root.Left, min, root.Val) && dfs(root.Right, root.Val, max)
}
