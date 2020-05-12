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
	fmt.Println(postorder(root))
}

func postorder(root *Node) []int {
	res := make([]int, 0)
	postorder2(root, &res)
	return res
}

func postorder2(node *Node, res *[]int) {
	if node == nil {
		return
	}
	for _, v := range node.Children {
		postorder2(v, res)
	}
	*res = append(*res, node.Val)
}
