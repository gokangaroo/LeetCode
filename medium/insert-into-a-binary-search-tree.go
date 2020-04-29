package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{10,
		&TreeNode{5, nil, nil}, &TreeNode{15,
			&TreeNode{6, nil, nil}, &TreeNode{20, nil, nil}}}
	fmt.Println(insertIntoBST(root, 3))
}

// 直接插入到叶子节点, 最简单的方式
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val, nil, nil}
	}
	cur := root
	for cur.Val != val {
		if cur.Val > val {
			if cur.Left == nil {
				cur.Left = &TreeNode{val, nil, nil}
			}
			cur = cur.Left
		} else {
			if cur.Right == nil {
				cur.Right = &TreeNode{val, nil, nil}
			}
			cur = cur.Right
		}
	}
	return root
}
