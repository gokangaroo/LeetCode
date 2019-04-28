package main

import (
	. "datastructure"
	"fmt"
)

func main() {
	l1 := ListNode{1, nil}
	l1.Next = &ListNode{2, nil}
	l2 := ListNode{1, nil}
	l2.Next = &ListNode{3, nil}
	l2.Next.Next = &ListNode{4, nil}
	resultList := mergeTwoLists(&l1, &l2)
	for resultList != nil {
		fmt.Println(resultList.Val)
		resultList = resultList.Next
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := ListNode{0, nil}
	p, q, curr := l1, l2, &dummyHead
	for p != nil || q != nil {
		if p == nil {
			curr.Next = q
			q = nil
		} else if q == nil {
			curr.Next = p
			p = nil
		}
		if p != nil && q != nil {
			if p.Val < q.Val {
				curr.Next = &ListNode{p.Val, nil}
				curr = curr.Next
				p = p.Next
			} else if p.Val > q.Val {
				curr.Next = &ListNode{q.Val, nil}
				curr = curr.Next
				q = q.Next
			} else { //相等的情况
				curr.Next = &ListNode{p.Val, nil}
				curr.Next.Next = &ListNode{q.Val, nil}
				curr = curr.Next.Next
				p = p.Next
				q = q.Next
			}
		}
	}
	return dummyHead.Next
}