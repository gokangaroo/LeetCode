package main

import (
	// 导入包到当前包, 类似java static import
	. "github/gokangaroo/LeetCode/datastructure"

	"fmt"
)

func main() {
	str := "]"
	fmt.Println(isValid(str))
}

func isValid(s string) bool {
	st := make(Stack, 0)
	for _, c := range s {
		switch string(c) {
		case "(", "{", "[":
			st.Push(c)
		case "]":
			pope, err := st.Pop()
			if err != nil || "[" != string(pope.(int32)) {
				return false
			}
		case "}":
			pope, err := st.Pop()
			if err != nil || "{" != string(pope.(int32)) {
				return false
			}
		case ")":
			pope, err := st.Pop()
			if err != nil || "(" != string(pope.(int32)) {
				return false
			}
		}
	}
	if st.Len() == 0 {
		return true
	} else {
		return false
	}
}
