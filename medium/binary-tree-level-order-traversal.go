package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{3,
		&TreeNode{9, nil, nil},
		&TreeNode{20,
			&TreeNode{15, nil, nil},
			&TreeNode{7, nil, nil}}}
	res := levelOrder(root)
	fmt.Println(res)
}

var res = make([][]int, 0)

func levelOrder2(root *TreeNode) [][]int {
	res = make([][]int, 0)
	if root == nil {
		return res
	}
	helper(root, 0)
	return res
}

func helper(node *TreeNode, depth int) {
	if len(res) == depth {
		res = append(res, make([]int, 0))
	}

	res[depth] = append(res[depth], node.Val)

	if node.Left != nil {
		helper(node.Left, depth+1)
	}
	if node.Right != nil {
		helper(node.Right, depth+1)
	}
}

// 广度遍历
func levelOrder(root *TreeNode) [][]int {
	h := 0
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	tmpQueue := make([]*TreeNode, 0) //用两个queue
	var cur *TreeNode
	for len(queue) != 0 {
		cur = queue[0]
		queue = queue[1:]
		if len(res) == h {
			res = append(res, make([]int, 0))
		}
		res[h] = append(res[h], cur.Val)
		if cur.Left != nil {
			tmpQueue = append(tmpQueue, cur.Left)
		}
		if cur.Right != nil {
			tmpQueue = append(tmpQueue, cur.Right)
		}
		if len(queue) == 0 { //一层一层放入不同的queue
			queue = tmpQueue
			tmpQueue = make([]*TreeNode, 0)
			h++
		}
	}
	return res
}
