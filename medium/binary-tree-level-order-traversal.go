package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res = make([][]int, 0)

func main() {
	root := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	res := levelOrder(root)
	fmt.Println(res)
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return res
	}
	helper(root, 0)
	return res
}

func helper(node *TreeNode, level int) {
	if len(res) == level {
		res = append(res, make([]int, 0))
	}

	res[level] = append(res[level], node.Val)

	if node.Left != nil {
		helper(node.Left, level+1)
	}
	if node.Right != nil {
		helper(node.Right, level+1)
	}
}
