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

// TODO 循环本地栈解法.
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	postorder(root, &res)
	return res
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
