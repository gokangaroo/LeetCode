package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isValidSerialization("#"))
}

func isValidSerialization(preorder string) bool {
	slice := strings.Split(preorder, ",")
	count := 1 //可承载的叶子数量
	for i := 0; i < len(slice); i++ {
		count-- // 叶子数量减一
		if count < 0 {
			return false
		}
		if slice[i] != "#" {
			count += 2 // 数字可以加两个叶子节点
		}
	}
	return count == 0
}
