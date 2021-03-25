package main

import "fmt"

func main() {
	fmt.Println(deleteDuplicates(
		&ListNode{Val: 3, Next:
		&ListNode{Val: 3, Next:
		&ListNode{Val: 5, Next: nil}}}))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	origin := &ListNode{
		Val:  0,
		Next: head,
	}
	cur := origin
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			cur2 := cur.Next.Next
			for cur2 != nil && cur2.Val == cur.Next.Val {
				cur2 = cur2.Next
			}
			cur.Next = cur2
		} else {
			cur = cur.Next
		}
	}
	return origin.Next
}
