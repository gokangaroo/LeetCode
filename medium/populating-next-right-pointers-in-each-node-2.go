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

// 满足非完全二叉树
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	for cur := root; getNextHead(cur) != nil; cur = getNextHead(cur) { //1.NextListHead不再是cur.Left
		for c := cur; c != nil; c = c.Next { //2.遍历链表
			if c.Left != nil && c.Right != nil { //one: 普通场景
				c.Left.Next = c.Right
			}
			if c.Next != nil { //two: 复杂场景
				bridge(c)
			}
		}
	}
	return root
}

func bridge(head *Node) {
	if head.Right != nil {
		head.Right.Next = getNextHead(head.Next)
		return
	}
	if head.Left != nil {
		head.Left.Next = getNextHead(head.Next)
		return
	}
}

func getNextHead(cur *Node) *Node {
	if cur == nil {
		return nil
	}
	if cur.Left != nil {
		return cur.Left
	}
	if cur.Right != nil {
		return cur.Right
	}
	return getNextHead(cur.Next)
}
