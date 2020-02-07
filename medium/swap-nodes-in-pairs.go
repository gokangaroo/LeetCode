package main

import (
	"fmt"
	. "github/gokangaroo/LeetCode/datastructure"
)

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	head = swapPairsNew(head)
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

//非递归
func swapPairsNew(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	//初始化临时节点
	ptr := &ListNode{Val: 0, Next: head}
	current := head.Next
	//进行交换
	for {
		next := ptr.Next.Next
		ptr.Next.Next = next.Next
		next.Next = ptr.Next
		ptr.Next = next

		ptr = ptr.Next.Next
		if ptr.Next == nil || ptr.Next.Next == nil {
			break
		}
	}
	return current
}
