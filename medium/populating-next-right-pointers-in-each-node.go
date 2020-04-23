package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func main() {
	x := &Node{1,
		&Node{2,
			&Node{4, nil, nil, nil}, &Node{5, nil, nil, nil}, nil},
		&Node{3,
			&Node{6, nil, nil, nil}, &Node{7, nil, nil, nil}, nil}, nil}
	fmt.Println(connect(x))
}

// 层次遍历, 该题每层都是一个链表, 我们的目的就是组建链表
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	for cur := root; cur.Left != nil; cur = cur.Left { //到下一层最左边,最左游标
		for c := cur; c != nil; c = c.Next { //遍历完这一层, 层次游标
			c.Left.Next = c.Right //1.两个子节点, 情况1可以满足情况2
			if c.Next != nil {    //2.跨越节点
				c.Right.Next = c.Next.Left
			}
		}
	}
	return root
}
