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
	fmt.Println(levelOrder(root))
}

// 层次遍历用队列, 双队列
func levelOrder(root *Node) [][]int {
	res := make([][]int, 0)
	queue1 := make([]*Node, 0)
	queue2 := make([]*Node, 0)
	if root == nil {
		return res
	}
	queue1 = append(queue1, root)
	res = append(res, make([]int, 0))
	var depth int
	var cur *Node
	for len(queue1) != 0 {
		cur = queue1[0]
		queue1 = queue1[1:]
		res[depth] = append(res[depth], cur.Val)
		queue2 = append(queue2, cur.Children...)
		if len(queue1) == 0 {
			queue1 = queue2
			queue2 = make([]*Node, 0)
			res = append(res, make([]int, 0))
			depth++
		}
	}
	return res[:len(res)-1]
}
