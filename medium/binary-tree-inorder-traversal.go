package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	x := &TreeNode{1,
		&TreeNode{2, nil, nil}, nil}
	fmt.Println(inorderTraversal(x))
}

// 循环本地栈解法, 略微复杂, root压入后, 一股脑左边循环压入栈
// 取出来, 再看右边, 右边压入后, 再一股脑左边循环压入栈
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := []*TreeNode{root}
	var cur *TreeNode
	cur = root
	for cur.Left != nil {
		stack = append(stack, cur.Left)
		cur = cur.Left
	}
	for len(stack) != 0 {
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, cur.Val)
		if cur.Right != nil {
			stack = append(stack, cur.Right)
			for cur.Right.Left != nil {
				stack = append(stack, cur.Right.Left)
				cur.Right = cur.Right.Left
			}
		}
	}
	return res
}

func inorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0)
	inorder(root, &res)
	return res
}

func inorder(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	inorder(node.Left, res)
	*res = append(*res, node.Val)
	inorder(node.Right, res)
}
