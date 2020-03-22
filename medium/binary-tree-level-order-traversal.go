package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{3,
		&TreeNode{9, nil, nil},
		&TreeNode{20,
			&TreeNode{15, nil, nil},
			&TreeNode{7, nil, nil}}}
	res := levelOrder(root)
	fmt.Println(res)
}

var res = make([][]int, 0)

func levelOrder(root *TreeNode) [][]int {
	res = make([][]int, 0)
	if root == nil {
		return res
	}
	helper(root, 0)
	return res
}

func helper(node *TreeNode, depth int) {
	if len(res) == depth {
		res = append(res, make([]int, 0))
	}

	res[depth] = append(res[depth], node.Val)

	if node.Left != nil {
		helper(node.Left, depth+1)
	}
	if node.Right != nil {
		helper(node.Right, depth+1)
	}
}
