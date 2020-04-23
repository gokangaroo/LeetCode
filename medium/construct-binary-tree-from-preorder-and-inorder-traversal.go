package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	pre := []int{1, 2, 3}
	in := []int{3, 2, 1}
	fmt.Println(buildTree(pre, in))
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	idx := -1
	for i, v := range inorder {
		if v == preorder[0] {
			idx = i
		}
	}
	root := &TreeNode{Val: preorder[0]}
	root.Left = buildTree(preorder[1:idx+1], inorder[:idx])
	root.Right = buildTree(preorder[idx+1:], inorder[idx+1:])
	return root
}
