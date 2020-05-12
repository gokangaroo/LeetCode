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
	fmt.Println(maxDepth(root))
}

func maxDepth(root *Node) int {
	tmp := 0
	new := 0
	if root == nil {
		return tmp
	}
	for _, v := range root.Children {
		new = maxDepth(v)
		if new > tmp {
			tmp = new
		}
	}
	return tmp + 1
}
