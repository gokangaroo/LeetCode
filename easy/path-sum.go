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
	fmt.Println(hasPathSum(x, 3))
}

func hasPathSum(root *TreeNode, sum int) bool {
	return dfsroute(root, 0, sum)
}

func dfsroute(node *TreeNode, sum int, aim int) bool {
	if node == nil {
		return false
	}
	sum += node.Val
	if node.Left == nil && node.Right == nil { // base case
		return sum == aim
	}
	return dfsroute(node.Left, sum, aim) || dfsroute(node.Right, sum, aim)
}
