package main

import (
	. "datastructure"
	"fmt"
)
func main()  {
	node1 := ListNode{3,nil}
	node2 := ListNode{2,nil}
	node3 := ListNode{0,nil}
	node4 := ListNode{4,nil}
	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node2
	res := detectCycle(&node1)
	fmt.Println((*res).Val)
}

func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head

	isCycle := false

	for fast != nil && fast.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow{
			isCycle = true
			break
		}
	}

	//找到入口节点
	if isCycle {
		p := head
		for slow != p{
			p  = p.Next
			slow = slow.Next
		}
		return p
	}
	return nil
}