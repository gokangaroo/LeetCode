package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{1, nil}
	fmt.Println(removeNthFromEnd(head,1))
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cur1, cur2 := head, head
	for i := 0; i < n; i++ {
		cur2 = cur2.Next
	}
	if cur2 == nil { //删除head
		return head.Next
	}
	for cur2.Next != nil {
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	cur1.Next = cur1.Next.Next
	return head
}
