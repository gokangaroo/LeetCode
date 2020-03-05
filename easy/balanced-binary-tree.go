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
	root := &TreeNode{2,
		&TreeNode{1, nil, nil},
		&TreeNode{4,
			&TreeNode{3, nil, nil},
			&TreeNode{5, nil,
				&TreeNode{6, nil, nil}}}}
	fmt.Println(isBalanced(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) (res bool) {
	if root == nil {
		return true
	}
	defer x()
	xxx(root, 0)
	return true
}

func xxx(x *TreeNode, height float64) float64 {
	height++
	if x.Left == nil && x.Right == nil {
		return height
	}
	if x.Left == nil {
		if x.Right.Left != nil || x.Right.Right != nil {
			panic("x")
		}
		return height + 1
	}
	if x.Right == nil {
		if x.Left.Left != nil || x.Left.Right != nil {
			panic("x")
		}
		return height + 1
	}
	l := xxx(x.Left, height)
	r := xxx(x.Right, height)
	//fmt.Println(l, r)
	if math.Abs(l-r) > 1 {
		panic("x")
	}
	return math.Max(l, r)
}

func x() {
	if p := recover(); p != nil {
		//fmt.Printf("%+v\n", p)
	}
}
