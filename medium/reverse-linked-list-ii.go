package main

import "fmt"

func main() {
	fmt.Println(reverseBetween(&ListNode{Val: 3, Next: &ListNode{Val: 5, Next: nil}}, 1, 2))
}

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	/*
	   翻转列表使用头插法, 将left及之后的node, 逐个调度到right后面
	*/
	headPrev := &ListNode{Val: 0, Next: head}
	var l, r, tmp *ListNode
	all := left == 1
	for {
		if left == 1 {
			l = headPrev // leftPrev
		}
		if right == 0 {
			r = headPrev // right
			break
		}
		left--
		right--
		headPrev = headPrev.Next
	}
	if l == nil || r == nil {
		return head
	}
	for l.Next != r {
		// del
		tmp = l.Next
		l.Next = l.Next.Next
		// add
		tmp.Next = r.Next
		r.Next = tmp
	}
	if all { // 全部翻转
		return r
	}
	return head
}
