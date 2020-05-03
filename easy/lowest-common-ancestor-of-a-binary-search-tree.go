package main

import (
	"fmt"
)

func main() {
	//[6,2,8,0,4,7,9,null,null,3,5]
	p := &TreeNode{2,
		&TreeNode{0, nil, nil}, &TreeNode{4,
			&TreeNode{3, nil, nil}, &TreeNode{5, nil, nil}}}
	q := &TreeNode{8,
		&TreeNode{7, nil, nil}, &TreeNode{9, nil, nil}}
	x := &TreeNode{6, p, q}
	fmt.Println(lowestCommonAncestor(x, p, q))
}

// TODO 使用二叉搜索树的特性, 可以尝试简单解决?

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	aim := make(chan *TreeNode)
	go dfsSons(root, p, q, aim)
	return <-aim
}

// 当一个节点左右都返回true表明是最近(深)的父节点
func dfsSons(node, p, q *TreeNode, aim chan *TreeNode) bool {
	if node == nil {
		return false
	}
	left := dfsSons(node.Left, p, q, aim)
	right := dfsSons(node.Right, p, q, aim)
	if node == p || node == q { //1.回溯 2.放到dfs后面,防止正好另一个是其子集
		if left || right {
			aim <- node
		}
		return true
	}
	if left && right {
		aim <- node
	}
	return left || right //3.收到这个的上层父节点应该另一半是false, 或者比这个慢, 无所谓
}

// 蠢笨方法层次遍历
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	var queue = []*tmpNode{&tmpNode{root, make([]bool, 0)}}
	var cur *tmpNode
	var pF, qF []bool
	for len(queue) != 0 && (pF == nil || qF == nil) {
		// 1.出队
		cur = queue[0]
		queue = queue[1:]
		// 2.找到路径
		if cur.TreeNode == p {
			pF = cur.path
		}
		if cur.TreeNode == q {
			qF = cur.path
		}
		// 2.1 入队
		if cur.Left != nil {
			queue = append(queue, &tmpNode{cur.Left, append(cur.path, true)})
		}
		if cur.Right != nil {
			tmp := make([]bool, len(cur.path), cap(cur.path)) // 拷贝
			copy(tmp, cur.path)                               //每个左侧一列共享一个path
			queue = append(queue, &tmpNode{cur.Right, append(tmp, false)})
		}
	}
	for i := 0; i < minLen(len(pF), len(qF)); i++ {
		if pF[i] == qF[i] {
			if pF[i] {
				root = root.Left
			} else {
				root = root.Right
			}
		} else {
			break
		}
	}
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type tmpNode struct {
	*TreeNode
	path []bool //true 为左, left为右
}

func minLen(x, y int) int {
	if x < y {
		return x
	}
	return y
}
