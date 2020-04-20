package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	x := &TreeNode{1,
		&TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	fmt.Println(isSymmetric(x))
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return compareReverse(root.Left, root.Right)
}

func compareReverse(left *TreeNode, right *TreeNode) bool {
	if left == nil || right == nil { //base case
		return left == nil && right == nil
	}
	if left.Val != right.Val {
		return false
	}
	return compareReverse(left.Left, right.Right) && compareReverse(left.Right, right.Left)
}
