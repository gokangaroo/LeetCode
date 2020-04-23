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
	inorder := []int{3, 2, 1}
	postorder := []int{3, 2, 1}
	fmt.Println(buildTree(inorder, postorder))
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	idx := -1
	for i, v := range inorder {
		if v == postorder[len(postorder)-1] {
			idx = i
		}
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	root.Left = buildTree(inorder[:idx], postorder[:idx])
	root.Right = buildTree(inorder[idx+1:], postorder[idx:len(postorder)-1])
	return root
}
