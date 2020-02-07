package main

import (
	"fmt"
	. "github/gokangaroo/LeetCode/datastructure"
)

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	head = swapPairs(head)
	fmt.Println(head)
}

//递归
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//进行交换
	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return next
}