package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l := &ListNode{Val: 2, Next: &ListNode{Val: 4}}
	l1 := &ListNode{Next: &ListNode{Val: 9, Next: &ListNode{Val: 1, Next: l}}}
	l2 := &ListNode{Val: 3, Next: l}
	fmt.Println(getIntersectionNode(l1, l2))
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	//走到头大家都一样, 但是需要不能有环
	curA, curB := headA, headB
	if curA == nil || curB == nil {
		return nil
	}
	var countA, countB int
	for curA.Next != nil {
		curA = curA.Next
		countA++
	}
	for curB.Next != nil {
		curB = curB.Next
		countB++
	}
	if curA != curB { //大家不是一路人
		return nil
	}
	curA, curB = headA, headB
	for countA != countB { // 谁走到前面了就等等
		if countA > countB {
			curA = curA.Next
			countA--
		} else {
			curB = curB.Next
			countB--
		}
	}
	for curA != curB {
		curA = curA.Next
		curB = curB.Next
	}
	return curA
	// 还有种解法就是, A走完了A去走B, B走完了B去走A, 如果是一样的, 会直接走到一起(两人步数一样)
}
