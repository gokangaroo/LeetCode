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
	fmt.Println(searchBST(x, 3))
}

func searchBST(root *TreeNode, val int) *TreeNode {
	cur := root
	for cur != nil {
		if cur.Val == val {
			return cur
		}
		if cur.Val < val {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}
	return cur
}
