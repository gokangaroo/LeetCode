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
	x := &TreeNode{1,
		nil, &TreeNode{2,
			&TreeNode{3, nil, nil}, nil}}
	fmt.Println(preorderTraversal(x))
}

// 循环本地栈解法, 取出来再放右边和左边
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root==nil {
		return res
	}
	stack := []*TreeNode{root}
	var cur *TreeNode
	for len(stack) != 0 {
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, cur.Val)
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}
	return res
}

func preorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0)
	preorder(root, &res)
	return res
}

func preorder(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	*res = append(*res, node.Val)
	preorder(node.Left, res)
	preorder(node.Right, res)
}
