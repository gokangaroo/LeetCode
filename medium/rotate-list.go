package main

import "fmt"

func main() {
	fmt.Println(rotateRight(
		&ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}}, 4))
}

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	/*
	   链表向右移动k个位置=>将最后k个元素移动到链表前面

	   链表的长度未知, k可能过于大, 先遍历一遍
	*/
	// 1. 处理k
	cur := head
	length := 0
	for cur != nil {
		length++
		cur = cur.Next
	}
	if length == 0 {
		return head
	}
	k %= length
	// 2. 找到需要转移的段
	origin := &ListNode{
		Val:  0,
		Next: head,
	}
	cur = origin
	for k > 0 {
		cur = cur.Next
		k--
	}
	cur2 := origin
	for cur.Next != nil {
		cur = cur.Next
		cur2 = cur2.Next
	}
	// 3. 将cur2.Next->cur都插入到orgin后面
	cur.Next = origin.Next
	origin.Next = cur2.Next
	cur2.Next = nil
	return origin.Next
}
