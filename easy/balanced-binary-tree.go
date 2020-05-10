package main

import (
	"fmt"
	"math"
)

func main() {
	//root := &TreeNode{1,
	//	&TreeNode{2,
	//		&TreeNode{3,
	//			&TreeNode{4, nil, nil},
	//			&TreeNode{4, nil, nil}},
	//		&TreeNode{3, nil, nil}},
	//	&TreeNode{2, nil, nil}}
	root := &TreeNode{1,
		&TreeNode{2,
			&TreeNode{3,
				&TreeNode{4, nil, nil}, &TreeNode{4, nil, nil}},
			&TreeNode{3, nil, nil}},
		&TreeNode{2, nil, nil}}
	fmt.Println(isBalanced(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced1(root *TreeNode) (res bool) {
	if root == nil {
		return true
	}
	c := make(chan bool)
	go checkBalance(root, 0, c)
	return <-c
}

func checkBalance(x *TreeNode, height float64, c chan bool) float64 {
	if height == 0 {
		defer func() {
			select {
			case c <- true:
			default:
			}
		}()
	}
	height++
	if x.Left == nil && x.Right == nil {
		return height
	}
	if x.Left == nil {
		if x.Right.Left != nil || x.Right.Right != nil {
			c <- false
		}
		return height + 1
	}
	if x.Right == nil {
		if x.Left.Left != nil || x.Left.Right != nil {
			c <- false
		}
		return height + 1
	}
	l := checkBalance(x.Left, height, c)
	r := checkBalance(x.Right, height, c)
	//fmt.Println(l, r)
	if math.Abs(l-r) > 1 {
		c <- false
	}
	return math.Max(l, r)
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 左右是平衡树, 且树深度差在1以内
	res, _ := isBalancedNode(root)
	return res
}

func isBalancedNode(node *TreeNode) (res bool, height int) {
	if node == nil {
		return true, 0
	}
	l, lh := isBalancedNode(node.Left)
	r, rh := isBalancedNode(node.Right)
	if !(l && r) {
		return
	}
	if lh-rh < -1 || lh-rh > 1 {
		return
	}
	return true, max(lh, rh) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
