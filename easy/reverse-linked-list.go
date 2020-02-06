package main

import (
	"fmt"
	. "github/gokangaroo/LeetCode/datastructure"
)

func main() {
	node1 := ListNode{-16557, nil}
	node2 := ListNode{-8725, nil}
	node3 := ListNode{-29125, nil}
	node4 := ListNode{28873, nil}
	node5 := ListNode{-21702, nil}
	node6 := ListNode{15483, nil}
	node7 := ListNode{-28441, nil}
	node8 := ListNode{-17845, nil}
	node9 := ListNode{-4317, nil}
	node10 := ListNode{-10914, nil}
	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5
	node5.Next = &node6
	node6.Next = &node7
	node7.Next = &node8
	node8.Next = &node9
	node9.Next = &node10
	list := reverseList(&node1)
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}

func reverseList(cur *ListNode) *ListNode {
	var nextNode *ListNode
	var preNode *ListNode = nil
	for cur != nil {
		// 1.保存cur的下一个节点
		nextNode = cur.Next
		// 2.将cur的node放到pre位置
		cur.Next = preNode
		preNode = cur
		// 3.cur指向后一个位置
		cur = nextNode
	}
	return preNode
}
