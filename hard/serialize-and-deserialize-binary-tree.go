package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	x := &TreeNode{1,
		nil, &TreeNode{2,
			&TreeNode{3, nil, nil}, nil}}
	constructor := Constructor()
	serialize := (&constructor).serialize(x)
	fmt.Println(serialize)
	fmt.Println((&constructor).deserialize(serialize))
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	res := new(strings.Builder)
	res.WriteString("[")
	dfsStr(root, res)
	res.WriteString("]")
	return res.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	data = data[1 : len(data)-1]
	strSlice := strings.Split(data, ",")
	if len(strSlice) == 0 {
		return nil
	}
	var root = &TreeNode{0, nil, nil}
	_, err := splitTree(root, 0, &strSlice)
	if err != nil {
		root = nil
	}
	return root
}

func dfsStr(node *TreeNode, str *strings.Builder) {
	if str.Len() != 1 {
		str.WriteString(",")
	}
	if node == nil {
		str.WriteString("null")
		return
	} else {
		str.WriteString(fmt.Sprintf("%d", node.Val))
	}
	dfsStr(node.Left, str)
	dfsStr(node.Right, str)
}

func splitTree(node *TreeNode, idx int, all *[]string) (newIdx int, err error) {
	i, err := strconv.Atoi((*all)[idx])
	if err == nil {
		node.Val = i
		idx++
		node.Left = &TreeNode{0, nil, nil}
		idx, err = splitTree(node.Left, idx, all)
		if err != nil {
			node.Left = nil
		}
		node.Right = &TreeNode{0, nil, nil}
		idx, err = splitTree(node.Right, idx, all)
		if err != nil {
			node.Right = nil
		}
		return idx, nil
	}
	return idx + 1, err //返回error表示应该设置为nil
}
