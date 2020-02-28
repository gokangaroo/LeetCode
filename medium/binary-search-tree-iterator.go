package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	x := &TreeNode{7, &TreeNode{3, nil, nil}, &TreeNode{15, &TreeNode{9, nil, nil}, &TreeNode{20, nil, nil}}}
	iterator := Constructor_bsti(x)
	fmt.Println(iterator.Next())    // return 3
	fmt.Println(iterator.Next())    // return 7
	fmt.Println(iterator.HasNext()) // return true
	fmt.Println(iterator.Next())    // return 9
	fmt.Println(iterator.HasNext()) // return true
	fmt.Println(iterator.Next())    // return 15
	fmt.Println(iterator.HasNext()) // return true
	fmt.Println(iterator.Next())    // return 20
	fmt.Println(iterator.HasNext()) // return false
}

type BSTIterator struct {
	// 最容易想到的,就是中序遍历,然后排列成数组
	// TODO 也可以使用把TreeNode当成list来用,只存一个cursor
	inOrderList []int
	index       int
}

func Constructor_bsti(root *TreeNode) BSTIterator {
	var inOrderList []int

	// 中序遍历.
	var ldr func(root *TreeNode)
	ldr = func(root *TreeNode) {
		if root == nil {
			return
		}
		// left
		ldr(root.Left)
		// root
		inOrderList = append(inOrderList, root.Val)
		// right
		ldr(root.Right)
	}
	ldr(root)

	return BSTIterator{inOrderList, -1}
}

// return the Next smallest number
func (this *BSTIterator) Next() int {
	this.index++
	return this.inOrderList[this.index]
}

// return whether we have a Next smallest number
func (this *BSTIterator) HasNext() bool {
	return this.index < len(this.inOrderList)-1
}
