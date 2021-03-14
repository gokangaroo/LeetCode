package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(flipMatchVoyage(&TreeNode{Val: 1,
		Left: nil, Right: &TreeNode{Val: 2, Left: nil, Right: nil}},
		[]int{1, 2}))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var idx int

func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	var res []int
	idx = 0
	flipMatchVoyageDFS(root, voyage, &res)
	return res
}

// 使用res记录需要反转的节点的值, 先没考虑溢出, 也没发生溢出
func flipMatchVoyageDFS(root *TreeNode, voyage []int, res *[]int) bool {
	if root == nil {
		return true
	}
	// 根节点的值和voyage第0个元素不符，无法通过反转节点使使先序遍历符合voyage
	if root.Val != voyage[idx] {
		*res = []int{-1}
		return false
	}

	// 当右儿子的值等于voyage的第1个元素，需要反转根节点的左右儿子
	// FIXME: 需要左子不空, 因为左子空, 先序遍历即使不翻转, 也是正确的.(个人觉得翻不翻都可以,但是题目答案是不翻转)
	// FIXME: voyage[idx]和[idx+1]竟然没溢出过, 题目给的条件应该是避免了这种情况,不然考虑起来麻烦了
	if root.Left != nil && root.Right != nil && root.Right.Val == voyage[idx+1] {
		*res = append(*res, root.Val)
		root.Left, root.Right = root.Right, root.Left
	}
	// 消耗掉voyage第0个元素
	idx++
	// 左边无法满足翻转,无需再去看右边, 直接回溯返回答案
	if !flipMatchVoyageDFS(root.Left, voyage, res) {
		return false
	}
	return flipMatchVoyageDFS(root.Right, voyage, res)
}
