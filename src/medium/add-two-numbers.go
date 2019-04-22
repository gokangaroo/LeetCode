package main

import (
	. "datastructure"
	"fmt"
)

func main() {
	l1 := ListNode{1,
		&ListNode{8,
			nil},
	}
	l2 := ListNode{0, nil}
	l3 := addTwoNumbers(&l1, &l2)
	for l3 != nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}

}

// 传入的是指针, 返回的也是指针
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{0, nil}  //声明一个指针
	p, q, curr := l1, l2, dummyHead //引用传递
	carry := 0
	for p != nil || q != nil {
		x := 0
		y := 0
		if p != nil {
			x = p.Val
		}
		if q != nil {
			y = q.Val
		}
		//x := If(p != nil, p.Val, 0).(int)
		//y := If(q != nil, q.Val, 0).(int)
		sum := carry + x + y
		carry = sum / 10
		curr.Next = &ListNode{sum % 10, nil}
		curr = curr.Next
		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}
	if carry > 0 {
		curr.Next = &ListNode{carry, nil}
	}
	return dummyHead.Next
}

// 模拟一个三元运算符
//func If(condition bool, trueVal, falseVal interface{}) interface{} {
//	if condition {
//		return trueVal
//	}
//	return falseVal
//}
