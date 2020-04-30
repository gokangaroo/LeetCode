package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// [5,3,6,2,4,null,7] 3
	x := &TreeNode{5,
		&TreeNode{3,
			&TreeNode{2, nil, nil}, &TreeNode{4, nil, nil}},
		&TreeNode{6, nil, &TreeNode{7, nil, nil}}}
	fmt.Println(deleteNode(x, 3))
}

// 写一个递归, 返回替代的节点
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		//没找到
		return nil
	}
	if root.Val == key { // 就是需要删除的节点
		if root.Left == nil && root.Right == nil {
			return nil // 1.如果两子为空, 直接删
		}                                          // 2.如果有子节点, 选一个替代, 但是要细分两种情况.
		if root.Left != nil && root.Right != nil { // 2.1两个子都在,找到"最靠近"的一个数字
			cur := root.Left
			father := root
			for cur.Right != nil {
				father = cur
				cur = cur.Right
			}
			// cur就是"左边最靠近"的数字
			cur.Right = root.Right
			if father != root { //特殊情况father==root
				father.Right = cur.Left
				cur.Left = root.Left
			}
			return cur
		}
		if root.Left == nil {
			return root.Right
		}
		return root.Left
	}
	root.Left = deleteNode(root.Left, key)
	root.Right = deleteNode(root.Right, key)
	return root
}
