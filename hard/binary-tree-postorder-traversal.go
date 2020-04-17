package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	x := &TreeNode{1,
		nil, &TreeNode{2,
			&TreeNode{3, nil, nil}, nil}}
	fmt.Println(postorderTraversal2(x))
}

// 循环本地栈解法, 与前序遍历左右翻转, 上下翻转
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := []*TreeNode{root}
	var cur *TreeNode
	for len(stack) != 0 {
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, cur.Val)
		// 与前序左右相反
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
	}
}

func postorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0)
	postorder(root, &res)
	return res
}

func postorder(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	postorder(node.Left, res)
	postorder(node.Right, res)
	*res = append(*res, node.Val)
}
