package main

import (
	"fmt"
	"github/gokangaroo/LeetCode/datastructure"
	"strings"
)

func main() {
	path := "/a//b////c/d//././/.."
	res := simplifyPath(path)
	fmt.Println(res)
}

//非常典型的出入栈
func simplifyPath(path string) string {
	var stack = make(datastructure.Stack, 0)
	var res strings.Builder
	//开始入栈
	pathArr := strings.Split(path, "/")
	for i := 0; i < len(pathArr); i++ {
		//多个/ 或者 当前目录
		if pathArr[i] == "" || pathArr[i] == "." {
			continue
		}
		//上级目录
		if pathArr[i] == ".." {
			if stack.Len() > 0 {
				stack.Pop()
			}
			continue
		}
		stack.Push(pathArr[i])
	}
	for _, v := range stack {
		res.WriteString(fmt.Sprintf("/%s", v.(string)))
	}
	if stack.Len() == 0 {
		return "/"
	}
	return res.String()
}
