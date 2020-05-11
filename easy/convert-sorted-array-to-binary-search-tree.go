package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	xx := []int{-10, -3, 0, 5, 9}
	fmt.Println(sortedArrayToBST(xx))
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	// sort.Ints(nums)
	// 直接用递归
	var node *TreeNode
	if len(nums) == 1 {
		node = &TreeNode{nums[0], nil, nil}
		return node
	}
	if len(nums) == 2 {
		node = &TreeNode{nums[1], nil, nil}
		node.Left = &TreeNode{nums[0], nil, nil}
		return node
	}
	node = &TreeNode{nums[len(nums)/2], nil, nil}
	node.Left = sortedArrayToBST(nums[:len(nums)/2])
	node.Right = sortedArrayToBST(nums[len(nums)/2+1:])
	return node
}
