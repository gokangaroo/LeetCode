package main

import (
	"fmt"
	. "github/gokangaroo/LeetCode/datastructure"
)

func main() {
	//-16557,-8725,-29125,28873,-21702,15483,-28441,-17845,-4317,-10914
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

	node20 := ListNode{-16557, nil}
	node19 := ListNode{-8725, nil}
	node18 := ListNode{-29125, nil}
	node17 := ListNode{28873, nil}
	node16 := ListNode{-21702, nil}
	node15 := ListNode{15483, nil}
	node14 := ListNode{-28441, nil}
	node13 := ListNode{-17845, nil}
	node12 := ListNode{-4317, nil}
	node11 := ListNode{-10914, nil}
	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5
	node5.Next = &node6
	node6.Next = &node7
	node7.Next = &node8
	node8.Next = &node9
	node9.Next = &node10
	node10.Next = &node11
	node11.Next = &node12
	node12.Next = &node13
	node13.Next = &node14
	node14.Next = &node15
	node15.Next = &node16
	node16.Next = &node17
	node17.Next = &node18
	node18.Next = &node19
	node19.Next = &node20
	fmt.Println(isPalindrome_list2(&node1))
}

// 这是翻转后半部分, 需要多费一半时间, 也可以直接翻转前半部分.
func isPalindrome_list(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	// 可以用快慢指针, 当快指针走到结尾说明走到一半了
	node1 := head
	node2 := head.Next
	for node2.Next != nil {
		node1 = node1.Next
		node2 = node2.Next
		if node2.Next != nil {
			node2 = node2.Next
		}
	}
	// 无论是单双, 现在node1都在前半部分的最后一个node
	cur := node1.Next //cur扫到就插到node1后面,来翻转后半部分
	var nextNode *ListNode
	var leftHead *ListNode = nil //后半部分另起炉灶
	for cur != nil {
		// 1.保存cur的下一个节点,最主要的一步
		nextNode = cur.Next
		// 2.cur.Next链接leftHead,就是把cur插到原head前面成为head
		cur.Next = leftHead
		leftHead = cur
		// 3.cur往后移动一次
		cur = nextNode
	}
	node1.Next = leftHead
	node1 = node1.Next
	for node1 != nil {
		if head.Val != node1.Val {
			return false
		}
		node1 = node1.Next
		head = head.Next
	}
	return true
}

func isPalindrome_list2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	even := false
	// 可以用快慢指针, 当快指针走到结尾说明走到一半了
	node1 := head
	node2 := head.Next
	var nextNode *ListNode
	var leftHead *ListNode = nil //后半部分另起炉灶
	for node2.Next != nil {
		nextNode = node1.Next
		node1.Next = leftHead
		leftHead = node1
		// 3.cur往后移动一次
		node1 = nextNode
		node2 = node2.Next
		if node2.Next != nil {
			node2 = node2.Next
		} else {
			//说明是奇数
			even = true
		}
	}
	nextNode = node1.Next
	node1.Next = leftHead
	leftHead = node1
	node1 = nextNode
	// 前面已经全部翻转, 如果是奇数会导致开头多一个
	if even {
		leftHead = leftHead.Next
	}
	for node1 != nil {
		if leftHead.Val != node1.Val {
			return false
		}
		node1 = node1.Next
		leftHead = leftHead.Next
	}
	return true
}
