package main

import (
	"fmt"
	"github/gokangaroo/LeetCode/datastructure"
	"path"
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

type dir struct {
	absolutePath string
	parentDir    *dir
	sons         map[string]*dir
}

func simplifyPath2(path string) string {
	p := strings.Split(path, "/")
	start := &dir{"/", nil, make(map[string]*dir)}
	cursor := start
	for _, v := range p {
		switch v {
		case ".", "":
			break
		case "..":
			if cursor.parentDir != nil {
				cursor = cursor.parentDir
			}
		default:
			// 是否是新节点
			if d, ok := cursor.sons[v]; ok {
				cursor = d
			} else {
				if cursor.absolutePath != "/" {
					cursor.sons[v] = &dir{cursor.absolutePath + "/" + v, cursor, make(map[string]*dir)}
				} else {
					cursor.sons[v] = &dir{cursor.absolutePath + v, cursor, make(map[string]*dir)}
				}
				cursor = cursor.sons[v]
			}
		}
	}
	return cursor.absolutePath
}

func simplifyPath3(path string) string {
	arr := strings.Split(path, "/")
	stack := make([]string, 0, len(arr))
	for i := 1; i < len(arr); i++ {
		switch arr[i] {
		case ".", "":
			continue
		case "..":
			// 出栈
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, arr[i])
		}
	}
	if len(stack) == 0 {
		return "/"
	}
	// 使用strings.Builder减少内存消耗, 提升速度
	var res = new(strings.Builder)
	for i := 0; i < len(stack); i++ {
		if stack[i] != "" {
			res.WriteString("/" + stack[i])
		}
	}
	return res.String()
}

// FIXME 标准库写法, 内存进一步降低..
func simplifyPath4(path2 string) string {
	return path.Clean(path2)
}
