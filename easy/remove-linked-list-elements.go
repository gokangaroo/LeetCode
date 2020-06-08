package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	fmt.Println(removeElements(head, 1))
}

func removeElements(head *ListNode, val int) *ListNode {
	h := &ListNode{0, head}
	cur1, cur2 := h, head
	if cur2 == nil {
		return h.Next
	}
	for cur2.Next != nil {
		if cur2.Val == val {
			cur1.Next = cur2.Next
		} else {
			cur1 = cur1.Next
		}
		cur2 = cur2.Next
	}
	if cur2.Val == val {
		cur1.Next = nil
	}
	return h.Next
}
