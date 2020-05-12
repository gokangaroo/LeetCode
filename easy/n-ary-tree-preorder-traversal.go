package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func main() {
	root := &Node{1, []*Node{
		{3, []*Node{
			{5, nil},
			{8, nil},
		}},
		{2, nil},
		{4, nil},
	}}
	fmt.Println(preorder(root))
}

func preorder(root *Node) []int {
	res := make([]int, 0)
	preorder2(root, &res)
	return res
}

func preorder2(node *Node, res *[]int) {
	if node == nil {
		return
	}
	*res = append(*res, node.Val)
	for _, v := range node.Children {
		preorder2(v, res)
	}
}
